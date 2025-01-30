package odoo

// HrExpenseSplit represents hr.expense.split model.
type HrExpenseSplit struct {
	AnalyticDistribution           *String   `xmlrpc:"analytic_distribution,omitempty"`
	AnalyticPrecision              *Int      `xmlrpc:"analytic_precision,omitempty"`
	CanBeReinvoiced                *Bool     `xmlrpc:"can_be_reinvoiced,omitempty"`
	CompanyId                      *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate                     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                      *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId                     *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName                    *String   `xmlrpc:"display_name,omitempty"`
	DistributionAnalyticAccountIds *Relation `xmlrpc:"distribution_analytic_account_ids,omitempty"`
	EmployeeId                     *Many2One `xmlrpc:"employee_id,omitempty"`
	ExpenseId                      *Many2One `xmlrpc:"expense_id,omitempty"`
	Id                             *Int      `xmlrpc:"id,omitempty"`
	Name                           *String   `xmlrpc:"name,omitempty"`
	ProductHasCost                 *Bool     `xmlrpc:"product_has_cost,omitempty"`
	ProductHasTax                  *Bool     `xmlrpc:"product_has_tax,omitempty"`
	ProductId                      *Many2One `xmlrpc:"product_id,omitempty"`
	SaleOrderId                    *Many2One `xmlrpc:"sale_order_id,omitempty"`
	TaxAmountCurrency              *Float    `xmlrpc:"tax_amount_currency,omitempty"`
	TaxIds                         *Relation `xmlrpc:"tax_ids,omitempty"`
	TotalAmountCurrency            *Float    `xmlrpc:"total_amount_currency,omitempty"`
	WizardId                       *Many2One `xmlrpc:"wizard_id,omitempty"`
	WriteDate                      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrExpenseSplits represents array of hr.expense.split model.
type HrExpenseSplits []HrExpenseSplit

// HrExpenseSplitModel is the odoo model name.
const HrExpenseSplitModel = "hr.expense.split"

// Many2One convert HrExpenseSplit to *Many2One.
func (hes *HrExpenseSplit) Many2One() *Many2One {
	return NewMany2One(hes.Id.Get(), "")
}

// CreateHrExpenseSplit creates a new hr.expense.split model and returns its id.
func (c *Client) CreateHrExpenseSplit(hes *HrExpenseSplit) (int64, error) {
	ids, err := c.CreateHrExpenseSplits([]*HrExpenseSplit{hes})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrExpenseSplit creates a new hr.expense.split model and returns its id.
func (c *Client) CreateHrExpenseSplits(hess []*HrExpenseSplit) ([]int64, error) {
	var vv []interface{}
	for _, v := range hess {
		vv = append(vv, v)
	}
	return c.Create(HrExpenseSplitModel, vv, nil)
}

// UpdateHrExpenseSplit updates an existing hr.expense.split record.
func (c *Client) UpdateHrExpenseSplit(hes *HrExpenseSplit) error {
	return c.UpdateHrExpenseSplits([]int64{hes.Id.Get()}, hes)
}

// UpdateHrExpenseSplits updates existing hr.expense.split records.
// All records (represented by ids) will be updated by hes values.
func (c *Client) UpdateHrExpenseSplits(ids []int64, hes *HrExpenseSplit) error {
	return c.Update(HrExpenseSplitModel, ids, hes, nil)
}

// DeleteHrExpenseSplit deletes an existing hr.expense.split record.
func (c *Client) DeleteHrExpenseSplit(id int64) error {
	return c.DeleteHrExpenseSplits([]int64{id})
}

// DeleteHrExpenseSplits deletes existing hr.expense.split records.
func (c *Client) DeleteHrExpenseSplits(ids []int64) error {
	return c.Delete(HrExpenseSplitModel, ids)
}

// GetHrExpenseSplit gets hr.expense.split existing record.
func (c *Client) GetHrExpenseSplit(id int64) (*HrExpenseSplit, error) {
	hess, err := c.GetHrExpenseSplits([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hess)[0]), nil
}

// GetHrExpenseSplits gets hr.expense.split existing records.
func (c *Client) GetHrExpenseSplits(ids []int64) (*HrExpenseSplits, error) {
	hess := &HrExpenseSplits{}
	if err := c.Read(HrExpenseSplitModel, ids, nil, hess); err != nil {
		return nil, err
	}
	return hess, nil
}

// FindHrExpenseSplit finds hr.expense.split record by querying it with criteria.
func (c *Client) FindHrExpenseSplit(criteria *Criteria) (*HrExpenseSplit, error) {
	hess := &HrExpenseSplits{}
	if err := c.SearchRead(HrExpenseSplitModel, criteria, NewOptions().Limit(1), hess); err != nil {
		return nil, err
	}
	return &((*hess)[0]), nil
}

// FindHrExpenseSplits finds hr.expense.split records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSplits(criteria *Criteria, options *Options) (*HrExpenseSplits, error) {
	hess := &HrExpenseSplits{}
	if err := c.SearchRead(HrExpenseSplitModel, criteria, options, hess); err != nil {
		return nil, err
	}
	return hess, nil
}

// FindHrExpenseSplitIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSplitIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrExpenseSplitModel, criteria, options)
}

// FindHrExpenseSplitId finds record id by querying it with criteria.
func (c *Client) FindHrExpenseSplitId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrExpenseSplitModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
