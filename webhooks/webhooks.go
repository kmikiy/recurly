package webhooks

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/kmikiy/recurly"
)

// Webhook notification constants.
const (
	// Invoice notifications.
	NewInvoice     = "new_invoice_notification"
	PastDueInvoice = "past_due_invoice_notification"

	// Payment notifications.
	SuccessfulPayment = "successful_payment_notification"
	FailedPayment     = "failed_payment_notification"
	VoidPayment       = "void_payment_notification"
	SuccessfulRefund  = "successful_refund_notification"
)

type notificationName struct {
	XMLName xml.Name
}

// Account represents the account object sent in webhooks.
type Account struct {
	XMLName     xml.Name `xml:"account"`
	Code        string   `xml:"account_code,omitempty"`
	Username    string   `xml:"username,omitempty"`
	Email       string   `xml:"email,omitempty"`
	FirstName   string   `xml:"first_name,omitempty"`
	LastName    string   `xml:"last_name,omitempty"`
	CompanyName string   `xml:"company_name,omitempty"`
	Phone       string   `xml:"phone,omitempty"`
}

// ShippingAddress represents the shipping_address object sent in webhooks.
type ShippingAddress struct {
	XMLName     xml.Name `xml:"shipping_address"`
	ID          int      `xml:"id,omitempty"`
	Nickname    string   `xml:"nickname,omitempty"`
	FirstName   string   `xml:"first_name,omitempty"`
	LastName    string   `xml:"last_name,omitempty"`
	CompanyName string   `xml:"company_name,omitempty"`
	VATNumber   string   `xml:"vat_number,omitempty"`
	Street      string   `xml:"street1,omitempty"`
	Street2     string   `xml:"street2,omitempty"`
	City        string   `xml:"city,omitempty"`
	State       string   `xml:"state,omitempty"`
	ZIP         string   `xml:"zip,omitempty"`
	Country     string   `xml:"country,omitempty"`
	Email       string   `xml:"email,omitempty"`
	Phone       string   `xml:"phone,omitempty"`
}

// Transaction represents the transaction object sent in webhooks.
type Transaction struct {
	XMLName           xml.Name         `xml:"transaction"`
	UUID              string           `xml:"id,omitempty"`
	InvoiceNumber     int              `xml:"invoice_number,omitempty"`
	SubscriptionUUID  string           `xml:"subscription_id,omitempty"`
	Action            string           `xml:"action,omitempty"`
	AmountInCents     int              `xml:"amount_in_cents,omitempty"`
	Status            string           `xml:"status,omitempty"`
	Message           string           `xml:"message,omitempty"`
	GatewayErrorCodes string           `xml:"gateway_error_codes,omitempty"`
	FailureType       string           `xml:"failure_type,omitempty"`
	Reference         string           `xml:"reference,omitempty"`
	Source            string           `xml:"source,omitempty"`
	Test              recurly.NullBool `xml:"test,omitempty"`
	Voidable          recurly.NullBool `xml:"voidable,omitempty"`
	Refundable        recurly.NullBool `xml:"refundable,omitempty"`
}

// Invoice represents the invoice object sent in webhooks.
type Invoice struct {
	XMLName             xml.Name         `xml:"invoice,omitempty"`
	SubscriptionUUID    string           `xml:"subscription_id,omitempty"`
	UUID                string           `xml:"uuid,omitempty"`
	State               string           `xml:"state,omitempty"`
	InvoiceNumberPrefix string           `xml:"invoice_number_prefix,omitempty"`
	InvoiceNumber       int              `xml:"invoice_number,omitempty"`
	PONumber            string           `xml:"po_number,omitempty"`
	VATNumber           string           `xml:"vat_number,omitempty"`
	TotalInCents        int              `xml:"total_in_cents,omitempty"`
	Currency            string           `xml:"currency,omitempty"`
	CreatedAt           recurly.NullTime `xml:"date,omitempty"`
	ClosedAt            recurly.NullTime `xml:"closed_at,omitempty"`
	NetTerms            recurly.NullInt  `xml:"net_terms,omitempty"`
	CollectionMethod    string           `xml:"collection_method,omitempty"`
}

// Transaction constants.
const (
	TransactionFailureTypeDeclined  = "declined"
	TransactionFailureTypeDuplicate = "duplicate_transaction"
)

// Invoice types.
type (
	// NewInvoiceNotification is sent when an invoice generated.
	// https://dev.recurly.com/page/webhooks#section-new-invoice
	NewInvoiceNotification struct {
		Account Account `xml:"account"`
		Invoice Invoice `xml:"invoice"`
	}

	// PastDueInvoiceNotification is sent when an invoice is past due.
	// https://dev.recurly.com/v2.4/page/webhooks#section-past-due-invoice
	PastDueInvoiceNotification struct {
		Account Account `xml:"account"`
		Invoice Invoice `xml:"invoice"`
	}
)

// Payment types.
type (
	// SuccessfulPaymentNotification is sent when a payment is successful.
	// https://dev.recurly.com/v2.4/page/webhooks#section-successful-payment
	SuccessfulPaymentNotification struct {
		Account     Account     `xml:"account"`
		Transaction Transaction `xml:"transaction"`
	}

	// FailedPaymentNotification is sent when a payment fails.
	// https://dev.recurly.com/v2.4/page/webhooks#section-failed-payment
	FailedPaymentNotification struct {
		Account     Account     `xml:"account"`
		Transaction Transaction `xml:"transaction"`
	}

	// VoidPaymentNotification is sent when a successful payment is voided.
	// https://dev.recurly.com/page/webhooks#section-void-payment
	VoidPaymentNotification struct {
		Account     Account     `xml:"account"`
		Transaction Transaction `xml:"transaction"`
	}

	// SuccessfulRefundNotification is sent when an amount is refunded.
	// https://dev.recurly.com/page/webhooks#section-successful-refund
	SuccessfulRefundNotification struct {
		Account     Account     `xml:"account"`
		Transaction Transaction `xml:"transaction"`
	}
)

// ErrUnknownNotification is used when the incoming webhook does not match a
// predefined notification type. It implements the error interface.
type ErrUnknownNotification struct {
	name string
}

// Error implements the error interface.
func (e ErrUnknownNotification) Error() string {
	return fmt.Sprintf("unknown notification: %s", e.name)
}

// Name returns the name of the unknown notification.
func (e ErrUnknownNotification) Name() string {
	return e.name
}

// Parse parses an incoming webhook and returns the notification.
func Parse(r io.Reader) (interface{}, error) {
	if closer, ok := r.(io.Closer); ok {
		defer closer.Close()
	}

	notification, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var n notificationName
	if err := xml.Unmarshal(notification, &n); err != nil {
		return nil, err
	}

	var dst interface{}
	switch n.XMLName.Local {
	// Account notifications
	case AccountNotificationNewAccountXMLName:
		dst = &AccountNotificationNewAccount{}
	case AccountNotificationUpdatedAccountXMLName:
		dst = &AccountNotificationUpdatedAccount{}
	case AccountNotificationCanceledAccountXMLName:
		dst = &AccountNotificationCanceledAccount{}
	case AccountNotificationBillingInfoUpdatedXMLName:
		dst = &AccountNotificationBillingInfoUpdated{}
	case AccountNotificationBillingInfoUpdateFailedXMLName:
		dst = &AccountNotificationBillingInfoUpdateFailed{}
	case AccountNotificationNewShippingAddressXMLName:
		dst = &AccountNotificationNewShippingAddress{}
	case AccountNotificationUpdatedShippingAddressXMLName:
		dst = &AccountNotificationUpdatedShippingAddress{}
	case AccountNotificationDeletedShippingAddressXMLName:
		dst = &AccountNotificationDeletedShippingAddress{}

	// Subscription notifications
	case SubscriptionNotificationNewSubscriptionXMLName:
		dst = &SubscriptionNotificationNewSubscription{}
	case SubscriptionNotificationUpdatedSubscriptionXMLName:
		dst = &SubscriptionNotificationUpdatedSubscription{}
	case SubscriptionNotificationCanceledSubscriptionXMLName:
		dst = &SubscriptionNotificationCanceledSubscription{}
	case SubscriptionNotificationExpiredSubscriptionXMLName:
		dst = &SubscriptionNotificationExpiredSubscription{}
	case SubscriptionNotificationRenewedSubscriptionXMLName:
		dst = &SubscriptionNotificationRenewedSubscription{}
	case SubscriptionNotificationReactivatedSubscriptionXMLName:
		dst = &SubscriptionNotificationReactivatedSubscription{}

	case NewInvoice:
		dst = &NewInvoiceNotification{}
	case PastDueInvoice:
		dst = &PastDueInvoiceNotification{}
	case SuccessfulPayment:
		dst = &SuccessfulPaymentNotification{}
	case FailedPayment:
		dst = &FailedPaymentNotification{}
	case VoidPayment:
		dst = &VoidPaymentNotification{}
	case SuccessfulRefund:
		dst = &SuccessfulRefundNotification{}
	default:
		return nil, ErrUnknownNotification{name: n.XMLName.Local}
	}

	if err := xml.Unmarshal(notification, dst); err != nil {
		return nil, err
	}

	return dst, nil
}
