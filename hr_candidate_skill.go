package odoo

// HrCandidateSkill represents hr.candidate.skill model.
type HrCandidateSkill struct {
	CandidateId   *Many2One `xmlrpc:"candidate_id,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	LevelProgress *Int      `xmlrpc:"level_progress,omitempty"`
	SkillId       *Many2One `xmlrpc:"skill_id,omitempty"`
	SkillLevelId  *Many2One `xmlrpc:"skill_level_id,omitempty"`
	SkillTypeId   *Many2One `xmlrpc:"skill_type_id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrCandidateSkills represents array of hr.candidate.skill model.
type HrCandidateSkills []HrCandidateSkill

// HrCandidateSkillModel is the odoo model name.
const HrCandidateSkillModel = "hr.candidate.skill"

// Many2One convert HrCandidateSkill to *Many2One.
func (hcs *HrCandidateSkill) Many2One() *Many2One {
	return NewMany2One(hcs.Id.Get(), "")
}

// CreateHrCandidateSkill creates a new hr.candidate.skill model and returns its id.
func (c *Client) CreateHrCandidateSkill(hcs *HrCandidateSkill) (int64, error) {
	ids, err := c.CreateHrCandidateSkills([]*HrCandidateSkill{hcs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrCandidateSkill creates a new hr.candidate.skill model and returns its id.
func (c *Client) CreateHrCandidateSkills(hcss []*HrCandidateSkill) ([]int64, error) {
	var vv []interface{}
	for _, v := range hcss {
		vv = append(vv, v)
	}
	return c.Create(HrCandidateSkillModel, vv, nil)
}

// UpdateHrCandidateSkill updates an existing hr.candidate.skill record.
func (c *Client) UpdateHrCandidateSkill(hcs *HrCandidateSkill) error {
	return c.UpdateHrCandidateSkills([]int64{hcs.Id.Get()}, hcs)
}

// UpdateHrCandidateSkills updates existing hr.candidate.skill records.
// All records (represented by ids) will be updated by hcs values.
func (c *Client) UpdateHrCandidateSkills(ids []int64, hcs *HrCandidateSkill) error {
	return c.Update(HrCandidateSkillModel, ids, hcs, nil)
}

// DeleteHrCandidateSkill deletes an existing hr.candidate.skill record.
func (c *Client) DeleteHrCandidateSkill(id int64) error {
	return c.DeleteHrCandidateSkills([]int64{id})
}

// DeleteHrCandidateSkills deletes existing hr.candidate.skill records.
func (c *Client) DeleteHrCandidateSkills(ids []int64) error {
	return c.Delete(HrCandidateSkillModel, ids)
}

// GetHrCandidateSkill gets hr.candidate.skill existing record.
func (c *Client) GetHrCandidateSkill(id int64) (*HrCandidateSkill, error) {
	hcss, err := c.GetHrCandidateSkills([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hcss)[0]), nil
}

// GetHrCandidateSkills gets hr.candidate.skill existing records.
func (c *Client) GetHrCandidateSkills(ids []int64) (*HrCandidateSkills, error) {
	hcss := &HrCandidateSkills{}
	if err := c.Read(HrCandidateSkillModel, ids, nil, hcss); err != nil {
		return nil, err
	}
	return hcss, nil
}

// FindHrCandidateSkill finds hr.candidate.skill record by querying it with criteria.
func (c *Client) FindHrCandidateSkill(criteria *Criteria) (*HrCandidateSkill, error) {
	hcss := &HrCandidateSkills{}
	if err := c.SearchRead(HrCandidateSkillModel, criteria, NewOptions().Limit(1), hcss); err != nil {
		return nil, err
	}
	return &((*hcss)[0]), nil
}

// FindHrCandidateSkills finds hr.candidate.skill records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrCandidateSkills(criteria *Criteria, options *Options) (*HrCandidateSkills, error) {
	hcss := &HrCandidateSkills{}
	if err := c.SearchRead(HrCandidateSkillModel, criteria, options, hcss); err != nil {
		return nil, err
	}
	return hcss, nil
}

// FindHrCandidateSkillIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrCandidateSkillIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrCandidateSkillModel, criteria, options)
}

// FindHrCandidateSkillId finds record id by querying it with criteria.
func (c *Client) FindHrCandidateSkillId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrCandidateSkillModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
