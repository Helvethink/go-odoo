package odoo

// CandidateSendMail represents candidate.send.mail model.
type CandidateSendMail struct {
	AttachmentIds        *Relation `xmlrpc:"attachment_ids,omitempty"`
	AuthorId             *Many2One `xmlrpc:"author_id,omitempty"`
	Body                 *String   `xmlrpc:"body,omitempty"`
	BodyHasTemplateValue *Bool     `xmlrpc:"body_has_template_value,omitempty"`
	CanEditBody          *Bool     `xmlrpc:"can_edit_body,omitempty"`
	CandidateIds         *Relation `xmlrpc:"candidate_ids,omitempty"`
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

// CandidateSendMails represents array of candidate.send.mail model.
type CandidateSendMails []CandidateSendMail

// CandidateSendMailModel is the odoo model name.
const CandidateSendMailModel = "candidate.send.mail"

// Many2One convert CandidateSendMail to *Many2One.
func (csm *CandidateSendMail) Many2One() *Many2One {
	return NewMany2One(csm.Id.Get(), "")
}

// CreateCandidateSendMail creates a new candidate.send.mail model and returns its id.
func (c *Client) CreateCandidateSendMail(csm *CandidateSendMail) (int64, error) {
	ids, err := c.CreateCandidateSendMails([]*CandidateSendMail{csm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCandidateSendMail creates a new candidate.send.mail model and returns its id.
func (c *Client) CreateCandidateSendMails(csms []*CandidateSendMail) ([]int64, error) {
	var vv []interface{}
	for _, v := range csms {
		vv = append(vv, v)
	}
	return c.Create(CandidateSendMailModel, vv, nil)
}

// UpdateCandidateSendMail updates an existing candidate.send.mail record.
func (c *Client) UpdateCandidateSendMail(csm *CandidateSendMail) error {
	return c.UpdateCandidateSendMails([]int64{csm.Id.Get()}, csm)
}

// UpdateCandidateSendMails updates existing candidate.send.mail records.
// All records (represented by ids) will be updated by csm values.
func (c *Client) UpdateCandidateSendMails(ids []int64, csm *CandidateSendMail) error {
	return c.Update(CandidateSendMailModel, ids, csm, nil)
}

// DeleteCandidateSendMail deletes an existing candidate.send.mail record.
func (c *Client) DeleteCandidateSendMail(id int64) error {
	return c.DeleteCandidateSendMails([]int64{id})
}

// DeleteCandidateSendMails deletes existing candidate.send.mail records.
func (c *Client) DeleteCandidateSendMails(ids []int64) error {
	return c.Delete(CandidateSendMailModel, ids)
}

// GetCandidateSendMail gets candidate.send.mail existing record.
func (c *Client) GetCandidateSendMail(id int64) (*CandidateSendMail, error) {
	csms, err := c.GetCandidateSendMails([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*csms)[0]), nil
}

// GetCandidateSendMails gets candidate.send.mail existing records.
func (c *Client) GetCandidateSendMails(ids []int64) (*CandidateSendMails, error) {
	csms := &CandidateSendMails{}
	if err := c.Read(CandidateSendMailModel, ids, nil, csms); err != nil {
		return nil, err
	}
	return csms, nil
}

// FindCandidateSendMail finds candidate.send.mail record by querying it with criteria.
func (c *Client) FindCandidateSendMail(criteria *Criteria) (*CandidateSendMail, error) {
	csms := &CandidateSendMails{}
	if err := c.SearchRead(CandidateSendMailModel, criteria, NewOptions().Limit(1), csms); err != nil {
		return nil, err
	}
	return &((*csms)[0]), nil
}

// FindCandidateSendMails finds candidate.send.mail records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCandidateSendMails(criteria *Criteria, options *Options) (*CandidateSendMails, error) {
	csms := &CandidateSendMails{}
	if err := c.SearchRead(CandidateSendMailModel, criteria, options, csms); err != nil {
		return nil, err
	}
	return csms, nil
}

// FindCandidateSendMailIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCandidateSendMailIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CandidateSendMailModel, criteria, options)
}

// FindCandidateSendMailId finds record id by querying it with criteria.
func (c *Client) FindCandidateSendMailId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CandidateSendMailModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
