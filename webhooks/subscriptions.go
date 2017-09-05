package webhooks

import (
	"github.com/kmikiy/recurly"
)

const (
	// Subscription notifications.
	SubscriptionNotificationNewXMLName         = "new_subscription_notification"
	SubscriptionNotificationUpdatedXMLName     = "updated_subscription_notification"
	SubscriptionNotificationCanceledXMLName    = "canceled_subscription_notification"
	SubscriptionNotificationExpiredXMLName     = "expired_subscription_notification"
	SubscriptionNotificationRenewedXMLName     = "renewed_subscription_notification"
	SubscriptionNotificationReactivatedXMLName = "reactivated_account_notification"
)

// Subscription types.
type (
	// SubscriptionNotificationNew is sent when a new subscription is created.
	// https://dev.recurly.com/page/webhooks#section-new-subscription
	SubscriptionNotificationNew struct {
		Account      Account              `xml:"account"`
		Subscription recurly.Subscription `xml:"subscription"`
	}

	// SubscriptionNotificationUpdated is sent when a subscription is upgraded or downgraded.
	// https://dev.recurly.com/page/webhooks#section-updated-subscription
	SubscriptionNotificationUpdated struct {
		Account      Account              `xml:"account"`
		Subscription recurly.Subscription `xml:"subscription"`
	}

	// SubscriptionNotificationCanceled is sent when a subscription is canceled.
	// https://dev.recurly.com/page/webhooks#section-canceled-subscription
	SubscriptionNotificationCanceled struct {
		Account      Account              `xml:"account"`
		Subscription recurly.Subscription `xml:"subscription"`
	}

	// SubscriptionNotificationExpired is sent when a subscription is no longer valid.
	// https://dev.recurly.com/v2.4/page/webhooks#section-expired-subscription
	SubscriptionNotificationExpired struct {
		Account      Account              `xml:"account"`
		Subscription recurly.Subscription `xml:"subscription"`
	}

	// SubscriptionNotificationRenewed is sent when a subscription renew.
	// https://dev.recurly.com/page/webhooks#section-renewed-subscription
	SubscriptionNotificationRenewed struct {
		Account      Account              `xml:"account"`
		Subscription recurly.Subscription `xml:"subscription"`
	}

	// SubscriptionNotificationReactivated is sent when a subscription is reactivated after having been canceled.
	// https://dev.recurly.com/page/webhooks#section-reactivated-subscription
	SubscriptionNotificationReactivated struct {
		Account      Account              `xml:"account"`
		Subscription recurly.Subscription `xml:"subscription"`
	}
)
