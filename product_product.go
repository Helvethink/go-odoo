package odoo

// ProductProduct represents product.product model.
type ProductProduct struct {
	AccountTagIds                        *Relation   `xmlrpc:"account_tag_ids,omitempty"`
	Active                               *Bool       `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId              *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline                 *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration          *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                          *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState                        *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                      *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                     *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                       *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                       *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	AdditionalProductTagIds              *Relation   `xmlrpc:"additional_product_tag_ids,omitempty"`
	AllProductTagIds                     *Relation   `xmlrpc:"all_product_tag_ids,omitempty"`
	AttributeLineIds                     *Relation   `xmlrpc:"attribute_line_ids,omitempty"`
	Barcode                              *String     `xmlrpc:"barcode,omitempty"`
	CanBeExpensed                        *Bool       `xmlrpc:"can_be_expensed,omitempty"`
	CanImage1024BeZoomed                 *Bool       `xmlrpc:"can_image_1024_be_zoomed,omitempty"`
	CanImageVariant1024BeZoomed          *Bool       `xmlrpc:"can_image_variant_1024_be_zoomed,omitempty"`
	CategId                              *Many2One   `xmlrpc:"categ_id,omitempty"`
	Code                                 *String     `xmlrpc:"code,omitempty"`
	Color                                *Int        `xmlrpc:"color,omitempty"`
	CombinationIndices                   *String     `xmlrpc:"combination_indices,omitempty"`
	ComboIds                             *Relation   `xmlrpc:"combo_ids,omitempty"`
	CompanyId                            *Many2One   `xmlrpc:"company_id,omitempty"`
	CostCurrencyId                       *Many2One   `xmlrpc:"cost_currency_id,omitempty"`
	CreateDate                           *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                            *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                           *Many2One   `xmlrpc:"currency_id,omitempty"`
	DefaultCode                          *String     `xmlrpc:"default_code,omitempty"`
	Description                          *String     `xmlrpc:"description,omitempty"`
	DescriptionPurchase                  *String     `xmlrpc:"description_purchase,omitempty"`
	DescriptionSale                      *String     `xmlrpc:"description_sale,omitempty"`
	DisplayName                          *String     `xmlrpc:"display_name,omitempty"`
	ExpensePolicy                        *Selection  `xmlrpc:"expense_policy,omitempty"`
	ExpensePolicyTooltip                 *String     `xmlrpc:"expense_policy_tooltip,omitempty"`
	FiscalCountryCodes                   *String     `xmlrpc:"fiscal_country_codes,omitempty"`
	HasConfigurableAttributes            *Bool       `xmlrpc:"has_configurable_attributes,omitempty"`
	HasMessage                           *Bool       `xmlrpc:"has_message,omitempty"`
	Id                                   *Int        `xmlrpc:"id,omitempty"`
	Image1024                            *String     `xmlrpc:"image_1024,omitempty"`
	Image128                             *String     `xmlrpc:"image_128,omitempty"`
	Image1920                            *String     `xmlrpc:"image_1920,omitempty"`
	Image256                             *String     `xmlrpc:"image_256,omitempty"`
	Image512                             *String     `xmlrpc:"image_512,omitempty"`
	ImageVariant1024                     *String     `xmlrpc:"image_variant_1024,omitempty"`
	ImageVariant128                      *String     `xmlrpc:"image_variant_128,omitempty"`
	ImageVariant1920                     *String     `xmlrpc:"image_variant_1920,omitempty"`
	ImageVariant256                      *String     `xmlrpc:"image_variant_256,omitempty"`
	ImageVariant512                      *String     `xmlrpc:"image_variant_512,omitempty"`
	InvoicePolicy                        *Selection  `xmlrpc:"invoice_policy,omitempty"`
	IsFavorite                           *Bool       `xmlrpc:"is_favorite,omitempty"`
	IsInPurchaseOrder                    *Bool       `xmlrpc:"is_in_purchase_order,omitempty"`
	IsProductVariant                     *Bool       `xmlrpc:"is_product_variant,omitempty"`
	ListPrice                            *Float      `xmlrpc:"list_price,omitempty"`
	LstPrice                             *Float      `xmlrpc:"lst_price,omitempty"`
	MessageAttachmentCount               *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds                   *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                      *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter               *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                   *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                           *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                    *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction                    *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter             *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                    *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline               *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                 *String     `xmlrpc:"name,omitempty"`
	OptionalProductIds                   *Relation   `xmlrpc:"optional_product_ids,omitempty"`
	PackagingIds                         *Relation   `xmlrpc:"packaging_ids,omitempty"`
	PartnerRef                           *String     `xmlrpc:"partner_ref,omitempty"`
	PriceExtra                           *Float      `xmlrpc:"price_extra,omitempty"`
	PricelistItemCount                   *Int        `xmlrpc:"pricelist_item_count,omitempty"`
	ProductCatalogProductIsInSaleOrder   *Bool       `xmlrpc:"product_catalog_product_is_in_sale_order,omitempty"`
	ProductDocumentCount                 *Int        `xmlrpc:"product_document_count,omitempty"`
	ProductDocumentIds                   *Relation   `xmlrpc:"product_document_ids,omitempty"`
	ProductProperties                    interface{} `xmlrpc:"product_properties,omitempty"`
	ProductTagIds                        *Relation   `xmlrpc:"product_tag_ids,omitempty"`
	ProductTemplateAttributeValueIds     *Relation   `xmlrpc:"product_template_attribute_value_ids,omitempty"`
	ProductTemplateVariantValueIds       *Relation   `xmlrpc:"product_template_variant_value_ids,omitempty"`
	ProductTmplId                        *Many2One   `xmlrpc:"product_tmpl_id,omitempty"`
	ProductTooltip                       *String     `xmlrpc:"product_tooltip,omitempty"`
	ProductVariantCount                  *Int        `xmlrpc:"product_variant_count,omitempty"`
	ProductVariantId                     *Many2One   `xmlrpc:"product_variant_id,omitempty"`
	ProductVariantIds                    *Relation   `xmlrpc:"product_variant_ids,omitempty"`
	ProjectId                            *Many2One   `xmlrpc:"project_id,omitempty"`
	ProjectTemplateId                    *Many2One   `xmlrpc:"project_template_id,omitempty"`
	PropertyAccountExpenseId             *Many2One   `xmlrpc:"property_account_expense_id,omitempty"`
	PropertyAccountIncomeId              *Many2One   `xmlrpc:"property_account_income_id,omitempty"`
	PurchaseLineWarn                     *Selection  `xmlrpc:"purchase_line_warn,omitempty"`
	PurchaseLineWarnMsg                  *String     `xmlrpc:"purchase_line_warn_msg,omitempty"`
	PurchaseMethod                       *Selection  `xmlrpc:"purchase_method,omitempty"`
	PurchaseOk                           *Bool       `xmlrpc:"purchase_ok,omitempty"`
	PurchasedProductQty                  *Float      `xmlrpc:"purchased_product_qty,omitempty"`
	RatingIds                            *Relation   `xmlrpc:"rating_ids,omitempty"`
	SaleLineWarn                         *Selection  `xmlrpc:"sale_line_warn,omitempty"`
	SaleLineWarnMsg                      *String     `xmlrpc:"sale_line_warn_msg,omitempty"`
	SaleOk                               *Bool       `xmlrpc:"sale_ok,omitempty"`
	SalesCount                           *Float      `xmlrpc:"sales_count,omitempty"`
	SellerIds                            *Relation   `xmlrpc:"seller_ids,omitempty"`
	Sequence                             *Int        `xmlrpc:"sequence,omitempty"`
	ServicePolicy                        *Selection  `xmlrpc:"service_policy,omitempty"`
	ServiceToPurchase                    *Bool       `xmlrpc:"service_to_purchase,omitempty"`
	ServiceTracking                      *Selection  `xmlrpc:"service_tracking,omitempty"`
	ServiceType                          *Selection  `xmlrpc:"service_type,omitempty"`
	ServiceUpsellThreshold               *Float      `xmlrpc:"service_upsell_threshold,omitempty"`
	ServiceUpsellThresholdRatio          *String     `xmlrpc:"service_upsell_threshold_ratio,omitempty"`
	StandardPrice                        *Float      `xmlrpc:"standard_price,omitempty"`
	StandardPriceUpdateWarning           *String     `xmlrpc:"standard_price_update_warning,omitempty"`
	SupplierTaxesId                      *Relation   `xmlrpc:"supplier_taxes_id,omitempty"`
	TaxString                            *String     `xmlrpc:"tax_string,omitempty"`
	TaxesId                              *Relation   `xmlrpc:"taxes_id,omitempty"`
	Type                                 *Selection  `xmlrpc:"type,omitempty"`
	UomId                                *Many2One   `xmlrpc:"uom_id,omitempty"`
	UomName                              *String     `xmlrpc:"uom_name,omitempty"`
	UomPoId                              *Many2One   `xmlrpc:"uom_po_id,omitempty"`
	ValidProductTemplateAttributeLineIds *Relation   `xmlrpc:"valid_product_template_attribute_line_ids,omitempty"`
	VariantSellerIds                     *Relation   `xmlrpc:"variant_seller_ids,omitempty"`
	VisibleExpensePolicy                 *Bool       `xmlrpc:"visible_expense_policy,omitempty"`
	Volume                               *Float      `xmlrpc:"volume,omitempty"`
	VolumeUomName                        *String     `xmlrpc:"volume_uom_name,omitempty"`
	WebsiteMessageIds                    *Relation   `xmlrpc:"website_message_ids,omitempty"`
	Weight                               *Float      `xmlrpc:"weight,omitempty"`
	WeightUomName                        *String     `xmlrpc:"weight_uom_name,omitempty"`
	WriteDate                            *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                             *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// ProductProducts represents array of product.product model.
type ProductProducts []ProductProduct

// ProductProductModel is the odoo model name.
const ProductProductModel = "product.product"

// Many2One convert ProductProduct to *Many2One.
func (pp *ProductProduct) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductProduct creates a new product.product model and returns its id.
func (c *Client) CreateProductProduct(pp *ProductProduct) (int64, error) {
	ids, err := c.CreateProductProducts([]*ProductProduct{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductProduct creates a new product.product model and returns its id.
func (c *Client) CreateProductProducts(pps []*ProductProduct) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(ProductProductModel, vv, nil)
}

// UpdateProductProduct updates an existing product.product record.
func (c *Client) UpdateProductProduct(pp *ProductProduct) error {
	return c.UpdateProductProducts([]int64{pp.Id.Get()}, pp)
}

// UpdateProductProducts updates existing product.product records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductProducts(ids []int64, pp *ProductProduct) error {
	return c.Update(ProductProductModel, ids, pp, nil)
}

// DeleteProductProduct deletes an existing product.product record.
func (c *Client) DeleteProductProduct(id int64) error {
	return c.DeleteProductProducts([]int64{id})
}

// DeleteProductProducts deletes existing product.product records.
func (c *Client) DeleteProductProducts(ids []int64) error {
	return c.Delete(ProductProductModel, ids)
}

// GetProductProduct gets product.product existing record.
func (c *Client) GetProductProduct(id int64) (*ProductProduct, error) {
	pps, err := c.GetProductProducts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// GetProductProducts gets product.product existing records.
func (c *Client) GetProductProducts(ids []int64) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.Read(ProductProductModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProduct finds product.product record by querying it with criteria.
func (c *Client) FindProductProduct(criteria *Criteria) (*ProductProduct, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// FindProductProducts finds product.product records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProducts(criteria *Criteria, options *Options) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProductIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProductIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductProductModel, criteria, options)
}

// FindProductProductId finds record id by querying it with criteria.
func (c *Client) FindProductProductId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
