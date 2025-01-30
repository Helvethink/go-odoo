package odoo

// HrExpenseApproveDuplicate represents hr.expense.approve.duplicate model.
type HrExpenseApproveDuplicate struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	ExpenseIds  *Relation `xmlrpc:"expense_ids,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	SheetIds    *Relation `xmlrpc:"sheet_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrExpenseApproveDuplicates represents array of hr.expense.approve.duplicate model.
type HrExpenseApproveDuplicates []HrExpenseApproveDuplicate

// HrExpenseApproveDuplicateModel is the odoo model name.
const HrExpenseApproveDuplicateModel = "hr.expense.approve.duplicate"

// Many2One convert HrExpenseApproveDuplicate to *Many2One.
func (head *HrExpenseApproveDuplicate) Many2One() *Many2One {
	return NewMany2One(head.Id.Get(), "")
}

// CreateHrExpenseApproveDuplicate creates a new hr.expense.approve.duplicate model and returns its id.
func (c *Client) CreateHrExpenseApproveDuplicate(head *HrExpenseApproveDuplicate) (int64, error) {
	ids, err := c.CreateHrExpenseApproveDuplicates([]*HrExpenseApproveDuplicate{head})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrExpenseApproveDuplicate creates a new hr.expense.approve.duplicate model and returns its id.
func (c *Client) CreateHrExpenseApproveDuplicates(heads []*HrExpenseApproveDuplicate) ([]int64, error) {
	var vv []interface{}
	for _, v := range heads {
		vv = append(vv, v)
	}
	return c.Create(HrExpenseApproveDuplicateModel, vv, nil)
}

// UpdateHrExpenseApproveDuplicate updates an existing hr.expense.approve.duplicate record.
func (c *Client) UpdateHrExpenseApproveDuplicate(head *HrExpenseApproveDuplicate) error {
	return c.UpdateHrExpenseApproveDuplicates([]int64{head.Id.Get()}, head)
}

// UpdateHrExpenseApproveDuplicates updates existing hr.expense.approve.duplicate records.
// All records (represented by ids) will be updated by head values.
func (c *Client) UpdateHrExpenseApproveDuplicates(ids []int64, head *HrExpenseApproveDuplicate) error {
	return c.Update(HrExpenseApproveDuplicateModel, ids, head, nil)
}

// DeleteHrExpenseApproveDuplicate deletes an existing hr.expense.approve.duplicate record.
func (c *Client) DeleteHrExpenseApproveDuplicate(id int64) error {
	return c.DeleteHrExpenseApproveDuplicates([]int64{id})
}

// DeleteHrExpenseApproveDuplicates deletes existing hr.expense.approve.duplicate records.
func (c *Client) DeleteHrExpenseApproveDuplicates(ids []int64) error {
	return c.Delete(HrExpenseApproveDuplicateModel, ids)
}

// GetHrExpenseApproveDuplicate gets hr.expense.approve.duplicate existing record.
func (c *Client) GetHrExpenseApproveDuplicate(id int64) (*HrExpenseApproveDuplicate, error) {
	heads, err := c.GetHrExpenseApproveDuplicates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*heads)[0]), nil
}

// GetHrExpenseApproveDuplicates gets hr.expense.approve.duplicate existing records.
func (c *Client) GetHrExpenseApproveDuplicates(ids []int64) (*HrExpenseApproveDuplicates, error) {
	heads := &HrExpenseApproveDuplicates{}
	if err := c.Read(HrExpenseApproveDuplicateModel, ids, nil, heads); err != nil {
		return nil, err
	}
	return heads, nil
}

// FindHrExpenseApproveDuplicate finds hr.expense.approve.duplicate record by querying it with criteria.
func (c *Client) FindHrExpenseApproveDuplicate(criteria *Criteria) (*HrExpenseApproveDuplicate, error) {
	heads := &HrExpenseApproveDuplicates{}
	if err := c.SearchRead(HrExpenseApproveDuplicateModel, criteria, NewOptions().Limit(1), heads); err != nil {
		return nil, err
	}
	return &((*heads)[0]), nil
}

// FindHrExpenseApproveDuplicates finds hr.expense.approve.duplicate records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseApproveDuplicates(criteria *Criteria, options *Options) (*HrExpenseApproveDuplicates, error) {
	heads := &HrExpenseApproveDuplicates{}
	if err := c.SearchRead(HrExpenseApproveDuplicateModel, criteria, options, heads); err != nil {
		return nil, err
	}
	return heads, nil
}

// FindHrExpenseApproveDuplicateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseApproveDuplicateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrExpenseApproveDuplicateModel, criteria, options)
}

// FindHrExpenseApproveDuplicateId finds record id by querying it with criteria.
func (c *Client) FindHrExpenseApproveDuplicateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrExpenseApproveDuplicateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
