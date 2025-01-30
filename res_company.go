package odoo

// ResCompany represents res.company model.
type ResCompany struct {
	AccountCashBasisBaseAccountId               *Many2One   `xmlrpc:"account_cash_basis_base_account_id,omitempty"`
	AccountDefaultPosReceivableAccountId        *Many2One   `xmlrpc:"account_default_pos_receivable_account_id,omitempty"`
	AccountDiscountExpenseAllocationId          *Many2One   `xmlrpc:"account_discount_expense_allocation_id,omitempty"`
	AccountDiscountIncomeAllocationId           *Many2One   `xmlrpc:"account_discount_income_allocation_id,omitempty"`
	AccountEnabledTaxCountryIds                 *Relation   `xmlrpc:"account_enabled_tax_country_ids,omitempty"`
	AccountFiscalCountryId                      *Many2One   `xmlrpc:"account_fiscal_country_id,omitempty"`
	AccountJournalEarlyPayDiscountGainAccountId *Many2One   `xmlrpc:"account_journal_early_pay_discount_gain_account_id,omitempty"`
	AccountJournalEarlyPayDiscountLossAccountId *Many2One   `xmlrpc:"account_journal_early_pay_discount_loss_account_id,omitempty"`
	AccountJournalSuspenseAccountId             *Many2One   `xmlrpc:"account_journal_suspense_account_id,omitempty"`
	AccountOpeningDate                          *Time       `xmlrpc:"account_opening_date,omitempty"`
	AccountOpeningJournalId                     *Many2One   `xmlrpc:"account_opening_journal_id,omitempty"`
	AccountOpeningMoveId                        *Many2One   `xmlrpc:"account_opening_move_id,omitempty"`
	AccountPriceInclude                         *Selection  `xmlrpc:"account_price_include,omitempty"`
	AccountPurchaseTaxId                        *Many2One   `xmlrpc:"account_purchase_tax_id,omitempty"`
	AccountSaleTaxId                            *Many2One   `xmlrpc:"account_sale_tax_id,omitempty"`
	AccountStorno                               *Bool       `xmlrpc:"account_storno,omitempty"`
	AccountUseCreditLimit                       *Bool       `xmlrpc:"account_use_credit_limit,omitempty"`
	Active                                      *Bool       `xmlrpc:"active,omitempty"`
	AliasDomainId                               *Many2One   `xmlrpc:"alias_domain_id,omitempty"`
	AliasDomainName                             *String     `xmlrpc:"alias_domain_name,omitempty"`
	AllChildIds                                 *Relation   `xmlrpc:"all_child_ids,omitempty"`
	AngloSaxonAccounting                        *Bool       `xmlrpc:"anglo_saxon_accounting,omitempty"`
	AutomaticEntryDefaultJournalId              *Many2One   `xmlrpc:"automatic_entry_default_journal_id,omitempty"`
	AutopostBills                               *Bool       `xmlrpc:"autopost_bills,omitempty"`
	BankAccountCodePrefix                       *String     `xmlrpc:"bank_account_code_prefix,omitempty"`
	BankIds                                     *Relation   `xmlrpc:"bank_ids,omitempty"`
	BankJournalIds                              *Relation   `xmlrpc:"bank_journal_ids,omitempty"`
	BatchPaymentSequenceId                      *Many2One   `xmlrpc:"batch_payment_sequence_id,omitempty"`
	BounceEmail                                 *String     `xmlrpc:"bounce_email,omitempty"`
	BounceFormatted                             *String     `xmlrpc:"bounce_formatted,omitempty"`
	CandidatePropertiesDefinition               interface{} `xmlrpc:"candidate_properties_definition,omitempty"`
	CashAccountCodePrefix                       *String     `xmlrpc:"cash_account_code_prefix,omitempty"`
	CatchallEmail                               *String     `xmlrpc:"catchall_email,omitempty"`
	CatchallFormatted                           *String     `xmlrpc:"catchall_formatted,omitempty"`
	ChartTemplate                               *Selection  `xmlrpc:"chart_template,omitempty"`
	CheckAccountAuditTrail                      *Bool       `xmlrpc:"check_account_audit_trail,omitempty"`
	ChildIds                                    *Relation   `xmlrpc:"child_ids,omitempty"`
	City                                        *String     `xmlrpc:"city,omitempty"`
	Color                                       *Int        `xmlrpc:"color,omitempty"`
	CompanyDetails                              *String     `xmlrpc:"company_details,omitempty"`
	CompanyExpenseAllowedPaymentMethodLineIds   *Relation   `xmlrpc:"company_expense_allowed_payment_method_line_ids,omitempty"`
	CompanyRegistry                             *String     `xmlrpc:"company_registry,omitempty"`
	CompanyVatPlaceholder                       *String     `xmlrpc:"company_vat_placeholder,omitempty"`
	ContractExpirationNoticePeriod              *Int        `xmlrpc:"contract_expiration_notice_period,omitempty"`
	CountryCode                                 *String     `xmlrpc:"country_code,omitempty"`
	CountryId                                   *Many2One   `xmlrpc:"country_id,omitempty"`
	CreateDate                                  *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                                   *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyExchangeJournalId                   *Many2One   `xmlrpc:"currency_exchange_journal_id,omitempty"`
	CurrencyId                                  *Many2One   `xmlrpc:"currency_id,omitempty"`
	DefaultCashDifferenceExpenseAccountId       *Many2One   `xmlrpc:"default_cash_difference_expense_account_id,omitempty"`
	DefaultCashDifferenceIncomeAccountId        *Many2One   `xmlrpc:"default_cash_difference_income_account_id,omitempty"`
	DefaultFromEmail                            *String     `xmlrpc:"default_from_email,omitempty"`
	DisplayInvoiceAmountTotalWords              *Bool       `xmlrpc:"display_invoice_amount_total_words,omitempty"`
	DisplayInvoiceTaxCompanyCurrency            *Bool       `xmlrpc:"display_invoice_tax_company_currency,omitempty"`
	DisplayName                                 *String     `xmlrpc:"display_name,omitempty"`
	Email                                       *String     `xmlrpc:"email,omitempty"`
	EmailFormatted                              *String     `xmlrpc:"email_formatted,omitempty"`
	EmailPrimaryColor                           *String     `xmlrpc:"email_primary_color,omitempty"`
	EmailSecondaryColor                         *String     `xmlrpc:"email_secondary_color,omitempty"`
	EmployeePropertiesDefinition                interface{} `xmlrpc:"employee_properties_definition,omitempty"`
	ExpectsChartOfAccounts                      *Bool       `xmlrpc:"expects_chart_of_accounts,omitempty"`
	ExpenseAccrualAccountId                     *Many2One   `xmlrpc:"expense_accrual_account_id,omitempty"`
	ExpenseCurrencyExchangeAccountId            *Many2One   `xmlrpc:"expense_currency_exchange_account_id,omitempty"`
	ExpenseJournalId                            *Many2One   `xmlrpc:"expense_journal_id,omitempty"`
	ExpenseOutstandingAccountId                 *Many2One   `xmlrpc:"expense_outstanding_account_id,omitempty"`
	ExternalReportLayoutId                      *Many2One   `xmlrpc:"external_report_layout_id,omitempty"`
	FiscalPositionIds                           *Relation   `xmlrpc:"fiscal_position_ids,omitempty"`
	FiscalyearLastDay                           *Int        `xmlrpc:"fiscalyear_last_day,omitempty"`
	FiscalyearLastMonth                         *Selection  `xmlrpc:"fiscalyear_last_month,omitempty"`
	FiscalyearLockDate                          *Time       `xmlrpc:"fiscalyear_lock_date,omitempty"`
	Font                                        *Selection  `xmlrpc:"font,omitempty"`
	HardLockDate                                *Time       `xmlrpc:"hard_lock_date,omitempty"`
	HasMessage                                  *Bool       `xmlrpc:"has_message,omitempty"`
	HrPresenceControlAttendance                 *Bool       `xmlrpc:"hr_presence_control_attendance,omitempty"`
	HrPresenceControlEmail                      *Bool       `xmlrpc:"hr_presence_control_email,omitempty"`
	HrPresenceControlEmailAmount                *Int        `xmlrpc:"hr_presence_control_email_amount,omitempty"`
	HrPresenceControlIp                         *Bool       `xmlrpc:"hr_presence_control_ip,omitempty"`
	HrPresenceControlIpList                     *String     `xmlrpc:"hr_presence_control_ip_list,omitempty"`
	HrPresenceControlLogin                      *Bool       `xmlrpc:"hr_presence_control_login,omitempty"`
	IapEnrichAutoDone                           *Bool       `xmlrpc:"iap_enrich_auto_done,omitempty"`
	Id                                          *Int        `xmlrpc:"id,omitempty"`
	IncomeCurrencyExchangeAccountId             *Many2One   `xmlrpc:"income_currency_exchange_account_id,omitempty"`
	IncotermId                                  *Many2One   `xmlrpc:"incoterm_id,omitempty"`
	InternalProjectId                           *Many2One   `xmlrpc:"internal_project_id,omitempty"`
	InvoiceTerms                                *String     `xmlrpc:"invoice_terms,omitempty"`
	InvoiceTermsHtml                            *String     `xmlrpc:"invoice_terms_html,omitempty"`
	IsCompanyDetailsEmpty                       *Bool       `xmlrpc:"is_company_details_empty,omitempty"`
	JobPropertiesDefinition                     interface{} `xmlrpc:"job_properties_definition,omitempty"`
	LayoutBackground                            *Selection  `xmlrpc:"layout_background,omitempty"`
	LayoutBackgroundImage                       *String     `xmlrpc:"layout_background_image,omitempty"`
	LeaveTimesheetTaskId                        *Many2One   `xmlrpc:"leave_timesheet_task_id,omitempty"`
	Logo                                        *String     `xmlrpc:"logo,omitempty"`
	LogoWeb                                     *String     `xmlrpc:"logo_web,omitempty"`
	MessageAttachmentCount                      *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds                          *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                             *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter                      *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                          *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                                  *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                           *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction                           *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter                    *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                           *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	Mobile                                      *String     `xmlrpc:"mobile,omitempty"`
	MultiVatForeignCountryIds                   *Relation   `xmlrpc:"multi_vat_foreign_country_ids,omitempty"`
	Name                                        *String     `xmlrpc:"name,omitempty"`
	PaperformatId                               *Many2One   `xmlrpc:"paperformat_id,omitempty"`
	ParentId                                    *Many2One   `xmlrpc:"parent_id,omitempty"`
	ParentIds                                   *Relation   `xmlrpc:"parent_ids,omitempty"`
	ParentPath                                  *String     `xmlrpc:"parent_path,omitempty"`
	PartnerGid                                  *Int        `xmlrpc:"partner_gid,omitempty"`
	PartnerId                                   *Many2One   `xmlrpc:"partner_id,omitempty"`
	PaymentOnboardingPaymentMethod              *Selection  `xmlrpc:"payment_onboarding_payment_method,omitempty"`
	Phone                                       *String     `xmlrpc:"phone,omitempty"`
	PoDoubleValidation                          *Selection  `xmlrpc:"po_double_validation,omitempty"`
	PoDoubleValidationAmount                    *Float      `xmlrpc:"po_double_validation_amount,omitempty"`
	PoLead                                      *Float      `xmlrpc:"po_lead,omitempty"`
	PoLock                                      *Selection  `xmlrpc:"po_lock,omitempty"`
	PortalConfirmationPay                       *Bool       `xmlrpc:"portal_confirmation_pay,omitempty"`
	PortalConfirmationSign                      *Bool       `xmlrpc:"portal_confirmation_sign,omitempty"`
	PrepaymentPercent                           *Float      `xmlrpc:"prepayment_percent,omitempty"`
	PrimaryColor                                *String     `xmlrpc:"primary_color,omitempty"`
	ProjectTimeModeId                           *Many2One   `xmlrpc:"project_time_mode_id,omitempty"`
	PurchaseLockDate                            *Time       `xmlrpc:"purchase_lock_date,omitempty"`
	QrCode                                      *Bool       `xmlrpc:"qr_code,omitempty"`
	QuickEditMode                               *Selection  `xmlrpc:"quick_edit_mode,omitempty"`
	QuotationValidityDays                       *Int        `xmlrpc:"quotation_validity_days,omitempty"`
	RatingIds                                   *Relation   `xmlrpc:"rating_ids,omitempty"`
	ReportFooter                                *String     `xmlrpc:"report_footer,omitempty"`
	ReportHeader                                *String     `xmlrpc:"report_header,omitempty"`
	ResourceCalendarId                          *Many2One   `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceCalendarIds                         *Relation   `xmlrpc:"resource_calendar_ids,omitempty"`
	RevenueAccrualAccountId                     *Many2One   `xmlrpc:"revenue_accrual_account_id,omitempty"`
	RootId                                      *Many2One   `xmlrpc:"root_id,omitempty"`
	SaleDiscountProductId                       *Many2One   `xmlrpc:"sale_discount_product_id,omitempty"`
	SaleLockDate                                *Time       `xmlrpc:"sale_lock_date,omitempty"`
	SaleOnboardingPaymentMethod                 *Selection  `xmlrpc:"sale_onboarding_payment_method,omitempty"`
	SaleOrderTemplateId                         *Many2One   `xmlrpc:"sale_order_template_id,omitempty"`
	SecondaryColor                              *String     `xmlrpc:"secondary_color,omitempty"`
	Sequence                                    *Int        `xmlrpc:"sequence,omitempty"`
	SnailmailColor                              *Bool       `xmlrpc:"snailmail_color,omitempty"`
	SnailmailCover                              *Bool       `xmlrpc:"snailmail_cover,omitempty"`
	SnailmailDuplex                             *Bool       `xmlrpc:"snailmail_duplex,omitempty"`
	StateId                                     *Many2One   `xmlrpc:"state_id,omitempty"`
	Street                                      *String     `xmlrpc:"street,omitempty"`
	Street2                                     *String     `xmlrpc:"street2,omitempty"`
	TaxCalculationRoundingMethod                *Selection  `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisJournalId                       *Many2One   `xmlrpc:"tax_cash_basis_journal_id,omitempty"`
	TaxExigibility                              *Bool       `xmlrpc:"tax_exigibility,omitempty"`
	TaxLockDate                                 *Time       `xmlrpc:"tax_lock_date,omitempty"`
	TermsType                                   *Selection  `xmlrpc:"terms_type,omitempty"`
	TimesheetEncodeUomId                        *Many2One   `xmlrpc:"timesheet_encode_uom_id,omitempty"`
	TransferAccountCodePrefix                   *String     `xmlrpc:"transfer_account_code_prefix,omitempty"`
	TransferAccountId                           *Many2One   `xmlrpc:"transfer_account_id,omitempty"`
	UninstalledL10NModuleIds                    *Relation   `xmlrpc:"uninstalled_l10n_module_ids,omitempty"`
	UserFiscalyearLockDate                      *Time       `xmlrpc:"user_fiscalyear_lock_date,omitempty"`
	UserHardLockDate                            *Time       `xmlrpc:"user_hard_lock_date,omitempty"`
	UserIds                                     *Relation   `xmlrpc:"user_ids,omitempty"`
	UserPurchaseLockDate                        *Time       `xmlrpc:"user_purchase_lock_date,omitempty"`
	UserSaleLockDate                            *Time       `xmlrpc:"user_sale_lock_date,omitempty"`
	UserTaxLockDate                             *Time       `xmlrpc:"user_tax_lock_date,omitempty"`
	UsesDefaultLogo                             *Bool       `xmlrpc:"uses_default_logo,omitempty"`
	Vat                                         *String     `xmlrpc:"vat,omitempty"`
	Website                                     *String     `xmlrpc:"website,omitempty"`
	WebsiteMessageIds                           *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WorkPermitExpirationNoticePeriod            *Int        `xmlrpc:"work_permit_expiration_notice_period,omitempty"`
	WriteDate                                   *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                                    *Many2One   `xmlrpc:"write_uid,omitempty"`
	Zip                                         *String     `xmlrpc:"zip,omitempty"`
}

// ResCompanys represents array of res.company model.
type ResCompanys []ResCompany

// ResCompanyModel is the odoo model name.
const ResCompanyModel = "res.company"

// Many2One convert ResCompany to *Many2One.
func (rc *ResCompany) Many2One() *Many2One {
	return NewMany2One(rc.Id.Get(), "")
}

// CreateResCompany creates a new res.company model and returns its id.
func (c *Client) CreateResCompany(rc *ResCompany) (int64, error) {
	ids, err := c.CreateResCompanys([]*ResCompany{rc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResCompany creates a new res.company model and returns its id.
func (c *Client) CreateResCompanys(rcs []*ResCompany) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcs {
		vv = append(vv, v)
	}
	return c.Create(ResCompanyModel, vv, nil)
}

// UpdateResCompany updates an existing res.company record.
func (c *Client) UpdateResCompany(rc *ResCompany) error {
	return c.UpdateResCompanys([]int64{rc.Id.Get()}, rc)
}

// UpdateResCompanys updates existing res.company records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResCompanys(ids []int64, rc *ResCompany) error {
	return c.Update(ResCompanyModel, ids, rc, nil)
}

// DeleteResCompany deletes an existing res.company record.
func (c *Client) DeleteResCompany(id int64) error {
	return c.DeleteResCompanys([]int64{id})
}

// DeleteResCompanys deletes existing res.company records.
func (c *Client) DeleteResCompanys(ids []int64) error {
	return c.Delete(ResCompanyModel, ids)
}

// GetResCompany gets res.company existing record.
func (c *Client) GetResCompany(id int64) (*ResCompany, error) {
	rcs, err := c.GetResCompanys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcs)[0]), nil
}

// GetResCompanys gets res.company existing records.
func (c *Client) GetResCompanys(ids []int64) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.Read(ResCompanyModel, ids, nil, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompany finds res.company record by querying it with criteria.
func (c *Client) FindResCompany(criteria *Criteria) (*ResCompany, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, NewOptions().Limit(1), rcs); err != nil {
		return nil, err
	}
	return &((*rcs)[0]), nil
}

// FindResCompanys finds res.company records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanys(criteria *Criteria, options *Options) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, options, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompanyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResCompanyModel, criteria, options)
}

// FindResCompanyId finds record id by querying it with criteria.
func (c *Client) FindResCompanyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
