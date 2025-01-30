package odoo

// HrJobPlatform represents hr.job.platform model.
type HrJobPlatform struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Email       *String   `xmlrpc:"email,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Regex       *String   `xmlrpc:"regex,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrJobPlatforms represents array of hr.job.platform model.
type HrJobPlatforms []HrJobPlatform

// HrJobPlatformModel is the odoo model name.
const HrJobPlatformModel = "hr.job.platform"

// Many2One convert HrJobPlatform to *Many2One.
func (hjp *HrJobPlatform) Many2One() *Many2One {
	return NewMany2One(hjp.Id.Get(), "")
}

// CreateHrJobPlatform creates a new hr.job.platform model and returns its id.
func (c *Client) CreateHrJobPlatform(hjp *HrJobPlatform) (int64, error) {
	ids, err := c.CreateHrJobPlatforms([]*HrJobPlatform{hjp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrJobPlatform creates a new hr.job.platform model and returns its id.
func (c *Client) CreateHrJobPlatforms(hjps []*HrJobPlatform) ([]int64, error) {
	var vv []interface{}
	for _, v := range hjps {
		vv = append(vv, v)
	}
	return c.Create(HrJobPlatformModel, vv, nil)
}

// UpdateHrJobPlatform updates an existing hr.job.platform record.
func (c *Client) UpdateHrJobPlatform(hjp *HrJobPlatform) error {
	return c.UpdateHrJobPlatforms([]int64{hjp.Id.Get()}, hjp)
}

// UpdateHrJobPlatforms updates existing hr.job.platform records.
// All records (represented by ids) will be updated by hjp values.
func (c *Client) UpdateHrJobPlatforms(ids []int64, hjp *HrJobPlatform) error {
	return c.Update(HrJobPlatformModel, ids, hjp, nil)
}

// DeleteHrJobPlatform deletes an existing hr.job.platform record.
func (c *Client) DeleteHrJobPlatform(id int64) error {
	return c.DeleteHrJobPlatforms([]int64{id})
}

// DeleteHrJobPlatforms deletes existing hr.job.platform records.
func (c *Client) DeleteHrJobPlatforms(ids []int64) error {
	return c.Delete(HrJobPlatformModel, ids)
}

// GetHrJobPlatform gets hr.job.platform existing record.
func (c *Client) GetHrJobPlatform(id int64) (*HrJobPlatform, error) {
	hjps, err := c.GetHrJobPlatforms([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hjps)[0]), nil
}

// GetHrJobPlatforms gets hr.job.platform existing records.
func (c *Client) GetHrJobPlatforms(ids []int64) (*HrJobPlatforms, error) {
	hjps := &HrJobPlatforms{}
	if err := c.Read(HrJobPlatformModel, ids, nil, hjps); err != nil {
		return nil, err
	}
	return hjps, nil
}

// FindHrJobPlatform finds hr.job.platform record by querying it with criteria.
func (c *Client) FindHrJobPlatform(criteria *Criteria) (*HrJobPlatform, error) {
	hjps := &HrJobPlatforms{}
	if err := c.SearchRead(HrJobPlatformModel, criteria, NewOptions().Limit(1), hjps); err != nil {
		return nil, err
	}
	return &((*hjps)[0]), nil
}

// FindHrJobPlatforms finds hr.job.platform records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrJobPlatforms(criteria *Criteria, options *Options) (*HrJobPlatforms, error) {
	hjps := &HrJobPlatforms{}
	if err := c.SearchRead(HrJobPlatformModel, criteria, options, hjps); err != nil {
		return nil, err
	}
	return hjps, nil
}

// FindHrJobPlatformIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrJobPlatformIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrJobPlatformModel, criteria, options)
}

// FindHrJobPlatformId finds record id by querying it with criteria.
func (c *Client) FindHrJobPlatformId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrJobPlatformModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
