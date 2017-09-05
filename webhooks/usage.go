package webhooks

const (
	// Usage notifications.
	UsageNotificationNewUsageXMLName = "new_usage_notification"
)

// Usage types.
type (
	// UsageNotificationNewUsage is sent when a new subscription is created.
	// https://dev.recurly.com/page/webhooks#section-new-subscription
	UsageNotificationNewUsage struct {
		Account Account `xml:"account"`
		Usage   Usage   `xml:"usage"`
	}
)
