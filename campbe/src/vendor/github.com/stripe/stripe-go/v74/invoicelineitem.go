//
//
// File generated from our OpenAPI spec
//
//

package stripe

// A string identifying the type of the source of this line item, either an `invoiceitem` or a `subscription`.
type InvoiceLineItemType string

// List of values that InvoiceLineItemType can take
const (
	InvoiceLineItemTypeInvoiceItem  InvoiceLineItemType = "invoiceitem"
	InvoiceLineItemTypeSubscription InvoiceLineItemType = "subscription"
)

// The amount of discount calculated per discount for this line item.
type InvoiceLineItemDiscountAmount struct {
	// The amount, in cents (or local equivalent), of the discount.
	Amount int64 `json:"amount"`
	// The discount that was applied to get this discount amount.
	Discount *Discount `json:"discount"`
}

// For a credit proration `line_item`, the original debit line_items to which the credit proration applies.
type InvoiceLineItemProrationDetailsCreditedItems struct {
	// Invoice containing the credited invoice line items
	Invoice string `json:"invoice"`
	// Credited invoice line items
	InvoiceLineItems []string `json:"invoice_line_items"`
}

// Additional details for proration line items
type InvoiceLineItemProrationDetails struct {
	// For a credit proration `line_item`, the original debit line_items to which the credit proration applies.
	CreditedItems *InvoiceLineItemProrationDetailsCreditedItems `json:"credited_items"`
}
type InvoiceLineItem struct {
	// The amount, in cents (or local equivalent).
	Amount int64 `json:"amount"`
	// The integer amount in cents (or local equivalent) representing the amount for this line item, excluding all tax and discounts.
	AmountExcludingTax int64 `json:"amount_excluding_tax"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// If true, discounts will apply to this line item. Always false for prorations.
	Discountable bool `json:"discountable"`
	// The amount of discount calculated per discount for this line item.
	DiscountAmounts []*InvoiceLineItemDiscountAmount `json:"discount_amounts"`
	// The discounts applied to the invoice line item. Line item discounts are applied before invoice discounts. Use `expand[]=discounts` to expand each discount.
	Discounts []*Discount `json:"discounts"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The ID of the [invoice item](https://stripe.com/docs/api/invoiceitems) associated with this line item if any.
	InvoiceItem string `json:"invoice_item"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Note that for line items with `type=subscription` this will reflect the metadata of the subscription that caused the line item to be created.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string  `json:"object"`
	Period *Period `json:"period"`
	// The plan of the subscription, if the line item is a subscription or a proration.
	Plan *Plan `json:"plan"`
	// The price of the line item.
	Price *Price `json:"price"`
	// Whether this is a proration.
	Proration bool `json:"proration"`
	// Additional details for proration line items
	ProrationDetails *InvoiceLineItemProrationDetails `json:"proration_details"`
	// The quantity of the subscription, if the line item is a subscription or a proration.
	Quantity int64 `json:"quantity"`
	// The subscription that the invoice item pertains to, if any.
	Subscription string `json:"subscription"`
	// The subscription item that generated this line item. Left empty if the line item is not an explicit result of a subscription.
	SubscriptionItem string `json:"subscription_item"`
	// The amount of tax calculated per tax rate for this line item
	TaxAmounts []*InvoiceTotalTaxAmount `json:"tax_amounts"`
	// The tax rates which apply to the line item.
	TaxRates []*TaxRate `json:"tax_rates"`
	// A string identifying the type of the source of this line item, either an `invoiceitem` or a `subscription`.
	Type InvoiceLineItemType `json:"type"`
	// The amount in cents (or local equivalent) representing the unit amount for this line item, excluding all tax and discounts.
	UnitAmountExcludingTax float64 `json:"unit_amount_excluding_tax,string"`
}

// Period is a structure representing a start and end dates.
type Period struct {
	End   int64 `json:"end"`
	Start int64 `json:"start"`
}

// InvoiceLineItemList is a list of InvoiceLineItems as retrieved from a list endpoint.
type InvoiceLineItemList struct {
	APIResource
	ListMeta
	Data []*InvoiceLineItem `json:"data"`
}
