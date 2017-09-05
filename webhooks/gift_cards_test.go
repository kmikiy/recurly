package webhooks

import (
	"encoding/xml"
	"reflect"
	"testing"
	"time"

	"github.com/kmikiy/recurly"
)

func TestParse_GiftCardNotificationPurchased(t *testing.T) {
	createdTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-03T20:37:21Z")
	updatedTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-03T20:37:21Z")

	xmlFile := MustOpenFile("testdata/gift_cards/purchased_gift_card_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*GiftCardNotificationPurchased); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &GiftCardNotificationPurchased{
		GiftCard: GiftCard{
			XMLName:           xml.Name{Local: "gift_card"},
			RedemptionCode:    "1A5069E266AED435",
			ID:                2008976331180115114,
			ProductCode:       "gift_card",
			UnitAmountInCents: recurly.NewInt(1000),
			Currency:          "USD",
			GifterAccountCode: "84395",
			InvoiceNumber:     1105,
			Delivery: GiftCardDelivery{
				Method:          "email",
				EmailAddress:    "john@example.com",
				FirstName:       "John",
				LastName:        "Smith",
				GifterName:      "Sally",
				PersonalMessage: "Hi John, Happy Birthday! I hope you have a great day! Love, Sally",
			},
			CreatedAt: recurly.NewTime(createdTs),
			UpdatedAt: recurly.NewTime(updatedTs),
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_GiftCardNotificationCanceled(t *testing.T) {
	createdTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-03T20:37:21Z")
	updatedTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-03T22:00:00Z")
	deliveredTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-03T20:37:22Z")
	canceledTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-04T20:30:22Z")

	xmlFile := MustOpenFile("testdata/gift_cards/canceled_gift_card_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*GiftCardNotificationCanceled); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &GiftCardNotificationCanceled{
		GiftCard: GiftCard{
			XMLName:           xml.Name{Local: "gift_card"},
			RedemptionCode:    "1A5069E266AED435",
			ID:                2008976331180115114,
			ProductCode:       "gift_card",
			UnitAmountInCents: recurly.NewInt(1000),
			Currency:          "USD",
			GifterAccountCode: "84395",
			InvoiceNumber:     1105,
			Delivery: GiftCardDelivery{
				Method:          "email",
				EmailAddress:    "john@example.com",
				FirstName:       "John",
				LastName:        "Smith",
				GifterName:      "Sally",
				PersonalMessage: "Hi John, Happy Birthday! I hope you have a great day! Love, Sally",
			},
			CreatedAt:   recurly.NewTime(createdTs),
			UpdatedAt:   recurly.NewTime(updatedTs),
			DeliveredAt: recurly.NewTime(deliveredTs),
			CanceledAt:  recurly.NewTime(canceledTs),
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_GiftCardNotificationUpdated(t *testing.T) {
	createdTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-03T20:37:21Z")
	updatedTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-03T22:00:00Z")
	deliveredTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-03T20:37:22Z")

	xmlFile := MustOpenFile("testdata/gift_cards/updated_gift_card_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*GiftCardNotificationUpdated); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &GiftCardNotificationUpdated{
		GiftCard: GiftCard{
			XMLName:           xml.Name{Local: "gift_card"},
			RedemptionCode:    "1A5069E266AED435",
			ID:                2008976331180115114,
			ProductCode:       "gift_card",
			UnitAmountInCents: recurly.NewInt(1000),
			Currency:          "USD",
			GifterAccountCode: "84395",
			InvoiceNumber:     1105,
			Delivery: GiftCardDelivery{
				Method:          "email",
				EmailAddress:    "john@example.com",
				FirstName:       "John",
				LastName:        "Smith",
				GifterName:      "Sally",
				PersonalMessage: "Hi John, Happy Birthday! I hope you have a great day! Love, Sally",
			},
			CreatedAt:   recurly.NewTime(createdTs),
			UpdatedAt:   recurly.NewTime(updatedTs),
			DeliveredAt: recurly.NewTime(deliveredTs),
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_GiftCardNotificationRegenerated(t *testing.T) {
	createdTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-03T20:37:21Z")
	updatedTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-03T22:00:00Z")
	deliveredTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-03T20:37:22Z")

	xmlFile := MustOpenFile("testdata/gift_cards/regenerated_gift_card_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*GiftCardNotificationRegenerated); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &GiftCardNotificationRegenerated{
		GiftCard: GiftCard{
			XMLName:           xml.Name{Local: "gift_card"},
			RedemptionCode:    "1A5069E266AED435",
			ID:                2008976331180115114,
			ProductCode:       "gift_card",
			UnitAmountInCents: recurly.NewInt(1000),
			Currency:          "USD",
			GifterAccountCode: "84395",
			InvoiceNumber:     1105,
			Delivery: GiftCardDelivery{
				Method:          "email",
				EmailAddress:    "john@example.com",
				FirstName:       "John",
				LastName:        "Smith",
				GifterName:      "Sally",
				PersonalMessage: "Hi John, Happy Birthday! I hope you have a great day! Love, Sally",
			},
			CreatedAt:   recurly.NewTime(createdTs),
			UpdatedAt:   recurly.NewTime(updatedTs),
			DeliveredAt: recurly.NewTime(deliveredTs),
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_GiftCardNotificationRedeemed(t *testing.T) {
	createdTs, _ := time.Parse(recurly.DateTimeFormat, "2016-07-29T21:41:11Z")
	updatedTs, _ := time.Parse(recurly.DateTimeFormat, "2016-07-29T21:50:38Z")
	deliveredTs, _ := time.Parse(recurly.DateTimeFormat, "2016-07-29T21:50:38Z")
	redeemedTs, _ := time.Parse(recurly.DateTimeFormat, "2016-07-29T21:50:38Z")

	xmlFile := MustOpenFile("testdata/gift_cards/redeemed_gift_card_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*GiftCardNotificationRedeemed); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &GiftCardNotificationRedeemed{
		GiftCard: GiftCard{
			XMLName:              xml.Name{Local: "gift_card"},
			RedemptionCode:       "AB54200960E33C93",
			ID:                   2005384587788419212,
			ProductCode:          "gift_card",
			UnitAmountInCents:    recurly.NewInt(1000),
			Currency:             "USD",
			GifterAccountCode:    "3543456",
			RecipientAccountCode: "3547000",
			InvoiceNumber:        1099,
			Delivery: GiftCardDelivery{
				Method:          "email",
				FirstName:       "John",
				LastName:        "Smith",
				GifterName:      "Sally",
				PersonalMessage: "Hi John, Happy Birthday! I hope you have a great day! Love, Sally",
			},
			CreatedAt:   recurly.NewTime(createdTs),
			UpdatedAt:   recurly.NewTime(updatedTs),
			DeliveredAt: recurly.NewTime(deliveredTs),
			RedeemedAt:  recurly.NewTime(redeemedTs),
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_GiftCardNotificationUpdatedBalance(t *testing.T) {
	createdTs, _ := time.Parse(recurly.DateTimeFormat, "2016-07-29T21:41:11Z")
	updatedTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-02T23:50:38Z")
	deliveredTs, _ := time.Parse(recurly.DateTimeFormat, "2016-07-29T21:50:38Z")
	redeemedTs, _ := time.Parse(recurly.DateTimeFormat, "2016-07-29T21:50:38Z")

	xmlFile := MustOpenFile("testdata/gift_cards/updated_balance_gift_card_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*GiftCardNotificationUpdatedBalance); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &GiftCardNotificationUpdatedBalance{
		GiftCard: GiftCard{
			XMLName:              xml.Name{Local: "gift_card"},
			RedemptionCode:       "AB54200960E33C93",
			ID:                   2005384587788419212,
			ProductCode:          "gift_card",
			UnitAmountInCents:    recurly.NewInt(1000),
			Currency:             "USD",
			GifterAccountCode:    "3543456",
			RecipientAccountCode: "3547000",
			InvoiceNumber:        1099,
			Delivery: GiftCardDelivery{
				Method:          "email",
				FirstName:       "John",
				LastName:        "Smith",
				GifterName:      "Sally",
				PersonalMessage: "Hi John, Happy Birthday! I hope you have a great day! Love, Sally",
			},
			CreatedAt:   recurly.NewTime(createdTs),
			UpdatedAt:   recurly.NewTime(updatedTs),
			DeliveredAt: recurly.NewTime(deliveredTs),
			RedeemedAt:  recurly.NewTime(redeemedTs),
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}
