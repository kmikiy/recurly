package webhooks

const (
	// Gift Card notifications.
	GiftCardNotificationPurchasedXMLName      = "purchased_gift_card_notification"
	GiftCardNotificationCanceledXMLName       = "canceled_gift_card_notification"
	GiftCardNotificationUpdatedXMLName        = "updated_gift_card_notification"
	GiftCardNotificationRegeneratedXMLName    = "regenerated_gift_card_notification"
	GiftCardNotificationRedeemedXMLName       = "redeemed_gift_card_notification"
	GiftCardNotificationUpdatedBalanceXMLName = "updated_balance_gift_card_notification"
)

// Subscription types.
type (
	// GiftCardNotificationPurchased is sent when a gift card is purchased by a gifter.
	// https://dev.recurly.com/page/webhooks#section-purchased-gift-card
	GiftCardNotificationPurchased struct {
		GiftCard GiftCard `xml:"gift_card,omitempty"`
	}

	// GiftCardNotificationCanceled is sent when you cancel a gift card from the Admin Console.
	// https://dev.recurly.com/page/webhooks#section-canceled-gift-card
	GiftCardNotificationCanceled struct {
		GiftCard GiftCard `xml:"gift_card,omitempty"`
	}

	// GiftCardNotificationUpdated is sent when you edit a gift card's delivery information from the Admin Console.
	// https://dev.recurly.com/page/webhooks#section-updated-gift-card
	GiftCardNotificationUpdated struct {
		GiftCard GiftCard `xml:"gift_card,omitempty"`
	}

	// GiftCardNotificationRegenerated is sent when you regenerate a gift card's redemption code from the Admin Console.
	// https://dev.recurly.com/page/webhooks#section-regenerated-gift-card
	GiftCardNotificationRegenerated struct {
		GiftCard GiftCard `xml:"gift_card,omitempty"`
	}

	// GiftCardNotificationRedeemed is sent when a gift card is redeemed by a recipient.
	// https://dev.recurly.com/page/webhooks#section-redeemed-gift-card
	GiftCardNotificationRedeemed struct {
		GiftCard GiftCard `xml:"gift_card,omitempty"`
	}

	// GiftCardNotificationUpdatedBalance is sent when the gift card's balance decreases from use on an invoice or increases if credit is returned to the account from a failed invoice.
	// https://dev.recurly.com/page/webhooks#section-updated-balance-gift-card
	GiftCardNotificationUpdatedBalance struct {
		GiftCard GiftCard `xml:"gift_card,omitempty"`
	}
)
