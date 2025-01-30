package odoo

// HrContract represents hr.contract model.
type HrContract struct {
	Active                      *Bool      `xmlrpc:"active,omitempty"`
	ActiveEmployee              *Bool      `xmlrpc:"active_employee,omitempty"`
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	CalendarMismatch            *Bool      `xmlrpc:"calendar_mismatch,omitempty"`
	CompanyCountryId            *Many2One  `xmlrpc:"company_country_id,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	ContractTypeId              *Many2One  `xmlrpc:"contract_type_id,omitempty"`
	ContractWage                *Float     `xmlrpc:"contract_wage,omitempty"`
	ContractsCount              *Int       `xmlrpc:"contracts_count,omitempty"`
	CountryCode                 *String    `xmlrpc:"country_code,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	DateEnd                     *Time      `xmlrpc:"date_end,omitempty"`
	DateStart                   *Time      `xmlrpc:"date_start,omitempty"`
	DepartmentId                *Many2One  `xmlrpc:"department_id,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omitempty"`
	FirstContractDate           *Time      `xmlrpc:"first_contract_date,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	HrResponsibleId             *Many2One  `xmlrpc:"hr_responsible_id,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	JobId                       *Many2One  `xmlrpc:"job_id,omitempty"`
	KanbanState                 *Selection `xmlrpc:"kanban_state,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	Notes                       *String    `xmlrpc:"notes,omitempty"`
	PermitNo                    *String    `xmlrpc:"permit_no,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	ResourceCalendarId          *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	StructureTypeId             *Many2One  `xmlrpc:"structure_type_id,omitempty"`
	TrialDateEnd                *Time      `xmlrpc:"trial_date_end,omitempty"`
	VisaNo                      *String    `xmlrpc:"visa_no,omitempty"`
	Wage                        *Float     `xmlrpc:"wage,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrContracts represents array of hr.contract model.
type HrContracts []HrContract

// HrContractModel is the odoo model name.
const HrContractModel = "hr.contract"

// Many2One convert HrContract to *Many2One.
func (hc *HrContract) Many2One() *Many2One {
	return NewMany2One(hc.Id.Get(), "")
}

// CreateHrContract creates a new hr.contract model and returns its id.
func (c *Client) CreateHrContract(hc *HrContract) (int64, error) {
	ids, err := c.CreateHrContracts([]*HrContract{hc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrContract creates a new hr.contract model and returns its id.
func (c *Client) CreateHrContracts(hcs []*HrContract) ([]int64, error) {
	var vv []interface{}
	for _, v := range hcs {
		vv = append(vv, v)
	}
	return c.Create(HrContractModel, vv, nil)
}

// UpdateHrContract updates an existing hr.contract record.
func (c *Client) UpdateHrContract(hc *HrContract) error {
	return c.UpdateHrContracts([]int64{hc.Id.Get()}, hc)
}

// UpdateHrContracts updates existing hr.contract records.
// All records (represented by ids) will be updated by hc values.
func (c *Client) UpdateHrContracts(ids []int64, hc *HrContract) error {
	return c.Update(HrContractModel, ids, hc, nil)
}

// DeleteHrContract deletes an existing hr.contract record.
func (c *Client) DeleteHrContract(id int64) error {
	return c.DeleteHrContracts([]int64{id})
}

// DeleteHrContracts deletes existing hr.contract records.
func (c *Client) DeleteHrContracts(ids []int64) error {
	return c.Delete(HrContractModel, ids)
}

// GetHrContract gets hr.contract existing record.
func (c *Client) GetHrContract(id int64) (*HrContract, error) {
	hcs, err := c.GetHrContracts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hcs)[0]), nil
}

// GetHrContracts gets hr.contract existing records.
func (c *Client) GetHrContracts(ids []int64) (*HrContracts, error) {
	hcs := &HrContracts{}
	if err := c.Read(HrContractModel, ids, nil, hcs); err != nil {
		return nil, err
	}
	return hcs, nil
}

// FindHrContract finds hr.contract record by querying it with criteria.
func (c *Client) FindHrContract(criteria *Criteria) (*HrContract, error) {
	hcs := &HrContracts{}
	if err := c.SearchRead(HrContractModel, criteria, NewOptions().Limit(1), hcs); err != nil {
		return nil, err
	}
	return &((*hcs)[0]), nil
}

// FindHrContracts finds hr.contract records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContracts(criteria *Criteria, options *Options) (*HrContracts, error) {
	hcs := &HrContracts{}
	if err := c.SearchRead(HrContractModel, criteria, options, hcs); err != nil {
		return nil, err
	}
	return hcs, nil
}

// FindHrContractIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContractIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrContractModel, criteria, options)
}

// FindHrContractId finds record id by querying it with criteria.
func (c *Client) FindHrContractId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrContractModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
