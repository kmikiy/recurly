package recurly

import (
	"encoding/xml"
	"strconv"
	"time"
)

const (
	// InvoiceStateOpen is an invoice state for invoices that are open, pending
	// collection.
	InvoiceStateOpen = "open"

	// InvoiceStateCollected is an invoice state for invoices that have been
	// successfully collected.
	InvoiceStateCollected = "collected"

	// InvoiceStateFailed is an invoice state for invoices that failed to collect.
	InvoiceStateFailed = "failed"

	// InvoiceStatePastDue is an invoice state for invoices where initial collection
	// failed, but Recurly is still attempting collection.
	InvoiceStatePastDue = "past_due"

	// CollectionMethodAutomatic is a collection method where the customer's
	// credit card is charged.
	CollectionMethodAutomatic = "automatic"

	// CollectionMethodManual is a collection method where the customer pays via
	// an accepted payment method outside of Recurly.
	CollectionMethodManual = "manual"
)

// Payment method constants.
const (
	PaymentMethodCreditCard   = "credit_card"
	PaymentMethodPayPal       = "paypal"
	PaymentMethodEFT          = "eft"
	PaymentMethodWireTransfer = "wire_transfer"
	PaymentMethodMoneyOrder   = "money_order"
	PaymentMethodCheck        = "check"
	PaymentMethodOther        = "other"
)

type InvoiceCommon struct {
	InvoiceNumberFromHref         int           `xml:"-"`
	AccountCode                   string        `xml:"-"`
	Address                       Address       `xml:"-"`
	ShippingAddress               Address       `xml:"-"`
	SubscriptionUUID              string        `xml:"-"`
	OriginalInvoiceNumber         int           `xml:"-"`
	UUID                          string        `xml:"-"`
	State                         string        `xml:"-"`
	InvoiceNumberPrefix           string        `xml:"-"`
	InvoiceNumber                 int           `xml:"-"`
	PONumber                      string        `xml:"po_number,omitempty"` // PostInvoice param
	VATNumber                     string        `xml:"-"`
	SubtotalInCents               int           `xml:"-"`
	TaxInCents                    int           `xml:"-"`
	TotalInCents                  int           `xml:"-"`
	Currency                      string        `xml:"-"`
	CreatedAt                     NullTime      `xml:"-"`
	UpdatedAt                     NullTime      `xml:"-"`
	AttemptNextCollectionAt       NullTime      `xml:"-"`
	ClosedAt                      NullTime      `xml:"-"`
	RecoveryReason                string        `xml:"-"`
	SubtotalBeforeDiscountInCents int           `xml:"-"`
	DiscountInCents               int           `xml:"-"`
	BalanceInCents                int           `xml:"-"`
	Type                          string        `xml:"-"`
	Origin                        string        `xml:"-"`
	DueOn                         NullTime      `xml:"-"`
	TaxType                       string        `xml:"-"`
	TaxRegion                     string        `xml:"-"`
	TaxRate                       float64       `xml:"-"`
	NetTerms                      NullInt       `xml:"net_terms,omitempty"`                // PostInvoice param
	CollectionMethod              string        `xml:"collection_method,omitempty"`        // PostInvoice param
	TermsAndConditions            string        `xml:"terms_and_conditions,omitempty"`     // PostInvoice param
	CustomerNotes                 string        `xml:"customer_notes,omitempty"`           // PostInvoice param
	VatReverseChargeNotes         string        `xml:"vat_reverse_charge_notes,omitempty"` // PostInvoice param
	LineItems                     []Adjustment  `xml:"-"`
	Transactions                  []Transaction `xml:"-"`
}

// Invoice is an individual invoice for an account.
// The only fields annotated with XML tags are those for posting an invoice.
// Unmarshaling an invoice is handled by the custom UnmarshalXML function.
type Invoice struct {
	XMLName xml.Name `xml:"invoice,omitempty"`
	InvoiceCommon
}

type invoiceCommon struct {
	InvoiceNumberFromHref         int           `xml:"-"`
	AccountCode                   hrefString    `xml:"account,omitempty"` // Read only
	Address                       Address       `xml:"address,omitempty"`
	ShippingAddress               Address       `xml:"shipping_address,omitempty"`
	SubscriptionUUID              hrefString    `xml:"subscription,omitempty"`
	OriginalInvoiceNumber         hrefInt       `xml:"original_invoice,omitempty"` // Read only
	UUID                          string        `xml:"uuid,omitempty"`
	State                         string        `xml:"state,omitempty"`
	InvoiceNumberPrefix           string        `xml:"invoice_number_prefix,omitempty"`
	InvoiceNumber                 int           `xml:"invoice_number,omitempty"`
	PONumber                      string        `xml:"po_number,omitempty"`
	VATNumber                     string        `xml:"vat_number,omitempty"`
	SubtotalInCents               int           `xml:"subtotal_in_cents,omitempty"`
	TaxInCents                    int           `xml:"tax_in_cents,omitempty"`
	TotalInCents                  int           `xml:"total_in_cents,omitempty"`
	Currency                      string        `xml:"currency,omitempty"`
	CreatedAt                     NullTime      `xml:"created_at,omitempty"`
	ClosedAt                      NullTime      `xml:"closed_at,omitempty"`
	TaxType                       string        `xml:"tax_type,omitempty"`
	TaxRegion                     string        `xml:"tax_region,omitempty"`
	TaxRate                       float64       `xml:"tax_rate,omitempty"`
	NetTerms                      NullInt       `xml:"net_terms,omitempty"`
	CollectionMethod              string        `xml:"collection_method,omitempty"`
	UpdatedAt                     NullTime      `xml:"updated_at,omitempty"`
	AttemptNextCollectionAt       NullTime      `xml:"attempt_next_collection_at,omitempty"`
	RecoveryReason                string        `xml:"recovery_reason,omitempty"`
	SubtotalBeforeDiscountInCents int           `xml:"subtotal_before_discount_in_cents,omitempty"`
	DiscountInCents               int           `xml:"discount_in_cents,omitempty"`
	BalanceInCents                int           `xml:"balance_in_cents,omitempty"`
	Type                          string        `xml:"type,omitempty"`
	Origin                        string        `xml:"origin,omitempty"`
	DueOn                         NullTime      `xml:"due_on,omitempty"`
	TermsAndConditions            string        `xml:"terms_and_conditions,omitempty"`     // PostInvoice param
	CustomerNotes                 string        `xml:"customer_notes,omitempty"`           // PostInvoice param
	VatReverseChargeNotes         string        `xml:"vat_reverse_charge_notes,omitempty"` // PostInvoice param
	LineItems                     []Adjustment  `xml:"line_items>adjustment,omitempty"`
	Transactions                  []Transaction `xml:"transactions>transaction,omitempty"`
}

// UnmarshalXML unmarshals invoices and handles intermediary state during unmarshaling
// for types like href.
func (i *Invoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v struct {
		XMLName xml.Name `xml:"invoice,omitempty"`
		HREF    string   `xml:"href,attr"`
		invoiceCommon
	}
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	InvoiceNumberFromHref, _ := strconv.Atoi(rxHREF.FindString(v.HREF))
	v.invoiceCommon.InvoiceNumberFromHref = InvoiceNumberFromHref

	*i = Invoice{
		XMLName:       v.XMLName,
		InvoiceCommon: convertInvoiceCommon(v.invoiceCommon),
	}
	return nil
}

func convertInvoiceCommon(ic invoiceCommon) InvoiceCommon {
	return InvoiceCommon{
		InvoiceNumberFromHref: ic.InvoiceNumberFromHref,
		AccountCode:           string(ic.AccountCode),
		Address:               ic.Address,
		ShippingAddress:       ic.ShippingAddress,
		SubscriptionUUID:      string(ic.SubscriptionUUID),
		OriginalInvoiceNumber: int(ic.OriginalInvoiceNumber),
		UUID:                          ic.UUID,
		State:                         ic.State,
		InvoiceNumberPrefix:           ic.InvoiceNumberPrefix,
		InvoiceNumber:                 ic.InvoiceNumber,
		PONumber:                      ic.PONumber,
		VATNumber:                     ic.VATNumber,
		SubtotalInCents:               ic.SubtotalInCents,
		TaxInCents:                    ic.TaxInCents,
		TotalInCents:                  ic.TotalInCents,
		Currency:                      ic.Currency,
		CreatedAt:                     ic.CreatedAt,
		UpdatedAt:                     ic.UpdatedAt,
		AttemptNextCollectionAt:       ic.AttemptNextCollectionAt,
		ClosedAt:                      ic.ClosedAt,
		RecoveryReason:                ic.RecoveryReason,
		SubtotalBeforeDiscountInCents: ic.SubtotalBeforeDiscountInCents,
		DiscountInCents:               ic.DiscountInCents,
		BalanceInCents:                ic.BalanceInCents,
		Type:                          ic.Type,
		Origin:                        ic.Origin,
		DueOn:                         ic.DueOn,
		TaxType:                       ic.TaxType,
		TaxRegion:                     ic.TaxRegion,
		TaxRate:                       ic.TaxRate,
		NetTerms:                      ic.NetTerms,
		CollectionMethod:              ic.CollectionMethod,
		TermsAndConditions:            ic.TermsAndConditions,
		CustomerNotes:                 ic.CustomerNotes,
		VatReverseChargeNotes:         ic.VatReverseChargeNotes,
		LineItems:                     ic.LineItems,
		Transactions:                  ic.Transactions,
	}
}

// OfflinePayment is a payment received outside the system to be recorded in Recurly.
type OfflinePayment struct {
	XMLName       xml.Name   `xml:"transaction"`
	InvoiceNumber int        `xml:"-"`
	PaymentMethod string     `xml:"payment_method"`
	CollectedAt   *time.Time `xml:"collected_at,omitempty"`
	Amount        int        `xml:"amount_in_cents,omitempty"`
	Description   string     `xml:"description,omitempty"`
}

// InvoiceCollection can be found in preview subscriptions
// The only fields annotated with XML tags are those for posting an invoice.
// Unmarshaling an invoice is handled by the custom UnmarshalXML function.
type InvoiceCollection struct {
	XMLName        xml.Name        `xml:"invoice_collection,omitempty"`
	ChargeInvoice  ChargeInvoice   `xml:"charge_invoice,omitempty"`
	CreditInvoices []CreditInvoice `xml:"credit_invoices>credit_invoice,omitempty"`
}

type ChargeInvoice struct {
	XMLName xml.Name `xml:"charge_invoice,omitempty"`
	InvoiceCommon
}

type CreditInvoice struct {
	XMLName xml.Name `xml:"credit_invoice,omitempty"`
	InvoiceCommon
}

// UnmarshalXML unmarshals invoices and handles intermediary state during unmarshaling
// for types like href.
func (i *ChargeInvoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v struct {
		XMLName xml.Name `xml:"charge_invoice,omitempty"`
		HREF    string   `xml:"href,attr"`
		invoiceCommon
	}
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	InvoiceNumberFromHref, _ := strconv.Atoi(rxHREF.FindString(v.HREF))
	v.invoiceCommon.InvoiceNumberFromHref = InvoiceNumberFromHref

	*i = ChargeInvoice{
		XMLName:       v.XMLName,
		InvoiceCommon: convertInvoiceCommon(v.invoiceCommon),
	}
	return nil
}

// UnmarshalXML unmarshals invoices and handles intermediary state during unmarshaling
// for types like href.
func (i *CreditInvoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v struct {
		XMLName xml.Name `xml:"credit_invoice,omitempty"`
		HREF    string   `xml:"href,attr"`
		invoiceCommon
	}
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	InvoiceNumberFromHref, _ := strconv.Atoi(rxHREF.FindString(v.HREF))
	v.invoiceCommon.InvoiceNumberFromHref = InvoiceNumberFromHref

	*i = CreditInvoice{
		XMLName:       v.XMLName,
		InvoiceCommon: convertInvoiceCommon(v.invoiceCommon),
	}
	return nil
}
