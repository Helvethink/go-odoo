package odoo

// HrJob represents hr.job model.
type HrJob struct {
	Active                        *Bool       `xmlrpc:"active,omitempty"`
	ActivitiesOverdue             *Int        `xmlrpc:"activities_overdue,omitempty"`
	ActivitiesToday               *Int        `xmlrpc:"activities_today,omitempty"`
	AddressId                     *Many2One   `xmlrpc:"address_id,omitempty"`
	AliasBouncedContent           *String     `xmlrpc:"alias_bounced_content,omitempty"`
	AliasContact                  *Selection  `xmlrpc:"alias_contact,omitempty"`
	AliasDefaults                 *String     `xmlrpc:"alias_defaults,omitempty"`
	AliasDomain                   *String     `xmlrpc:"alias_domain,omitempty"`
	AliasDomainId                 *Many2One   `xmlrpc:"alias_domain_id,omitempty"`
	AliasEmail                    *String     `xmlrpc:"alias_email,omitempty"`
	AliasForceThreadId            *Int        `xmlrpc:"alias_force_thread_id,omitempty"`
	AliasFullName                 *String     `xmlrpc:"alias_full_name,omitempty"`
	AliasId                       *Many2One   `xmlrpc:"alias_id,omitempty"`
	AliasIncomingLocal            *Bool       `xmlrpc:"alias_incoming_local,omitempty"`
	AliasModelId                  *Many2One   `xmlrpc:"alias_model_id,omitempty"`
	AliasName                     *String     `xmlrpc:"alias_name,omitempty"`
	AliasParentModelId            *Many2One   `xmlrpc:"alias_parent_model_id,omitempty"`
	AliasParentThreadId           *Int        `xmlrpc:"alias_parent_thread_id,omitempty"`
	AliasStatus                   *Selection  `xmlrpc:"alias_status,omitempty"`
	AllApplicationCount           *Int        `xmlrpc:"all_application_count,omitempty"`
	ApplicantHired                *Int        `xmlrpc:"applicant_hired,omitempty"`
	ApplicantPropertiesDefinition interface{} `xmlrpc:"applicant_properties_definition,omitempty"`
	ApplicationCount              *Int        `xmlrpc:"application_count,omitempty"`
	ApplicationIds                *Relation   `xmlrpc:"application_ids,omitempty"`
	Color                         *Int        `xmlrpc:"color,omitempty"`
	CompanyId                     *Many2One   `xmlrpc:"company_id,omitempty"`
	ContractTypeId                *Many2One   `xmlrpc:"contract_type_id,omitempty"`
	CreateDate                    *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                     *Many2One   `xmlrpc:"create_uid,omitempty"`
	DateFrom                      *Time       `xmlrpc:"date_from,omitempty"`
	DateTo                        *Time       `xmlrpc:"date_to,omitempty"`
	DepartmentId                  *Many2One   `xmlrpc:"department_id,omitempty"`
	Description                   *String     `xmlrpc:"description,omitempty"`
	DisplayName                   *String     `xmlrpc:"display_name,omitempty"`
	DocumentIds                   *Relation   `xmlrpc:"document_ids,omitempty"`
	DocumentsCount                *Int        `xmlrpc:"documents_count,omitempty"`
	EmployeeIds                   *Relation   `xmlrpc:"employee_ids,omitempty"`
	ExpectedEmployees             *Int        `xmlrpc:"expected_employees,omitempty"`
	ExtendedInterviewerIds        *Relation   `xmlrpc:"extended_interviewer_ids,omitempty"`
	FavoriteUserIds               *Relation   `xmlrpc:"favorite_user_ids,omitempty"`
	HasMessage                    *Bool       `xmlrpc:"has_message,omitempty"`
	Id                            *Int        `xmlrpc:"id,omitempty"`
	IndustryId                    *Many2One   `xmlrpc:"industry_id,omitempty"`
	InterviewerIds                *Relation   `xmlrpc:"interviewer_ids,omitempty"`
	IsFavorite                    *Bool       `xmlrpc:"is_favorite,omitempty"`
	JobProperties                 interface{} `xmlrpc:"job_properties,omitempty"`
	ManagerId                     *Many2One   `xmlrpc:"manager_id,omitempty"`
	MessageAttachmentCount        *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds            *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError               *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter        *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError            *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                    *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower             *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction             *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter      *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds             *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	Name                          *String     `xmlrpc:"name,omitempty"`
	NewApplicationCount           *Int        `xmlrpc:"new_application_count,omitempty"`
	NoOfEmployee                  *Int        `xmlrpc:"no_of_employee,omitempty"`
	NoOfHiredEmployee             *Int        `xmlrpc:"no_of_hired_employee,omitempty"`
	NoOfRecruitment               *Int        `xmlrpc:"no_of_recruitment,omitempty"`
	OldApplicationCount           *Int        `xmlrpc:"old_application_count,omitempty"`
	RatingIds                     *Relation   `xmlrpc:"rating_ids,omitempty"`
	Requirements                  *String     `xmlrpc:"requirements,omitempty"`
	Sequence                      *Int        `xmlrpc:"sequence,omitempty"`
	SkillIds                      *Relation   `xmlrpc:"skill_ids,omitempty"`
	UserId                        *Many2One   `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds             *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                     *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                      *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// HrJobs represents array of hr.job model.
type HrJobs []HrJob

// HrJobModel is the odoo model name.
const HrJobModel = "hr.job"

// Many2One convert HrJob to *Many2One.
func (hj *HrJob) Many2One() *Many2One {
	return NewMany2One(hj.Id.Get(), "")
}

// CreateHrJob creates a new hr.job model and returns its id.
func (c *Client) CreateHrJob(hj *HrJob) (int64, error) {
	ids, err := c.CreateHrJobs([]*HrJob{hj})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrJob creates a new hr.job model and returns its id.
func (c *Client) CreateHrJobs(hjs []*HrJob) ([]int64, error) {
	var vv []interface{}
	for _, v := range hjs {
		vv = append(vv, v)
	}
	return c.Create(HrJobModel, vv, nil)
}

// UpdateHrJob updates an existing hr.job record.
func (c *Client) UpdateHrJob(hj *HrJob) error {
	return c.UpdateHrJobs([]int64{hj.Id.Get()}, hj)
}

// UpdateHrJobs updates existing hr.job records.
// All records (represented by ids) will be updated by hj values.
func (c *Client) UpdateHrJobs(ids []int64, hj *HrJob) error {
	return c.Update(HrJobModel, ids, hj, nil)
}

// DeleteHrJob deletes an existing hr.job record.
func (c *Client) DeleteHrJob(id int64) error {
	return c.DeleteHrJobs([]int64{id})
}

// DeleteHrJobs deletes existing hr.job records.
func (c *Client) DeleteHrJobs(ids []int64) error {
	return c.Delete(HrJobModel, ids)
}

// GetHrJob gets hr.job existing record.
func (c *Client) GetHrJob(id int64) (*HrJob, error) {
	hjs, err := c.GetHrJobs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hjs)[0]), nil
}

// GetHrJobs gets hr.job existing records.
func (c *Client) GetHrJobs(ids []int64) (*HrJobs, error) {
	hjs := &HrJobs{}
	if err := c.Read(HrJobModel, ids, nil, hjs); err != nil {
		return nil, err
	}
	return hjs, nil
}

// FindHrJob finds hr.job record by querying it with criteria.
func (c *Client) FindHrJob(criteria *Criteria) (*HrJob, error) {
	hjs := &HrJobs{}
	if err := c.SearchRead(HrJobModel, criteria, NewOptions().Limit(1), hjs); err != nil {
		return nil, err
	}
	return &((*hjs)[0]), nil
}

// FindHrJobs finds hr.job records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrJobs(criteria *Criteria, options *Options) (*HrJobs, error) {
	hjs := &HrJobs{}
	if err := c.SearchRead(HrJobModel, criteria, options, hjs); err != nil {
		return nil, err
	}
	return hjs, nil
}

// FindHrJobIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrJobIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrJobModel, criteria, options)
}

// FindHrJobId finds record id by querying it with criteria.
func (c *Client) FindHrJobId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrJobModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
