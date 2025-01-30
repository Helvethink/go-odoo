package odoo

// HrRecruitmentSource represents hr.recruitment.source model.
type HrRecruitmentSource struct {
	AliasId     *Many2One `xmlrpc:"alias_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Email       *String   `xmlrpc:"email,omitempty"`
	HasDomain   *String   `xmlrpc:"has_domain,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	JobId       *Many2One `xmlrpc:"job_id,omitempty"`
	MediumId    *Many2One `xmlrpc:"medium_id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	SourceId    *Many2One `xmlrpc:"source_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrRecruitmentSources represents array of hr.recruitment.source model.
type HrRecruitmentSources []HrRecruitmentSource

// HrRecruitmentSourceModel is the odoo model name.
const HrRecruitmentSourceModel = "hr.recruitment.source"

// Many2One convert HrRecruitmentSource to *Many2One.
func (hrs *HrRecruitmentSource) Many2One() *Many2One {
	return NewMany2One(hrs.Id.Get(), "")
}

// CreateHrRecruitmentSource creates a new hr.recruitment.source model and returns its id.
func (c *Client) CreateHrRecruitmentSource(hrs *HrRecruitmentSource) (int64, error) {
	ids, err := c.CreateHrRecruitmentSources([]*HrRecruitmentSource{hrs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrRecruitmentSource creates a new hr.recruitment.source model and returns its id.
func (c *Client) CreateHrRecruitmentSources(hrss []*HrRecruitmentSource) ([]int64, error) {
	var vv []interface{}
	for _, v := range hrss {
		vv = append(vv, v)
	}
	return c.Create(HrRecruitmentSourceModel, vv, nil)
}

// UpdateHrRecruitmentSource updates an existing hr.recruitment.source record.
func (c *Client) UpdateHrRecruitmentSource(hrs *HrRecruitmentSource) error {
	return c.UpdateHrRecruitmentSources([]int64{hrs.Id.Get()}, hrs)
}

// UpdateHrRecruitmentSources updates existing hr.recruitment.source records.
// All records (represented by ids) will be updated by hrs values.
func (c *Client) UpdateHrRecruitmentSources(ids []int64, hrs *HrRecruitmentSource) error {
	return c.Update(HrRecruitmentSourceModel, ids, hrs, nil)
}

// DeleteHrRecruitmentSource deletes an existing hr.recruitment.source record.
func (c *Client) DeleteHrRecruitmentSource(id int64) error {
	return c.DeleteHrRecruitmentSources([]int64{id})
}

// DeleteHrRecruitmentSources deletes existing hr.recruitment.source records.
func (c *Client) DeleteHrRecruitmentSources(ids []int64) error {
	return c.Delete(HrRecruitmentSourceModel, ids)
}

// GetHrRecruitmentSource gets hr.recruitment.source existing record.
func (c *Client) GetHrRecruitmentSource(id int64) (*HrRecruitmentSource, error) {
	hrss, err := c.GetHrRecruitmentSources([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hrss)[0]), nil
}

// GetHrRecruitmentSources gets hr.recruitment.source existing records.
func (c *Client) GetHrRecruitmentSources(ids []int64) (*HrRecruitmentSources, error) {
	hrss := &HrRecruitmentSources{}
	if err := c.Read(HrRecruitmentSourceModel, ids, nil, hrss); err != nil {
		return nil, err
	}
	return hrss, nil
}

// FindHrRecruitmentSource finds hr.recruitment.source record by querying it with criteria.
func (c *Client) FindHrRecruitmentSource(criteria *Criteria) (*HrRecruitmentSource, error) {
	hrss := &HrRecruitmentSources{}
	if err := c.SearchRead(HrRecruitmentSourceModel, criteria, NewOptions().Limit(1), hrss); err != nil {
		return nil, err
	}
	return &((*hrss)[0]), nil
}

// FindHrRecruitmentSources finds hr.recruitment.source records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRecruitmentSources(criteria *Criteria, options *Options) (*HrRecruitmentSources, error) {
	hrss := &HrRecruitmentSources{}
	if err := c.SearchRead(HrRecruitmentSourceModel, criteria, options, hrss); err != nil {
		return nil, err
	}
	return hrss, nil
}

// FindHrRecruitmentSourceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRecruitmentSourceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrRecruitmentSourceModel, criteria, options)
}

// FindHrRecruitmentSourceId finds record id by querying it with criteria.
func (c *Client) FindHrRecruitmentSourceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrRecruitmentSourceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
