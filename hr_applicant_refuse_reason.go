package odoo

// HrApplicantRefuseReason represents hr.applicant.refuse.reason model.
type HrApplicantRefuseReason struct {
	Active      *Bool     `xmlrpc:"active,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	TemplateId  *Many2One `xmlrpc:"template_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrApplicantRefuseReasons represents array of hr.applicant.refuse.reason model.
type HrApplicantRefuseReasons []HrApplicantRefuseReason

// HrApplicantRefuseReasonModel is the odoo model name.
const HrApplicantRefuseReasonModel = "hr.applicant.refuse.reason"

// Many2One convert HrApplicantRefuseReason to *Many2One.
func (harr *HrApplicantRefuseReason) Many2One() *Many2One {
	return NewMany2One(harr.Id.Get(), "")
}

// CreateHrApplicantRefuseReason creates a new hr.applicant.refuse.reason model and returns its id.
func (c *Client) CreateHrApplicantRefuseReason(harr *HrApplicantRefuseReason) (int64, error) {
	ids, err := c.CreateHrApplicantRefuseReasons([]*HrApplicantRefuseReason{harr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrApplicantRefuseReason creates a new hr.applicant.refuse.reason model and returns its id.
func (c *Client) CreateHrApplicantRefuseReasons(harrs []*HrApplicantRefuseReason) ([]int64, error) {
	var vv []interface{}
	for _, v := range harrs {
		vv = append(vv, v)
	}
	return c.Create(HrApplicantRefuseReasonModel, vv, nil)
}

// UpdateHrApplicantRefuseReason updates an existing hr.applicant.refuse.reason record.
func (c *Client) UpdateHrApplicantRefuseReason(harr *HrApplicantRefuseReason) error {
	return c.UpdateHrApplicantRefuseReasons([]int64{harr.Id.Get()}, harr)
}

// UpdateHrApplicantRefuseReasons updates existing hr.applicant.refuse.reason records.
// All records (represented by ids) will be updated by harr values.
func (c *Client) UpdateHrApplicantRefuseReasons(ids []int64, harr *HrApplicantRefuseReason) error {
	return c.Update(HrApplicantRefuseReasonModel, ids, harr, nil)
}

// DeleteHrApplicantRefuseReason deletes an existing hr.applicant.refuse.reason record.
func (c *Client) DeleteHrApplicantRefuseReason(id int64) error {
	return c.DeleteHrApplicantRefuseReasons([]int64{id})
}

// DeleteHrApplicantRefuseReasons deletes existing hr.applicant.refuse.reason records.
func (c *Client) DeleteHrApplicantRefuseReasons(ids []int64) error {
	return c.Delete(HrApplicantRefuseReasonModel, ids)
}

// GetHrApplicantRefuseReason gets hr.applicant.refuse.reason existing record.
func (c *Client) GetHrApplicantRefuseReason(id int64) (*HrApplicantRefuseReason, error) {
	harrs, err := c.GetHrApplicantRefuseReasons([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*harrs)[0]), nil
}

// GetHrApplicantRefuseReasons gets hr.applicant.refuse.reason existing records.
func (c *Client) GetHrApplicantRefuseReasons(ids []int64) (*HrApplicantRefuseReasons, error) {
	harrs := &HrApplicantRefuseReasons{}
	if err := c.Read(HrApplicantRefuseReasonModel, ids, nil, harrs); err != nil {
		return nil, err
	}
	return harrs, nil
}

// FindHrApplicantRefuseReason finds hr.applicant.refuse.reason record by querying it with criteria.
func (c *Client) FindHrApplicantRefuseReason(criteria *Criteria) (*HrApplicantRefuseReason, error) {
	harrs := &HrApplicantRefuseReasons{}
	if err := c.SearchRead(HrApplicantRefuseReasonModel, criteria, NewOptions().Limit(1), harrs); err != nil {
		return nil, err
	}
	return &((*harrs)[0]), nil
}

// FindHrApplicantRefuseReasons finds hr.applicant.refuse.reason records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrApplicantRefuseReasons(criteria *Criteria, options *Options) (*HrApplicantRefuseReasons, error) {
	harrs := &HrApplicantRefuseReasons{}
	if err := c.SearchRead(HrApplicantRefuseReasonModel, criteria, options, harrs); err != nil {
		return nil, err
	}
	return harrs, nil
}

// FindHrApplicantRefuseReasonIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrApplicantRefuseReasonIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrApplicantRefuseReasonModel, criteria, options)
}

// FindHrApplicantRefuseReasonId finds record id by querying it with criteria.
func (c *Client) FindHrApplicantRefuseReasonId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrApplicantRefuseReasonModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
