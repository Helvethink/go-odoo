package odoo

// TimesheetsAnalysisReporting represents timesheets.analysis.reporting model.
type TimesheetsAnalysisReporting struct {
	Amount                     *Float     `xmlrpc:"amount,omitempty"`
	BillableTime               *Float     `xmlrpc:"billable_time,omitempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omitempty"`
	CurrencyId                 *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                       *Time      `xmlrpc:"date,omitempty"`
	DepartmentId               *Many2One  `xmlrpc:"department_id,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	EmployeeId                 *Many2One  `xmlrpc:"employee_id,omitempty"`
	HasDepartmentManagerAccess *Bool      `xmlrpc:"has_department_manager_access,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	InvoiceDate                *Time      `xmlrpc:"invoice_date,omitempty"`
	InvoiceDueDate             *Time      `xmlrpc:"invoice_due_date,omitempty"`
	ManagerId                  *Many2One  `xmlrpc:"manager_id,omitempty"`
	MessagePartnerIds          *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MilestoneId                *Many2One  `xmlrpc:"milestone_id,omitempty"`
	Name                       *String    `xmlrpc:"name,omitempty"`
	OrderId                    *Many2One  `xmlrpc:"order_id,omitempty"`
	ParentTaskId               *Many2One  `xmlrpc:"parent_task_id,omitempty"`
	PartnerId                  *Many2One  `xmlrpc:"partner_id,omitempty"`
	PayRef                     *String    `xmlrpc:"pay_ref,omitempty"`
	PayState                   *String    `xmlrpc:"pay_state,omitempty"`
	ProjectId                  *Many2One  `xmlrpc:"project_id,omitempty"`
	SoLine                     *Many2One  `xmlrpc:"so_line,omitempty"`
	TaskId                     *Many2One  `xmlrpc:"task_id,omitempty"`
	TaskName                   *String    `xmlrpc:"task_name,omitempty"`
	TimesheetInvoiceId         *Many2One  `xmlrpc:"timesheet_invoice_id,omitempty"`
	TimesheetInvoiceType       *Selection `xmlrpc:"timesheet_invoice_type,omitempty"`
	TimesheetRevenues          *Float     `xmlrpc:"timesheet_revenues,omitempty"`
	UnitAmount                 *Float     `xmlrpc:"unit_amount,omitempty"`
	UserId                     *Many2One  `xmlrpc:"user_id,omitempty"`
}

// TimesheetsAnalysisReportings represents array of timesheets.analysis.reporting model.
type TimesheetsAnalysisReportings []TimesheetsAnalysisReporting

// TimesheetsAnalysisReportingModel is the odoo model name.
const TimesheetsAnalysisReportingModel = "timesheets.analysis.reporting"

// Many2One convert TimesheetsAnalysisReporting to *Many2One.
func (tar *TimesheetsAnalysisReporting) Many2One() *Many2One {
	return NewMany2One(tar.Id.Get(), "")
}

// CreateTimesheetsAnalysisReporting creates a new timesheets.analysis.reporting model and returns its id.
func (c *Client) CreateTimesheetsAnalysisReporting(tar *TimesheetsAnalysisReporting) (int64, error) {
	ids, err := c.CreateTimesheetsAnalysisReportings([]*TimesheetsAnalysisReporting{tar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTimesheetsAnalysisReporting creates a new timesheets.analysis.reporting model and returns its id.
func (c *Client) CreateTimesheetsAnalysisReportings(tars []*TimesheetsAnalysisReporting) ([]int64, error) {
	var vv []interface{}
	for _, v := range tars {
		vv = append(vv, v)
	}
	return c.Create(TimesheetsAnalysisReportingModel, vv, nil)
}

// UpdateTimesheetsAnalysisReporting updates an existing timesheets.analysis.reporting record.
func (c *Client) UpdateTimesheetsAnalysisReporting(tar *TimesheetsAnalysisReporting) error {
	return c.UpdateTimesheetsAnalysisReportings([]int64{tar.Id.Get()}, tar)
}

// UpdateTimesheetsAnalysisReportings updates existing timesheets.analysis.reporting records.
// All records (represented by ids) will be updated by tar values.
func (c *Client) UpdateTimesheetsAnalysisReportings(ids []int64, tar *TimesheetsAnalysisReporting) error {
	return c.Update(TimesheetsAnalysisReportingModel, ids, tar, nil)
}

// DeleteTimesheetsAnalysisReporting deletes an existing timesheets.analysis.reporting record.
func (c *Client) DeleteTimesheetsAnalysisReporting(id int64) error {
	return c.DeleteTimesheetsAnalysisReportings([]int64{id})
}

// DeleteTimesheetsAnalysisReportings deletes existing timesheets.analysis.reporting records.
func (c *Client) DeleteTimesheetsAnalysisReportings(ids []int64) error {
	return c.Delete(TimesheetsAnalysisReportingModel, ids)
}

// GetTimesheetsAnalysisReporting gets timesheets.analysis.reporting existing record.
func (c *Client) GetTimesheetsAnalysisReporting(id int64) (*TimesheetsAnalysisReporting, error) {
	tars, err := c.GetTimesheetsAnalysisReportings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*tars)[0]), nil
}

// GetTimesheetsAnalysisReportings gets timesheets.analysis.reporting existing records.
func (c *Client) GetTimesheetsAnalysisReportings(ids []int64) (*TimesheetsAnalysisReportings, error) {
	tars := &TimesheetsAnalysisReportings{}
	if err := c.Read(TimesheetsAnalysisReportingModel, ids, nil, tars); err != nil {
		return nil, err
	}
	return tars, nil
}

// FindTimesheetsAnalysisReporting finds timesheets.analysis.reporting record by querying it with criteria.
func (c *Client) FindTimesheetsAnalysisReporting(criteria *Criteria) (*TimesheetsAnalysisReporting, error) {
	tars := &TimesheetsAnalysisReportings{}
	if err := c.SearchRead(TimesheetsAnalysisReportingModel, criteria, NewOptions().Limit(1), tars); err != nil {
		return nil, err
	}
	return &((*tars)[0]), nil
}

// FindTimesheetsAnalysisReportings finds timesheets.analysis.reporting records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimesheetsAnalysisReportings(criteria *Criteria, options *Options) (*TimesheetsAnalysisReportings, error) {
	tars := &TimesheetsAnalysisReportings{}
	if err := c.SearchRead(TimesheetsAnalysisReportingModel, criteria, options, tars); err != nil {
		return nil, err
	}
	return tars, nil
}

// FindTimesheetsAnalysisReportingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimesheetsAnalysisReportingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(TimesheetsAnalysisReportingModel, criteria, options)
}

// FindTimesheetsAnalysisReportingId finds record id by querying it with criteria.
func (c *Client) FindTimesheetsAnalysisReportingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TimesheetsAnalysisReportingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
