package webhooks

const (
	// Invoice notifications.
	InvoiceNotificationNewXMLName     = "new_invoice_notification"
	InvoiceNotificationPastDueXMLName = "past_due_invoice_notification"
	InvoiceNotificationClosedXMLName  = "closed_invoice_notification"
)

// Invoice types.
type (
	// InvoiceNotificationNew is sent when an invoice generated.
	// https://dev.recurly.com/page/webhooks#section-new-invoice
	InvoiceNotificationNew struct {
		Account Account `xml:"account"`
		Invoice Invoice `xml:"invoice"`
	}

	// InvoiceNotificationPastDue is sent when an invoice is past due.
	// https://dev.recurly.com/v2.4/page/webhooks#section-past-due-invoice
	InvoiceNotificationPastDue struct {
		Account Account `xml:"account"`
		Invoice Invoice `xml:"invoice"`
	}

	// InvoiceNotificationClosed is sent if an invoice is closed.
	// https://dev.recurly.com/page/webhooks#section-closed-invoice
	InvoiceNotificationClosed struct {
		Account Account `xml:"account"`
		Invoice Invoice `xml:"invoice"`
	}
)
