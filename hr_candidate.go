package odoo

// HrCandidate represents hr.candidate model.
type HrCandidate struct {
	AcceptedApplicationsCount   *Int        `xmlrpc:"accepted_applications_count,omitempty"`
	Active                      *Bool       `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId     *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	ApplicantIds                *Relation   `xmlrpc:"applicant_ids,omitempty"`
	ApplicationCount            *Int        `xmlrpc:"application_count,omitempty"`
	ApplicationsCount           *Int        `xmlrpc:"applications_count,omitempty"`
	AttachmentCount             *Int        `xmlrpc:"attachment_count,omitempty"`
	AttachmentIds               *Relation   `xmlrpc:"attachment_ids,omitempty"`
	Availability                *Time       `xmlrpc:"availability,omitempty"`
	CandidateProperties         interface{} `xmlrpc:"candidate_properties,omitempty"`
	CandidateSkillIds           *Relation   `xmlrpc:"candidate_skill_ids,omitempty"`
	CategIds                    *Relation   `xmlrpc:"categ_ids,omitempty"`
	Color                       *Int        `xmlrpc:"color,omitempty"`
	CompanyId                   *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayName                 *String     `xmlrpc:"display_name,omitempty"`
	EmailCc                     *String     `xmlrpc:"email_cc,omitempty"`
	EmailFrom                   *String     `xmlrpc:"email_from,omitempty"`
	EmailNormalized             *String     `xmlrpc:"email_normalized,omitempty"`
	EmpIsActive                 *Bool       `xmlrpc:"emp_is_active,omitempty"`
	EmployeeId                  *Many2One   `xmlrpc:"employee_id,omitempty"`
	EmployeeName                *String     `xmlrpc:"employee_name,omitempty"`
	HasMessage                  *Bool       `xmlrpc:"has_message,omitempty"`
	Id                          *Int        `xmlrpc:"id,omitempty"`
	IsBlacklisted               *Bool       `xmlrpc:"is_blacklisted,omitempty"`
	LinkedinProfile             *String     `xmlrpc:"linkedin_profile,omitempty"`
	MatchingScore               *Float      `xmlrpc:"matching_score,omitempty"`
	MatchingSkillIds            *Relation   `xmlrpc:"matching_skill_ids,omitempty"`
	MeetingDisplayDate          *Time       `xmlrpc:"meeting_display_date,omitempty"`
	MeetingDisplayText          *String     `xmlrpc:"meeting_display_text,omitempty"`
	MeetingIds                  *Relation   `xmlrpc:"meeting_ids,omitempty"`
	MessageAttachmentCount      *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageBounce               *Int        `xmlrpc:"message_bounce,omitempty"`
	MessageFollowerIds          *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId     *Many2One   `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction           *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MissingSkillIds             *Relation   `xmlrpc:"missing_skill_ids,omitempty"`
	MobileBlacklisted           *Bool       `xmlrpc:"mobile_blacklisted,omitempty"`
	MyActivityDateDeadline      *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	PartnerId                   *Many2One   `xmlrpc:"partner_id,omitempty"`
	PartnerName                 *String     `xmlrpc:"partner_name,omitempty"`
	PartnerPhone                *String     `xmlrpc:"partner_phone,omitempty"`
	PartnerPhoneSanitized       *String     `xmlrpc:"partner_phone_sanitized,omitempty"`
	PhoneBlacklisted            *Bool       `xmlrpc:"phone_blacklisted,omitempty"`
	PhoneMobileSearch           *String     `xmlrpc:"phone_mobile_search,omitempty"`
	PhoneSanitized              *String     `xmlrpc:"phone_sanitized,omitempty"`
	PhoneSanitizedBlacklisted   *Bool       `xmlrpc:"phone_sanitized_blacklisted,omitempty"`
	Priority                    *Selection  `xmlrpc:"priority,omitempty"`
	RatingIds                   *Relation   `xmlrpc:"rating_ids,omitempty"`
	RefusedApplicationsCount    *Int        `xmlrpc:"refused_applications_count,omitempty"`
	SimilarCandidatesCount      *Int        `xmlrpc:"similar_candidates_count,omitempty"`
	SkillIds                    *Relation   `xmlrpc:"skill_ids,omitempty"`
	TypeId                      *Many2One   `xmlrpc:"type_id,omitempty"`
	UserId                      *Many2One   `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds           *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// HrCandidates represents array of hr.candidate model.
type HrCandidates []HrCandidate

// HrCandidateModel is the odoo model name.
const HrCandidateModel = "hr.candidate"

// Many2One convert HrCandidate to *Many2One.
func (hc *HrCandidate) Many2One() *Many2One {
	return NewMany2One(hc.Id.Get(), "")
}

// CreateHrCandidate creates a new hr.candidate model and returns its id.
func (c *Client) CreateHrCandidate(hc *HrCandidate) (int64, error) {
	ids, err := c.CreateHrCandidates([]*HrCandidate{hc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrCandidate creates a new hr.candidate model and returns its id.
func (c *Client) CreateHrCandidates(hcs []*HrCandidate) ([]int64, error) {
	var vv []interface{}
	for _, v := range hcs {
		vv = append(vv, v)
	}
	return c.Create(HrCandidateModel, vv, nil)
}

// UpdateHrCandidate updates an existing hr.candidate record.
func (c *Client) UpdateHrCandidate(hc *HrCandidate) error {
	return c.UpdateHrCandidates([]int64{hc.Id.Get()}, hc)
}

// UpdateHrCandidates updates existing hr.candidate records.
// All records (represented by ids) will be updated by hc values.
func (c *Client) UpdateHrCandidates(ids []int64, hc *HrCandidate) error {
	return c.Update(HrCandidateModel, ids, hc, nil)
}

// DeleteHrCandidate deletes an existing hr.candidate record.
func (c *Client) DeleteHrCandidate(id int64) error {
	return c.DeleteHrCandidates([]int64{id})
}

// DeleteHrCandidates deletes existing hr.candidate records.
func (c *Client) DeleteHrCandidates(ids []int64) error {
	return c.Delete(HrCandidateModel, ids)
}

// GetHrCandidate gets hr.candidate existing record.
func (c *Client) GetHrCandidate(id int64) (*HrCandidate, error) {
	hcs, err := c.GetHrCandidates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hcs)[0]), nil
}

// GetHrCandidates gets hr.candidate existing records.
func (c *Client) GetHrCandidates(ids []int64) (*HrCandidates, error) {
	hcs := &HrCandidates{}
	if err := c.Read(HrCandidateModel, ids, nil, hcs); err != nil {
		return nil, err
	}
	return hcs, nil
}

// FindHrCandidate finds hr.candidate record by querying it with criteria.
func (c *Client) FindHrCandidate(criteria *Criteria) (*HrCandidate, error) {
	hcs := &HrCandidates{}
	if err := c.SearchRead(HrCandidateModel, criteria, NewOptions().Limit(1), hcs); err != nil {
		return nil, err
	}
	return &((*hcs)[0]), nil
}

// FindHrCandidates finds hr.candidate records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrCandidates(criteria *Criteria, options *Options) (*HrCandidates, error) {
	hcs := &HrCandidates{}
	if err := c.SearchRead(HrCandidateModel, criteria, options, hcs); err != nil {
		return nil, err
	}
	return hcs, nil
}

// FindHrCandidateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrCandidateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrCandidateModel, criteria, options)
}

// FindHrCandidateId finds record id by querying it with criteria.
func (c *Client) FindHrCandidateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrCandidateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
