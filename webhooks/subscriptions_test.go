package webhooks

import (
	"encoding/xml"
	"reflect"
	"testing"
	"time"

	"github.com/kmikiy/recurly"
)

func TestParse_SubscriptionNotificationNewSubscription(t *testing.T) {
	activatedTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-23T22:05:03Z")
	canceledTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-23T22:05:43Z")
	expiresTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-24T22:05:03Z")
	startedTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-23T22:05:03Z")
	endsTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-24T22:05:03Z")

	xmlFile := MustOpenFile("testdata/subscriptions/new_subscription_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*SubscriptionNotificationNewSubscription); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &SubscriptionNotificationNewSubscription{
		Account: Account{
			XMLName:   xml.Name{Local: "account"},
			Code:      "1",
			Email:     "verena@example.com",
			FirstName: "Verena",
			LastName:  "Example",
		},
		Subscription: recurly.Subscription{
			XMLName: xml.Name{Local: "subscription"},
			Plan: recurly.NestedPlan{
				Code: "bronze",
				Name: "Bronze Plan",
			},
			UUID:                   "d1b6d359a01ded71caed78eaa0fedf8e",
			State:                  "active",
			Quantity:               2,
			TotalAmountInCents:     17000,
			ActivatedAt:            recurly.NewTime(activatedTs),
			CanceledAt:             recurly.NewTime(canceledTs),
			ExpiresAt:              recurly.NewTime(expiresTs),
			CurrentPeriodStartedAt: recurly.NewTime(startedTs),
			CurrentPeriodEndsAt:    recurly.NewTime(endsTs),
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_SubscriptionNotificationUpdatedSubscription(t *testing.T) {
	activatedTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-23T22:05:03Z")
	canceledTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-23T22:05:43Z")
	expiresTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-24T22:05:03Z")
	startedTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-23T22:05:03Z")
	endsTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-24T22:05:03Z")

	xmlFile := MustOpenFile("testdata/subscriptions/updated_subscription_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*SubscriptionNotificationUpdatedSubscription); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &SubscriptionNotificationUpdatedSubscription{
		Account: Account{
			XMLName:   xml.Name{Local: "account"},
			Code:      "1",
			Email:     "verena@example.com",
			FirstName: "Verena",
			LastName:  "Example",
		},
		Subscription: recurly.Subscription{
			XMLName: xml.Name{Local: "subscription"},
			Plan: recurly.NestedPlan{
				Code: "1dpt",
				Name: "Subscription One",
			},
			UUID:                   "292332928954ca62fa48048be5ac98ec",
			State:                  "active",
			Quantity:               1,
			TotalAmountInCents:     200,
			ActivatedAt:            recurly.NewTime(activatedTs),
			CanceledAt:             recurly.NewTime(canceledTs),
			ExpiresAt:              recurly.NewTime(expiresTs),
			CurrentPeriodStartedAt: recurly.NewTime(startedTs),
			CurrentPeriodEndsAt:    recurly.NewTime(endsTs),
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_SubscriptionNotificationCanceledSubscription(t *testing.T) {
	activatedTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-23T22:05:03Z")
	canceledTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-23T22:05:43Z")
	expiresTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-24T22:05:03Z")
	startedTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-23T22:05:03Z")
	endsTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-24T22:05:03Z")

	xmlFile := MustOpenFile("testdata/subscriptions/canceled_subscription_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*SubscriptionNotificationCanceledSubscription); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &SubscriptionNotificationCanceledSubscription{
		Account: Account{
			XMLName:   xml.Name{Local: "account"},
			Code:      "1",
			Email:     "verena@example.com",
			FirstName: "Verena",
			LastName:  "Example",
		},
		Subscription: recurly.Subscription{
			XMLName: xml.Name{Local: "subscription"},
			Plan: recurly.NestedPlan{
				Code: "1dpt",
				Name: "Subscription One",
			},
			UUID:                   "dccd742f4710e78515714d275839f891",
			State:                  "canceled",
			Quantity:               1,
			TotalAmountInCents:     200,
			ActivatedAt:            recurly.NewTime(activatedTs),
			CanceledAt:             recurly.NewTime(canceledTs),
			ExpiresAt:              recurly.NewTime(expiresTs),
			CurrentPeriodStartedAt: recurly.NewTime(startedTs),
			CurrentPeriodEndsAt:    recurly.NewTime(endsTs),
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_SubscriptionNotificationExpiredSubscription(t *testing.T) {
	activatedTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-23T22:05:03Z")
	canceledTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-23T22:05:43Z")
	expiresTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-24T22:05:03Z")
	startedTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-23T22:05:03Z")
	endsTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-24T22:05:03Z")

	xmlFile := MustOpenFile("testdata/subscriptions/expired_subscription_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*SubscriptionNotificationExpiredSubscription); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &SubscriptionNotificationExpiredSubscription{
		Account: Account{
			XMLName:   xml.Name{Local: "account"},
			Code:      "1",
			Email:     "verena@example.com",
			FirstName: "Verena",
			LastName:  "Example",
		},
		Subscription: recurly.Subscription{
			XMLName: xml.Name{Local: "subscription"},
			Plan: recurly.NestedPlan{
				Code: "1dpt",
				Name: "Subscription One",
			},
			UUID:                   "d1b6d359a01ded71caed78eaa0fedf8e",
			State:                  "expired",
			Quantity:               1,
			TotalAmountInCents:     200,
			ActivatedAt:            recurly.NewTime(activatedTs),
			CanceledAt:             recurly.NewTime(canceledTs),
			ExpiresAt:              recurly.NewTime(expiresTs),
			CurrentPeriodStartedAt: recurly.NewTime(startedTs),
			CurrentPeriodEndsAt:    recurly.NewTime(endsTs),
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_SubscriptionNotificationRenewedSubscription(t *testing.T) {
	activatedTs, _ := time.Parse(recurly.DateTimeFormat, "2010-07-22T20:42:05Z")
	startedTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-22T20:42:05Z")
	endsTs, _ := time.Parse(recurly.DateTimeFormat, "2010-10-22T20:42:05Z")

	xmlFile := MustOpenFile("testdata/subscriptions/renewed_subscription_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*SubscriptionNotificationRenewedSubscription); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &SubscriptionNotificationRenewedSubscription{
		Account: Account{
			XMLName:     xml.Name{Local: "account"},
			Code:        "1",
			Email:       "verena@example.com",
			FirstName:   "Verena",
			LastName:    "Example",
			CompanyName: "Company, Inc.",
		},
		Subscription: recurly.Subscription{
			XMLName: xml.Name{Local: "subscription"},
			Plan: recurly.NestedPlan{
				Code: "bootstrap",
				Name: "Bootstrap",
			},
			UUID:                   "6ab458a887d38070807ebb3bed7ac1e5",
			State:                  "active",
			Quantity:               1,
			TotalAmountInCents:     9900,
			ActivatedAt:            recurly.NewTime(activatedTs),
			CurrentPeriodStartedAt: recurly.NewTime(startedTs),
			CurrentPeriodEndsAt:    recurly.NewTime(endsTs),
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_SubscriptionNotificationReactivatedSubscription(t *testing.T) {
	activatedTs, _ := time.Parse(recurly.DateTimeFormat, "2010-07-22T20:42:05Z")
	startedTs, _ := time.Parse(recurly.DateTimeFormat, "2010-09-22T20:42:05Z")
	endsTs, _ := time.Parse(recurly.DateTimeFormat, "2010-10-22T20:42:05Z")

	xmlFile := MustOpenFile("testdata/subscriptions/reactivated_account_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*SubscriptionNotificationReactivatedSubscription); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &SubscriptionNotificationReactivatedSubscription{
		Account: Account{
			XMLName:   xml.Name{Local: "account"},
			Code:      "1",
			Email:     "verena@example.com",
			FirstName: "Verena",
			LastName:  "Example",
		},
		Subscription: recurly.Subscription{
			XMLName: xml.Name{Local: "subscription"},
			Plan: recurly.NestedPlan{
				Code: "bootstrap",
				Name: "Bootstrap",
			},
			UUID:                   "6ab458a887d38070807ebb3bed7ac1e5",
			State:                  "active",
			Quantity:               1,
			TotalAmountInCents:     9900,
			ActivatedAt:            recurly.NewTime(activatedTs),
			CurrentPeriodStartedAt: recurly.NewTime(startedTs),
			CurrentPeriodEndsAt:    recurly.NewTime(endsTs),
			CollectionMethod:       "automatic",
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}
