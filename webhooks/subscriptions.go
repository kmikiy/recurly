package webhooks

import "github.com/kmikiy/recurly"

const (
	// Subscription notifications.
	SubscriptionNotificationNewSubscriptionXMLName         = "new_subscription_notification"
	SubscriptionNotificationUpdatedSubscriptionXMLName     = "updated_subscription_notification"
	SubscriptionNotificationCanceledSubscriptionXMLName    = "canceled_subscription_notification"
	SubscriptionNotificationExpiredSubscriptionXMLName     = "expired_subscription_notification"
	SubscriptionNotificationRenewedSubscriptionXMLName     = "renewed_subscription_notification"
	SubscriptionNotificationReactivatedSubscriptionXMLName = "reactivated_account_notification"
)

// Subscription types.
type (
	// SubscriptionNotificationNewSubscription is sent when a new subscription is created.
	// https://dev.recurly.com/page/webhooks#section-new-subscription
	SubscriptionNotificationNewSubscription struct {
		Account      Account              `xml:"account"`
		Subscription recurly.Subscription `xml:"subscription"`
	}

	// SubscriptionNotificationUpdatedSubscription is sent when a subscription is upgraded or downgraded.
	// https://dev.recurly.com/page/webhooks#section-updated-subscription
	SubscriptionNotificationUpdatedSubscription struct {
		Account      Account              `xml:"account"`
		Subscription recurly.Subscription `xml:"subscription"`
	}

	// SubscriptionNotificationCanceledSubscription is sent when a subscription is canceled.
	// https://dev.recurly.com/page/webhooks#section-canceled-subscription
	SubscriptionNotificationCanceledSubscription struct {
		Account      Account              `xml:"account"`
		Subscription recurly.Subscription `xml:"subscription"`
	}

	// SubscriptionNotificationExpiredSubscription is sent when a subscription is no longer valid.
	// https://dev.recurly.com/v2.4/page/webhooks#section-expired-subscription
	SubscriptionNotificationExpiredSubscription struct {
		Account      Account              `xml:"account"`
		Subscription recurly.Subscription `xml:"subscription"`
	}

	// SubscriptionNotificationRenewedSubscription is sent when a subscription renew.
	// https://dev.recurly.com/page/webhooks#section-renewed-subscription
	SubscriptionNotificationRenewedSubscription struct {
		Account      Account              `xml:"account"`
		Subscription recurly.Subscription `xml:"subscription"`
	}

	// SubscriptionNotificationReactivatedSubscription is sent when a subscription is reactivated after having been canceled.
	// https://dev.recurly.com/page/webhooks#section-reactivated-subscription
	SubscriptionNotificationReactivatedSubscription struct {
		Account      Account              `xml:"account"`
		Subscription recurly.Subscription `xml:"subscription"`
	}
)
