package odoo

// HrExpenseRefuseWizard represents hr.expense.refuse.wizard model.
type HrExpenseRefuseWizard struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Reason      *String   `xmlrpc:"reason,omitempty"`
	SheetIds    *Relation `xmlrpc:"sheet_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrExpenseRefuseWizards represents array of hr.expense.refuse.wizard model.
type HrExpenseRefuseWizards []HrExpenseRefuseWizard

// HrExpenseRefuseWizardModel is the odoo model name.
const HrExpenseRefuseWizardModel = "hr.expense.refuse.wizard"

// Many2One convert HrExpenseRefuseWizard to *Many2One.
func (herw *HrExpenseRefuseWizard) Many2One() *Many2One {
	return NewMany2One(herw.Id.Get(), "")
}

// CreateHrExpenseRefuseWizard creates a new hr.expense.refuse.wizard model and returns its id.
func (c *Client) CreateHrExpenseRefuseWizard(herw *HrExpenseRefuseWizard) (int64, error) {
	ids, err := c.CreateHrExpenseRefuseWizards([]*HrExpenseRefuseWizard{herw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrExpenseRefuseWizard creates a new hr.expense.refuse.wizard model and returns its id.
func (c *Client) CreateHrExpenseRefuseWizards(herws []*HrExpenseRefuseWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range herws {
		vv = append(vv, v)
	}
	return c.Create(HrExpenseRefuseWizardModel, vv, nil)
}

// UpdateHrExpenseRefuseWizard updates an existing hr.expense.refuse.wizard record.
func (c *Client) UpdateHrExpenseRefuseWizard(herw *HrExpenseRefuseWizard) error {
	return c.UpdateHrExpenseRefuseWizards([]int64{herw.Id.Get()}, herw)
}

// UpdateHrExpenseRefuseWizards updates existing hr.expense.refuse.wizard records.
// All records (represented by ids) will be updated by herw values.
func (c *Client) UpdateHrExpenseRefuseWizards(ids []int64, herw *HrExpenseRefuseWizard) error {
	return c.Update(HrExpenseRefuseWizardModel, ids, herw, nil)
}

// DeleteHrExpenseRefuseWizard deletes an existing hr.expense.refuse.wizard record.
func (c *Client) DeleteHrExpenseRefuseWizard(id int64) error {
	return c.DeleteHrExpenseRefuseWizards([]int64{id})
}

// DeleteHrExpenseRefuseWizards deletes existing hr.expense.refuse.wizard records.
func (c *Client) DeleteHrExpenseRefuseWizards(ids []int64) error {
	return c.Delete(HrExpenseRefuseWizardModel, ids)
}

// GetHrExpenseRefuseWizard gets hr.expense.refuse.wizard existing record.
func (c *Client) GetHrExpenseRefuseWizard(id int64) (*HrExpenseRefuseWizard, error) {
	herws, err := c.GetHrExpenseRefuseWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*herws)[0]), nil
}

// GetHrExpenseRefuseWizards gets hr.expense.refuse.wizard existing records.
func (c *Client) GetHrExpenseRefuseWizards(ids []int64) (*HrExpenseRefuseWizards, error) {
	herws := &HrExpenseRefuseWizards{}
	if err := c.Read(HrExpenseRefuseWizardModel, ids, nil, herws); err != nil {
		return nil, err
	}
	return herws, nil
}

// FindHrExpenseRefuseWizard finds hr.expense.refuse.wizard record by querying it with criteria.
func (c *Client) FindHrExpenseRefuseWizard(criteria *Criteria) (*HrExpenseRefuseWizard, error) {
	herws := &HrExpenseRefuseWizards{}
	if err := c.SearchRead(HrExpenseRefuseWizardModel, criteria, NewOptions().Limit(1), herws); err != nil {
		return nil, err
	}
	return &((*herws)[0]), nil
}

// FindHrExpenseRefuseWizards finds hr.expense.refuse.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseRefuseWizards(criteria *Criteria, options *Options) (*HrExpenseRefuseWizards, error) {
	herws := &HrExpenseRefuseWizards{}
	if err := c.SearchRead(HrExpenseRefuseWizardModel, criteria, options, herws); err != nil {
		return nil, err
	}
	return herws, nil
}

// FindHrExpenseRefuseWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseRefuseWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrExpenseRefuseWizardModel, criteria, options)
}

// FindHrExpenseRefuseWizardId finds record id by querying it with criteria.
func (c *Client) FindHrExpenseRefuseWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrExpenseRefuseWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
