package cmd

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/helvethink/go-odoo"
)

// GeneratorConfiguration is the configuration to create a new *generator by injecting
// its dependencies.
type GeneratorConfiguration struct {
	Odoo          *odoo.Client
	ModelTemplate *template.Template
	FormatModels  bool
	DestFolder    string
}

type Generator struct {
	odoo         *odoo.Client
	tmpl         *template.Template
	destFolder   string
	formatModels bool
}

// NewGenerator creates a new *generator.
func NewGenerator(cfg GeneratorConfiguration) *Generator {
	return &Generator{odoo: cfg.Odoo, tmpl: cfg.ModelTemplate, formatModels: cfg.FormatModels, destFolder: cfg.DestFolder}
}

func (g *Generator) handleModels(models []string) error {
	mm, err := g.getModels(models)
	if err != nil {
		return err
	}
	if err := g.generateModels(mm); err != nil {
		return err
	}
	return nil
}

func (g *Generator) getModels(models []string) ([]*model, error) {
	if len(models) == 0 {
		var err error
		models, err = g.getAllModelsName()
		if err != nil {
			return nil, err
		}
	}
	var mm []*model
	for _, model := range models {
		mfs, err := g.modelFieldsFromModel(model)
		if err != nil {
			return nil, err
		}
		if len(mfs) == 0 {
			log.Printf("error: cannot find fields for model %s, cannot generate it.\n", model)
			continue
		}
		idExists := false
		for _, mf := range mfs {
			if mf.Name == "id" {
				idExists = true
				break
			}
		}
		if !idExists {
			log.Printf("error: cannot find ID field for model %s, cannot generate it.\n", model)
			continue
		}
		mm = append(mm, newModel(model, mfs))
	}
	return mm, nil
}

func (g *Generator) getAllModelsName() ([]string, error) {
	var models []string
	ims, err := g.odoo.FindIrModels(nil, nil)
	if err != nil {
		return models, err
	}
	for _, im := range *ims {
		models = append(models, im.Model.Get())
	}
	return models, nil
}

func (g *Generator) modelFieldsFromModel(model string) ([]*modelField, error) {
	imfs, err := g.odoo.FindIrModelFieldss(odoo.NewCriteria().Add("model", "=", model), nil)
	if err != nil {
		if errors.Is(err, odoo.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return g.irModelFieldsToModelFields(imfs), nil
}

func (g *Generator) irModelFieldsToModelFields(imfs *odoo.IrModelFieldss) []*modelField {
	var mfs []*modelField
	for _, imf := range *imfs {
		mfs = append(mfs, newModelField(imf.Name.Get(), imf.Ttype.Get().(string)))
	}
	return mfs
}

func (g *Generator) generateModels(models []*model) error {
	for _, m := range models {
		filePath := g.destFolder + "/" + strings.Replace(m.Name, ".", "_", -1) + ".go"
		output, err := os.Create(filePath)
		if err != nil {
			return err
		}
		if err := g.tmpl.Execute(output, m); err != nil {
			return err
		}
		if g.formatModels {
			if err := exec.Command("gofmt", "-w", filePath).Run(); err != nil {
				return err
			}
		}
		log.Printf("%s has been generated\n", filePath)
	}
	return nil
}
