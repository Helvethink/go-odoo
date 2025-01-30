package odoo

// HrExpenseSplitWizard represents hr.expense.split.wizard model.
type HrExpenseSplitWizard struct {
	CreateDate                  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName                 *String   `xmlrpc:"display_name,omitempty"`
	ExpenseId                   *Many2One `xmlrpc:"expense_id,omitempty"`
	ExpenseSplitLineIds         *Relation `xmlrpc:"expense_split_line_ids,omitempty"`
	Id                          *Int      `xmlrpc:"id,omitempty"`
	SplitPossible               *Bool     `xmlrpc:"split_possible,omitempty"`
	TaxAmountCurrency           *Float    `xmlrpc:"tax_amount_currency,omitempty"`
	TotalAmountCurrency         *Float    `xmlrpc:"total_amount_currency,omitempty"`
	TotalAmountCurrencyOriginal *Float    `xmlrpc:"total_amount_currency_original,omitempty"`
	WriteDate                   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrExpenseSplitWizards represents array of hr.expense.split.wizard model.
type HrExpenseSplitWizards []HrExpenseSplitWizard

// HrExpenseSplitWizardModel is the odoo model name.
const HrExpenseSplitWizardModel = "hr.expense.split.wizard"

// Many2One convert HrExpenseSplitWizard to *Many2One.
func (hesw *HrExpenseSplitWizard) Many2One() *Many2One {
	return NewMany2One(hesw.Id.Get(), "")
}

// CreateHrExpenseSplitWizard creates a new hr.expense.split.wizard model and returns its id.
func (c *Client) CreateHrExpenseSplitWizard(hesw *HrExpenseSplitWizard) (int64, error) {
	ids, err := c.CreateHrExpenseSplitWizards([]*HrExpenseSplitWizard{hesw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrExpenseSplitWizard creates a new hr.expense.split.wizard model and returns its id.
func (c *Client) CreateHrExpenseSplitWizards(hesws []*HrExpenseSplitWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hesws {
		vv = append(vv, v)
	}
	return c.Create(HrExpenseSplitWizardModel, vv, nil)
}

// UpdateHrExpenseSplitWizard updates an existing hr.expense.split.wizard record.
func (c *Client) UpdateHrExpenseSplitWizard(hesw *HrExpenseSplitWizard) error {
	return c.UpdateHrExpenseSplitWizards([]int64{hesw.Id.Get()}, hesw)
}

// UpdateHrExpenseSplitWizards updates existing hr.expense.split.wizard records.
// All records (represented by ids) will be updated by hesw values.
func (c *Client) UpdateHrExpenseSplitWizards(ids []int64, hesw *HrExpenseSplitWizard) error {
	return c.Update(HrExpenseSplitWizardModel, ids, hesw, nil)
}

// DeleteHrExpenseSplitWizard deletes an existing hr.expense.split.wizard record.
func (c *Client) DeleteHrExpenseSplitWizard(id int64) error {
	return c.DeleteHrExpenseSplitWizards([]int64{id})
}

// DeleteHrExpenseSplitWizards deletes existing hr.expense.split.wizard records.
func (c *Client) DeleteHrExpenseSplitWizards(ids []int64) error {
	return c.Delete(HrExpenseSplitWizardModel, ids)
}

// GetHrExpenseSplitWizard gets hr.expense.split.wizard existing record.
func (c *Client) GetHrExpenseSplitWizard(id int64) (*HrExpenseSplitWizard, error) {
	hesws, err := c.GetHrExpenseSplitWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hesws)[0]), nil
}

// GetHrExpenseSplitWizards gets hr.expense.split.wizard existing records.
func (c *Client) GetHrExpenseSplitWizards(ids []int64) (*HrExpenseSplitWizards, error) {
	hesws := &HrExpenseSplitWizards{}
	if err := c.Read(HrExpenseSplitWizardModel, ids, nil, hesws); err != nil {
		return nil, err
	}
	return hesws, nil
}

// FindHrExpenseSplitWizard finds hr.expense.split.wizard record by querying it with criteria.
func (c *Client) FindHrExpenseSplitWizard(criteria *Criteria) (*HrExpenseSplitWizard, error) {
	hesws := &HrExpenseSplitWizards{}
	if err := c.SearchRead(HrExpenseSplitWizardModel, criteria, NewOptions().Limit(1), hesws); err != nil {
		return nil, err
	}
	return &((*hesws)[0]), nil
}

// FindHrExpenseSplitWizards finds hr.expense.split.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSplitWizards(criteria *Criteria, options *Options) (*HrExpenseSplitWizards, error) {
	hesws := &HrExpenseSplitWizards{}
	if err := c.SearchRead(HrExpenseSplitWizardModel, criteria, options, hesws); err != nil {
		return nil, err
	}
	return hesws, nil
}

// FindHrExpenseSplitWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSplitWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrExpenseSplitWizardModel, criteria, options)
}

// FindHrExpenseSplitWizardId finds record id by querying it with criteria.
func (c *Client) FindHrExpenseSplitWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrExpenseSplitWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
