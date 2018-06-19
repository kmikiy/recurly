package webhooks

import "github.com/kmikiy/recurly"

const (
	// Dunning event notifications.
	DunningEventNotificationNewXMLName = "new_dunning_event_notification"
)

// Dunning event types.
type (
	// This notification will be sent according to your dunning configuration in Recurly.
	//https://dev.recurly.com/page/webhooks#dunning-event-notifications
	DunningEventNotificationNew struct {
		Account      Account              `xml:"account"`
		Invoice      Invoice              `xml:"invoice"`
		Subscription recurly.Subscription `xml:"subscription"`
		Transaction  Transaction          `xml:"transaction"`
	}
)
