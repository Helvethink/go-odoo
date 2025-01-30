package odoo

// HrContractHistory represents hr.contract.history model.
type HrContractHistory struct {
	ActiveEmployee     *Bool      `xmlrpc:"active_employee,omitempty"`
	ActivityState      *Selection `xmlrpc:"activity_state,omitempty"`
	CompanyCountryId   *Many2One  `xmlrpc:"company_country_id,omitempty"`
	CompanyId          *Many2One  `xmlrpc:"company_id,omitempty"`
	ContractCount      *Int       `xmlrpc:"contract_count,omitempty"`
	ContractId         *Many2One  `xmlrpc:"contract_id,omitempty"`
	ContractIds        *Relation  `xmlrpc:"contract_ids,omitempty"`
	ContractTypeId     *Many2One  `xmlrpc:"contract_type_id,omitempty"`
	CountryCode        *String    `xmlrpc:"country_code,omitempty"`
	CurrencyId         *Many2One  `xmlrpc:"currency_id,omitempty"`
	DateEnd            *Time      `xmlrpc:"date_end,omitempty"`
	DateHired          *Time      `xmlrpc:"date_hired,omitempty"`
	DateStart          *Time      `xmlrpc:"date_start,omitempty"`
	DepartmentId       *Many2One  `xmlrpc:"department_id,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	EmployeeId         *Many2One  `xmlrpc:"employee_id,omitempty"`
	HrResponsibleId    *Many2One  `xmlrpc:"hr_responsible_id,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	IsUnderContract    *Bool      `xmlrpc:"is_under_contract,omitempty"`
	JobId              *Many2One  `xmlrpc:"job_id,omitempty"`
	Name               *String    `xmlrpc:"name,omitempty"`
	ResourceCalendarId *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	State              *Selection `xmlrpc:"state,omitempty"`
	StructureTypeId    *Many2One  `xmlrpc:"structure_type_id,omitempty"`
	UnderContractState *Selection `xmlrpc:"under_contract_state,omitempty"`
	Wage               *Float     `xmlrpc:"wage,omitempty"`
}

// HrContractHistorys represents array of hr.contract.history model.
type HrContractHistorys []HrContractHistory

// HrContractHistoryModel is the odoo model name.
const HrContractHistoryModel = "hr.contract.history"

// Many2One convert HrContractHistory to *Many2One.
func (hch *HrContractHistory) Many2One() *Many2One {
	return NewMany2One(hch.Id.Get(), "")
}

// CreateHrContractHistory creates a new hr.contract.history model and returns its id.
func (c *Client) CreateHrContractHistory(hch *HrContractHistory) (int64, error) {
	ids, err := c.CreateHrContractHistorys([]*HrContractHistory{hch})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrContractHistory creates a new hr.contract.history model and returns its id.
func (c *Client) CreateHrContractHistorys(hchs []*HrContractHistory) ([]int64, error) {
	var vv []interface{}
	for _, v := range hchs {
		vv = append(vv, v)
	}
	return c.Create(HrContractHistoryModel, vv, nil)
}

// UpdateHrContractHistory updates an existing hr.contract.history record.
func (c *Client) UpdateHrContractHistory(hch *HrContractHistory) error {
	return c.UpdateHrContractHistorys([]int64{hch.Id.Get()}, hch)
}

// UpdateHrContractHistorys updates existing hr.contract.history records.
// All records (represented by ids) will be updated by hch values.
func (c *Client) UpdateHrContractHistorys(ids []int64, hch *HrContractHistory) error {
	return c.Update(HrContractHistoryModel, ids, hch, nil)
}

// DeleteHrContractHistory deletes an existing hr.contract.history record.
func (c *Client) DeleteHrContractHistory(id int64) error {
	return c.DeleteHrContractHistorys([]int64{id})
}

// DeleteHrContractHistorys deletes existing hr.contract.history records.
func (c *Client) DeleteHrContractHistorys(ids []int64) error {
	return c.Delete(HrContractHistoryModel, ids)
}

// GetHrContractHistory gets hr.contract.history existing record.
func (c *Client) GetHrContractHistory(id int64) (*HrContractHistory, error) {
	hchs, err := c.GetHrContractHistorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hchs)[0]), nil
}

// GetHrContractHistorys gets hr.contract.history existing records.
func (c *Client) GetHrContractHistorys(ids []int64) (*HrContractHistorys, error) {
	hchs := &HrContractHistorys{}
	if err := c.Read(HrContractHistoryModel, ids, nil, hchs); err != nil {
		return nil, err
	}
	return hchs, nil
}

// FindHrContractHistory finds hr.contract.history record by querying it with criteria.
func (c *Client) FindHrContractHistory(criteria *Criteria) (*HrContractHistory, error) {
	hchs := &HrContractHistorys{}
	if err := c.SearchRead(HrContractHistoryModel, criteria, NewOptions().Limit(1), hchs); err != nil {
		return nil, err
	}
	return &((*hchs)[0]), nil
}

// FindHrContractHistorys finds hr.contract.history records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContractHistorys(criteria *Criteria, options *Options) (*HrContractHistorys, error) {
	hchs := &HrContractHistorys{}
	if err := c.SearchRead(HrContractHistoryModel, criteria, options, hchs); err != nil {
		return nil, err
	}
	return hchs, nil
}

// FindHrContractHistoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContractHistoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrContractHistoryModel, criteria, options)
}

// FindHrContractHistoryId finds record id by querying it with criteria.
func (c *Client) FindHrContractHistoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrContractHistoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
