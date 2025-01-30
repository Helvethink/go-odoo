package odoo

// HrApplicantCategory represents hr.applicant.category model.
type HrApplicantCategory struct {
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrApplicantCategorys represents array of hr.applicant.category model.
type HrApplicantCategorys []HrApplicantCategory

// HrApplicantCategoryModel is the odoo model name.
const HrApplicantCategoryModel = "hr.applicant.category"

// Many2One convert HrApplicantCategory to *Many2One.
func (hac *HrApplicantCategory) Many2One() *Many2One {
	return NewMany2One(hac.Id.Get(), "")
}

// CreateHrApplicantCategory creates a new hr.applicant.category model and returns its id.
func (c *Client) CreateHrApplicantCategory(hac *HrApplicantCategory) (int64, error) {
	ids, err := c.CreateHrApplicantCategorys([]*HrApplicantCategory{hac})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrApplicantCategory creates a new hr.applicant.category model and returns its id.
func (c *Client) CreateHrApplicantCategorys(hacs []*HrApplicantCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range hacs {
		vv = append(vv, v)
	}
	return c.Create(HrApplicantCategoryModel, vv, nil)
}

// UpdateHrApplicantCategory updates an existing hr.applicant.category record.
func (c *Client) UpdateHrApplicantCategory(hac *HrApplicantCategory) error {
	return c.UpdateHrApplicantCategorys([]int64{hac.Id.Get()}, hac)
}

// UpdateHrApplicantCategorys updates existing hr.applicant.category records.
// All records (represented by ids) will be updated by hac values.
func (c *Client) UpdateHrApplicantCategorys(ids []int64, hac *HrApplicantCategory) error {
	return c.Update(HrApplicantCategoryModel, ids, hac, nil)
}

// DeleteHrApplicantCategory deletes an existing hr.applicant.category record.
func (c *Client) DeleteHrApplicantCategory(id int64) error {
	return c.DeleteHrApplicantCategorys([]int64{id})
}

// DeleteHrApplicantCategorys deletes existing hr.applicant.category records.
func (c *Client) DeleteHrApplicantCategorys(ids []int64) error {
	return c.Delete(HrApplicantCategoryModel, ids)
}

// GetHrApplicantCategory gets hr.applicant.category existing record.
func (c *Client) GetHrApplicantCategory(id int64) (*HrApplicantCategory, error) {
	hacs, err := c.GetHrApplicantCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hacs)[0]), nil
}

// GetHrApplicantCategorys gets hr.applicant.category existing records.
func (c *Client) GetHrApplicantCategorys(ids []int64) (*HrApplicantCategorys, error) {
	hacs := &HrApplicantCategorys{}
	if err := c.Read(HrApplicantCategoryModel, ids, nil, hacs); err != nil {
		return nil, err
	}
	return hacs, nil
}

// FindHrApplicantCategory finds hr.applicant.category record by querying it with criteria.
func (c *Client) FindHrApplicantCategory(criteria *Criteria) (*HrApplicantCategory, error) {
	hacs := &HrApplicantCategorys{}
	if err := c.SearchRead(HrApplicantCategoryModel, criteria, NewOptions().Limit(1), hacs); err != nil {
		return nil, err
	}
	return &((*hacs)[0]), nil
}

// FindHrApplicantCategorys finds hr.applicant.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrApplicantCategorys(criteria *Criteria, options *Options) (*HrApplicantCategorys, error) {
	hacs := &HrApplicantCategorys{}
	if err := c.SearchRead(HrApplicantCategoryModel, criteria, options, hacs); err != nil {
		return nil, err
	}
	return hacs, nil
}

// FindHrApplicantCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrApplicantCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrApplicantCategoryModel, criteria, options)
}

// FindHrApplicantCategoryId finds record id by querying it with criteria.
func (c *Client) FindHrApplicantCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrApplicantCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
