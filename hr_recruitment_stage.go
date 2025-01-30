package odoo

// HrRecruitmentStage represents hr.recruitment.stage model.
type HrRecruitmentStage struct {
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	Fold             *Bool     `xmlrpc:"fold,omitempty"`
	HiredStage       *Bool     `xmlrpc:"hired_stage,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	IsWarningVisible *Bool     `xmlrpc:"is_warning_visible,omitempty"`
	JobIds           *Relation `xmlrpc:"job_ids,omitempty"`
	LegendBlocked    *String   `xmlrpc:"legend_blocked,omitempty"`
	LegendDone       *String   `xmlrpc:"legend_done,omitempty"`
	LegendNormal     *String   `xmlrpc:"legend_normal,omitempty"`
	Name             *String   `xmlrpc:"name,omitempty"`
	Requirements     *String   `xmlrpc:"requirements,omitempty"`
	Sequence         *Int      `xmlrpc:"sequence,omitempty"`
	TemplateId       *Many2One `xmlrpc:"template_id,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrRecruitmentStages represents array of hr.recruitment.stage model.
type HrRecruitmentStages []HrRecruitmentStage

// HrRecruitmentStageModel is the odoo model name.
const HrRecruitmentStageModel = "hr.recruitment.stage"

// Many2One convert HrRecruitmentStage to *Many2One.
func (hrs *HrRecruitmentStage) Many2One() *Many2One {
	return NewMany2One(hrs.Id.Get(), "")
}

// CreateHrRecruitmentStage creates a new hr.recruitment.stage model and returns its id.
func (c *Client) CreateHrRecruitmentStage(hrs *HrRecruitmentStage) (int64, error) {
	ids, err := c.CreateHrRecruitmentStages([]*HrRecruitmentStage{hrs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrRecruitmentStage creates a new hr.recruitment.stage model and returns its id.
func (c *Client) CreateHrRecruitmentStages(hrss []*HrRecruitmentStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range hrss {
		vv = append(vv, v)
	}
	return c.Create(HrRecruitmentStageModel, vv, nil)
}

// UpdateHrRecruitmentStage updates an existing hr.recruitment.stage record.
func (c *Client) UpdateHrRecruitmentStage(hrs *HrRecruitmentStage) error {
	return c.UpdateHrRecruitmentStages([]int64{hrs.Id.Get()}, hrs)
}

// UpdateHrRecruitmentStages updates existing hr.recruitment.stage records.
// All records (represented by ids) will be updated by hrs values.
func (c *Client) UpdateHrRecruitmentStages(ids []int64, hrs *HrRecruitmentStage) error {
	return c.Update(HrRecruitmentStageModel, ids, hrs, nil)
}

// DeleteHrRecruitmentStage deletes an existing hr.recruitment.stage record.
func (c *Client) DeleteHrRecruitmentStage(id int64) error {
	return c.DeleteHrRecruitmentStages([]int64{id})
}

// DeleteHrRecruitmentStages deletes existing hr.recruitment.stage records.
func (c *Client) DeleteHrRecruitmentStages(ids []int64) error {
	return c.Delete(HrRecruitmentStageModel, ids)
}

// GetHrRecruitmentStage gets hr.recruitment.stage existing record.
func (c *Client) GetHrRecruitmentStage(id int64) (*HrRecruitmentStage, error) {
	hrss, err := c.GetHrRecruitmentStages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hrss)[0]), nil
}

// GetHrRecruitmentStages gets hr.recruitment.stage existing records.
func (c *Client) GetHrRecruitmentStages(ids []int64) (*HrRecruitmentStages, error) {
	hrss := &HrRecruitmentStages{}
	if err := c.Read(HrRecruitmentStageModel, ids, nil, hrss); err != nil {
		return nil, err
	}
	return hrss, nil
}

// FindHrRecruitmentStage finds hr.recruitment.stage record by querying it with criteria.
func (c *Client) FindHrRecruitmentStage(criteria *Criteria) (*HrRecruitmentStage, error) {
	hrss := &HrRecruitmentStages{}
	if err := c.SearchRead(HrRecruitmentStageModel, criteria, NewOptions().Limit(1), hrss); err != nil {
		return nil, err
	}
	return &((*hrss)[0]), nil
}

// FindHrRecruitmentStages finds hr.recruitment.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRecruitmentStages(criteria *Criteria, options *Options) (*HrRecruitmentStages, error) {
	hrss := &HrRecruitmentStages{}
	if err := c.SearchRead(HrRecruitmentStageModel, criteria, options, hrss); err != nil {
		return nil, err
	}
	return hrss, nil
}

// FindHrRecruitmentStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRecruitmentStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrRecruitmentStageModel, criteria, options)
}

// FindHrRecruitmentStageId finds record id by querying it with criteria.
func (c *Client) FindHrRecruitmentStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrRecruitmentStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
