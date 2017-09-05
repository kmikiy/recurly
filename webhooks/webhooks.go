package webhooks

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/kmikiy/recurly"
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

// Usage represents the usage object sent in webhooks.
type Usage struct {
	XMLName            xml.Name          `xml:"usage,omitempty"`
	ID                 int               `xml:"id,omitempty"`
	SubscriptionUUID   string            `xml:"subscription_id,omitempty"`
	AddOnCode          string            `xml:"add_on_code,omitempty"`
	MeasuredUnitID     int               `xml:"measured_unit_id,omitempty"`
	Amount             int               `xml:"amount,omitempty"`
	MerchantTag        string            `xml:"merchant_tag,omitempty"`
	RecordingTimestamp recurly.NullTime  `xml:"recording_timestamp,omitempty"`
	UsageTimestamp     recurly.NullTime  `xml:"usage_timestamp,omitempty"`
	CreatedAt          recurly.NullTime  `xml:"created_at,omitempty"`
	ModifiedAt         recurly.NullTime  `xml:"modified_at,omitempty"`
	BilledAt           recurly.NullTime  `xml:"billed_at,omitempty"`
	UsageType          string            `xml:"usage_type,omitempty"`
	UnitAmountInCents  recurly.NullInt   `xml:"unit_amount_in_cents,omitempty"`
	UsagePercentage    recurly.NullFloat `xml:"usage_percentage,omitempty"`
}

// Usage represents the usage object sent in webhooks.
type GiftCard struct {
	XMLName              xml.Name         `xml:"gift_card,omitempty"`
	RedemptionCode       string           `xml:"redemption_code,omitempty"`
	ID                   int              `xml:"id,omitempty"`
	ProductCode          string           `xml:"product_code,omitempty"`
	UnitAmountInCents    recurly.NullInt  `xml:"unit_amount_in_cents,omitempty"`
	Currency             string           `xml:"currency,omitempty"`
	GifterAccountCode    string           `xml:"gifter_account_code,omitempty"`
	RecipientAccountCode string           `xml:"recipient_account_code,omitempty"`
	InvoiceNumber        int              `xml:"invoice_number,omitempty"`
	Delivery             GiftCardDelivery `xml:"delivery,omitempty"`
	CreatedAt            recurly.NullTime `xml:"created_at,omitempty"`
	UpdatedAt            recurly.NullTime `xml:"updated_at,omitempty"`
	DeliveredAt          recurly.NullTime `xml:"delivered_at,omitempty"`
	RedeemedAt           recurly.NullTime `xml:"redeemed_at,omitempty"`
	CanceledAt           recurly.NullTime `xml:"canceled_at,omitempty"`
}

type GiftCardDelivery struct {
	Method          string                  `xml:"method,omitempty"`
	EmailAddress    string                  `xml:"email_address,omitempty"`
	DeliverAt       recurly.NullTime        `xml:"deliver_at,omitempty"`
	FirstName       string                  `xml:"first_name,omitempty"`
	LastName        string                  `xml:"last_name,omitempty"`
	Address         GiftCardDeliveryAddress `xml:"address,omitempty"`
	GifterName      string                  `xml:"gifter_name"`
	PersonalMessage string                  `xml:"personal_message"`
}

type GiftCardDeliveryAddress struct {
	Address  string `xml:"address1"`
	Address2 string `xml:"address2"`
	City     string `xml:"city"`
	State    string `xml:"state"`
	ZIP      string `xml:"zip"`
	Country  string `xml:"country"`
	Phone    string `xml:"phone"`
}

// Transaction constants.
const (
	TransactionFailureTypeDeclined  = "declined"
	TransactionFailureTypeDuplicate = "duplicate_transaction"
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
	case AccountNotificationNewXMLName:
		dst = &AccountNotificationNew{}
	case AccountNotificationUpdatedXMLName:
		dst = &AccountNotificationUpdated{}
	case AccountNotificationCanceledXMLName:
		dst = &AccountNotificationCanceled{}
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
	case SubscriptionNotificationNewXMLName:
		dst = &SubscriptionNotificationNew{}
	case SubscriptionNotificationUpdatedXMLName:
		dst = &SubscriptionNotificationUpdated{}
	case SubscriptionNotificationCanceledXMLName:
		dst = &SubscriptionNotificationCanceled{}
	case SubscriptionNotificationExpiredXMLName:
		dst = &SubscriptionNotificationExpired{}
	case SubscriptionNotificationRenewedXMLName:
		dst = &SubscriptionNotificationRenewed{}
	case SubscriptionNotificationReactivatedXMLName:
		dst = &SubscriptionNotificationReactivated{}

	// Usage notifications
	case UsageNotificationNewUsageXMLName:
		dst = &UsageNotificationNewUsage{}

	// Gift Card notifications
	case GiftCardNotificationPurchasedXMLName:
		dst = &GiftCardNotificationPurchased{}
	case GiftCardNotificationCanceledXMLName:
		dst = &GiftCardNotificationCanceled{}
	case GiftCardNotificationUpdatedXMLName:
		dst = &GiftCardNotificationUpdated{}
	case GiftCardNotificationRegeneratedXMLName:
		dst = &GiftCardNotificationRegenerated{}
	case GiftCardNotificationRedeemedXMLName:
		dst = &GiftCardNotificationRedeemed{}
	case GiftCardNotificationUpdatedBalanceXMLName:
		dst = &GiftCardNotificationUpdatedBalance{}

	// Invoice notifications
	case InvoiceNotificationNewXMLName:
		dst = &InvoiceNotificationNew{}
	case InvoiceNotificationPastDueXMLName:
		dst = &InvoiceNotificationPastDue{}

	// Payment notifications
	case PaymentNotificationSuccessfulXMLName:
		dst = &PaymentNotificationSuccessful{}
	case PaymentNotificationFailedXMLName:
		dst = &PaymentNotificationFailed{}
	case PaymentNotificationVoidXMLName:
		dst = &PaymentNotificationVoid{}
	case PaymentNotificationSuccessfulRefundXMLName:
		dst = &PaymentNotificationSuccessfulRefund{}
	default:
		return nil, ErrUnknownNotification{name: n.XMLName.Local}
	}

	if err := xml.Unmarshal(notification, dst); err != nil {
		return nil, err
	}

	return dst, nil
}
