package odoo

// ApplicantSendMail represents applicant.send.mail model.
type ApplicantSendMail struct {
	ApplicantIds         *Relation `xmlrpc:"applicant_ids,omitempty"`
	AttachmentIds        *Relation `xmlrpc:"attachment_ids,omitempty"`
	AuthorId             *Many2One `xmlrpc:"author_id,omitempty"`
	Body                 *String   `xmlrpc:"body,omitempty"`
	BodyHasTemplateValue *Bool     `xmlrpc:"body_has_template_value,omitempty"`
	CanEditBody          *Bool     `xmlrpc:"can_edit_body,omitempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	IsMailTemplateEditor *Bool     `xmlrpc:"is_mail_template_editor,omitempty"`
	Lang                 *String   `xmlrpc:"lang,omitempty"`
	RenderModel          *String   `xmlrpc:"render_model,omitempty"`
	Subject              *String   `xmlrpc:"subject,omitempty"`
	TemplateId           *Many2One `xmlrpc:"template_id,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ApplicantSendMails represents array of applicant.send.mail model.
type ApplicantSendMails []ApplicantSendMail

// ApplicantSendMailModel is the odoo model name.
const ApplicantSendMailModel = "applicant.send.mail"

// Many2One convert ApplicantSendMail to *Many2One.
func (asm *ApplicantSendMail) Many2One() *Many2One {
	return NewMany2One(asm.Id.Get(), "")
}

// CreateApplicantSendMail creates a new applicant.send.mail model and returns its id.
func (c *Client) CreateApplicantSendMail(asm *ApplicantSendMail) (int64, error) {
	ids, err := c.CreateApplicantSendMails([]*ApplicantSendMail{asm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateApplicantSendMail creates a new applicant.send.mail model and returns its id.
func (c *Client) CreateApplicantSendMails(asms []*ApplicantSendMail) ([]int64, error) {
	var vv []interface{}
	for _, v := range asms {
		vv = append(vv, v)
	}
	return c.Create(ApplicantSendMailModel, vv, nil)
}

// UpdateApplicantSendMail updates an existing applicant.send.mail record.
func (c *Client) UpdateApplicantSendMail(asm *ApplicantSendMail) error {
	return c.UpdateApplicantSendMails([]int64{asm.Id.Get()}, asm)
}

// UpdateApplicantSendMails updates existing applicant.send.mail records.
// All records (represented by ids) will be updated by asm values.
func (c *Client) UpdateApplicantSendMails(ids []int64, asm *ApplicantSendMail) error {
	return c.Update(ApplicantSendMailModel, ids, asm, nil)
}

// DeleteApplicantSendMail deletes an existing applicant.send.mail record.
func (c *Client) DeleteApplicantSendMail(id int64) error {
	return c.DeleteApplicantSendMails([]int64{id})
}

// DeleteApplicantSendMails deletes existing applicant.send.mail records.
func (c *Client) DeleteApplicantSendMails(ids []int64) error {
	return c.Delete(ApplicantSendMailModel, ids)
}

// GetApplicantSendMail gets applicant.send.mail existing record.
func (c *Client) GetApplicantSendMail(id int64) (*ApplicantSendMail, error) {
	asms, err := c.GetApplicantSendMails([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*asms)[0]), nil
}

// GetApplicantSendMails gets applicant.send.mail existing records.
func (c *Client) GetApplicantSendMails(ids []int64) (*ApplicantSendMails, error) {
	asms := &ApplicantSendMails{}
	if err := c.Read(ApplicantSendMailModel, ids, nil, asms); err != nil {
		return nil, err
	}
	return asms, nil
}

// FindApplicantSendMail finds applicant.send.mail record by querying it with criteria.
func (c *Client) FindApplicantSendMail(criteria *Criteria) (*ApplicantSendMail, error) {
	asms := &ApplicantSendMails{}
	if err := c.SearchRead(ApplicantSendMailModel, criteria, NewOptions().Limit(1), asms); err != nil {
		return nil, err
	}
	return &((*asms)[0]), nil
}

// FindApplicantSendMails finds applicant.send.mail records by querying it
// and filtering it with criteria and options.
func (c *Client) FindApplicantSendMails(criteria *Criteria, options *Options) (*ApplicantSendMails, error) {
	asms := &ApplicantSendMails{}
	if err := c.SearchRead(ApplicantSendMailModel, criteria, options, asms); err != nil {
		return nil, err
	}
	return asms, nil
}

// FindApplicantSendMailIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindApplicantSendMailIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ApplicantSendMailModel, criteria, options)
}

// FindApplicantSendMailId finds record id by querying it with criteria.
func (c *Client) FindApplicantSendMailId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ApplicantSendMailModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
