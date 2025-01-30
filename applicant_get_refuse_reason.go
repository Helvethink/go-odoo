package odoo

// ApplicantGetRefuseReason represents applicant.get.refuse.reason model.
type ApplicantGetRefuseReason struct {
	ApplicantEmails       *String   `xmlrpc:"applicant_emails,omitempty"`
	ApplicantIds          *Relation `xmlrpc:"applicant_ids,omitempty"`
	ApplicantWithoutEmail *String   `xmlrpc:"applicant_without_email,omitempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	Duplicates            *Bool     `xmlrpc:"duplicates,omitempty"`
	DuplicatesCount       *Int      `xmlrpc:"duplicates_count,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	RefuseReasonId        *Many2One `xmlrpc:"refuse_reason_id,omitempty"`
	SendMail              *Bool     `xmlrpc:"send_mail,omitempty"`
	SingleApplicantEmail  *String   `xmlrpc:"single_applicant_email,omitempty"`
	TemplateId            *Many2One `xmlrpc:"template_id,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ApplicantGetRefuseReasons represents array of applicant.get.refuse.reason model.
type ApplicantGetRefuseReasons []ApplicantGetRefuseReason

// ApplicantGetRefuseReasonModel is the odoo model name.
const ApplicantGetRefuseReasonModel = "applicant.get.refuse.reason"

// Many2One convert ApplicantGetRefuseReason to *Many2One.
func (agrr *ApplicantGetRefuseReason) Many2One() *Many2One {
	return NewMany2One(agrr.Id.Get(), "")
}

// CreateApplicantGetRefuseReason creates a new applicant.get.refuse.reason model and returns its id.
func (c *Client) CreateApplicantGetRefuseReason(agrr *ApplicantGetRefuseReason) (int64, error) {
	ids, err := c.CreateApplicantGetRefuseReasons([]*ApplicantGetRefuseReason{agrr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateApplicantGetRefuseReason creates a new applicant.get.refuse.reason model and returns its id.
func (c *Client) CreateApplicantGetRefuseReasons(agrrs []*ApplicantGetRefuseReason) ([]int64, error) {
	var vv []interface{}
	for _, v := range agrrs {
		vv = append(vv, v)
	}
	return c.Create(ApplicantGetRefuseReasonModel, vv, nil)
}

// UpdateApplicantGetRefuseReason updates an existing applicant.get.refuse.reason record.
func (c *Client) UpdateApplicantGetRefuseReason(agrr *ApplicantGetRefuseReason) error {
	return c.UpdateApplicantGetRefuseReasons([]int64{agrr.Id.Get()}, agrr)
}

// UpdateApplicantGetRefuseReasons updates existing applicant.get.refuse.reason records.
// All records (represented by ids) will be updated by agrr values.
func (c *Client) UpdateApplicantGetRefuseReasons(ids []int64, agrr *ApplicantGetRefuseReason) error {
	return c.Update(ApplicantGetRefuseReasonModel, ids, agrr, nil)
}

// DeleteApplicantGetRefuseReason deletes an existing applicant.get.refuse.reason record.
func (c *Client) DeleteApplicantGetRefuseReason(id int64) error {
	return c.DeleteApplicantGetRefuseReasons([]int64{id})
}

// DeleteApplicantGetRefuseReasons deletes existing applicant.get.refuse.reason records.
func (c *Client) DeleteApplicantGetRefuseReasons(ids []int64) error {
	return c.Delete(ApplicantGetRefuseReasonModel, ids)
}

// GetApplicantGetRefuseReason gets applicant.get.refuse.reason existing record.
func (c *Client) GetApplicantGetRefuseReason(id int64) (*ApplicantGetRefuseReason, error) {
	agrrs, err := c.GetApplicantGetRefuseReasons([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*agrrs)[0]), nil
}

// GetApplicantGetRefuseReasons gets applicant.get.refuse.reason existing records.
func (c *Client) GetApplicantGetRefuseReasons(ids []int64) (*ApplicantGetRefuseReasons, error) {
	agrrs := &ApplicantGetRefuseReasons{}
	if err := c.Read(ApplicantGetRefuseReasonModel, ids, nil, agrrs); err != nil {
		return nil, err
	}
	return agrrs, nil
}

// FindApplicantGetRefuseReason finds applicant.get.refuse.reason record by querying it with criteria.
func (c *Client) FindApplicantGetRefuseReason(criteria *Criteria) (*ApplicantGetRefuseReason, error) {
	agrrs := &ApplicantGetRefuseReasons{}
	if err := c.SearchRead(ApplicantGetRefuseReasonModel, criteria, NewOptions().Limit(1), agrrs); err != nil {
		return nil, err
	}
	return &((*agrrs)[0]), nil
}

// FindApplicantGetRefuseReasons finds applicant.get.refuse.reason records by querying it
// and filtering it with criteria and options.
func (c *Client) FindApplicantGetRefuseReasons(criteria *Criteria, options *Options) (*ApplicantGetRefuseReasons, error) {
	agrrs := &ApplicantGetRefuseReasons{}
	if err := c.SearchRead(ApplicantGetRefuseReasonModel, criteria, options, agrrs); err != nil {
		return nil, err
	}
	return agrrs, nil
}

// FindApplicantGetRefuseReasonIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindApplicantGetRefuseReasonIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ApplicantGetRefuseReasonModel, criteria, options)
}

// FindApplicantGetRefuseReasonId finds record id by querying it with criteria.
func (c *Client) FindApplicantGetRefuseReasonId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ApplicantGetRefuseReasonModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
