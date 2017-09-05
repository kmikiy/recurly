package webhooks

const (
	// Payment notifications.
	PaymentNotificationSuccessfulXMLName       = "successful_payment_notification"
	PaymentNotificationFailedXMLName           = "failed_payment_notification"
	PaymentNotificationVoidXMLName             = "void_payment_notification"
	PaymentNotificationSuccessfulRefundXMLName = "successful_refund_notification"
)

// Payment types.
type (
	// PaymentNotificationSuccessful is sent when a payment is successful.
	// https://dev.recurly.com/v2.4/page/webhooks#section-successful-payment
	PaymentNotificationSuccessful struct {
		Account     Account     `xml:"account"`
		Transaction Transaction `xml:"transaction"`
	}

	// PaymentNotificationFailed is sent when a payment fails.
	// https://dev.recurly.com/v2.4/page/webhooks#section-failed-payment
	PaymentNotificationFailed struct {
		Account     Account     `xml:"account"`
		Transaction Transaction `xml:"transaction"`
	}

	// PaymentNotificationVoid is sent when a successful payment is voided.
	// https://dev.recurly.com/page/webhooks#section-void-payment
	PaymentNotificationVoid struct {
		Account     Account     `xml:"account"`
		Transaction Transaction `xml:"transaction"`
	}

	// PaymentNotificationSuccessfulRefund is sent when an amount is refunded.
	// https://dev.recurly.com/page/webhooks#section-successful-refund
	PaymentNotificationSuccessfulRefund struct {
		Account     Account     `xml:"account"`
		Transaction Transaction `xml:"transaction"`
	}
)
