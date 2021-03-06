package webhooks

const (
	// Account notifications.
	AccountNotificationNewXMLName                     = "new_account_notification"
	AccountNotificationUpdatedXMLName                 = "updated_account_notification"
	AccountNotificationCanceledXMLName                = "canceled_account_notification"
	AccountNotificationBillingInfoUpdatedXMLName      = "billing_info_updated_notification"
	AccountNotificationBillingInfoUpdateFailedXMLName = "billing_info_update_failed_notification"
	AccountNotificationNewShippingAddressXMLName      = "new_shipping_address_notification"
	AccountNotificationUpdatedShippingAddressXMLName  = "updated_shipping_address_notification"
	AccountNotificationDeletedShippingAddressXMLName  = "deleted_shipping_address_notification"
)

// Account types.
type (
	// AccountNotificationNew is sent when a new account is created.
	// https://dev.recurly.com/page/webhooks#section-new-account
	AccountNotificationNew struct {
		Account Account `xml:"account"`
	}

	// AccountNotificationUpdated is sent when an account is updated.
	// https://dev.recurly.com/page/webhooks#section-updated-account
	AccountNotificationUpdated struct {
		Account Account `xml:"account"`
	}

	// AccountNotificationCanceled is sent when an account is closed.
	// https://dev.recurly.com/page/webhooks#section-closed-account
	AccountNotificationCanceled struct {
		Account Account `xml:"account"`
	}

	// AccountNotificationBillingInfoUpdated is sent when billing information is successfully created with a credit card or updated with a credit card or token.
	// https://dev.recurly.com/page/webhooks#section-updated-billing-information
	AccountNotificationBillingInfoUpdated struct {
		Account Account `xml:"account"`
	}

	// AccountNotificationBillingInfoUpdateFailed is sent when an existing account unsuccessfully tries to update the billing information.
	// https://dev.recurly.com/page/webhooks#section-failed-billing-information-update
	AccountNotificationBillingInfoUpdateFailed struct {
		Account Account `xml:"account"`
	}

	// AccountNotificationNewShippingAddress is sent when shipping address is added.
	// https://dev.recurly.com/page/webhooks#section-a-new-shipping-address-is-created
	AccountNotificationNewShippingAddress struct {
		Account         Account         `xml:"account"`
		ShippingAddress ShippingAddress `xml:"shipping_address"`
	}

	// AccountNotificationUpdatedShippingAddress is sent when shipping address is updated.
	// https://dev.recurly.com/page/webhooks#section-an-existing-shipping-address-is-edited
	AccountNotificationUpdatedShippingAddress struct {
		Account         Account         `xml:"account"`
		ShippingAddress ShippingAddress `xml:"shipping_address"`
	}

	// AccountNotificationDeletedShippingAddress is sent when a customer updates or adds billing information.
	// https://dev.recurly.com/page/webhooks#section-an-existing-shipping-address-is-deleted
	AccountNotificationDeletedShippingAddress struct {
		Account         Account         `xml:"account"`
		ShippingAddress ShippingAddress `xml:"shipping_address"`
	}
)
