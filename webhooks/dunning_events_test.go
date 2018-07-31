package webhooks

import (
	"encoding/xml"
	"reflect"
	"testing"
	"time"

	"github.com/kmikiy/recurly"
)

func TestParse_DunningEventNew(t *testing.T) {
	// createdTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-03T20:37:21Z")
	// updatedTs, _ := time.Parse(recurly.DateTimeFormat, "2016-08-03T20:37:21Z")
	activatedAt, _ := time.Parse(recurly.DateTimeFormat, "2017-11-09T16:47:30Z")
	currentPeriodStartedAt, _ := time.Parse(recurly.DateTimeFormat, "2018-02-09T16:47:30Z")
	currentPeriodEndsAt, _ := time.Parse(recurly.DateTimeFormat, "2018-03-09T16:47:30Z")

	xmlFile := MustOpenFile("testdata/dunning_events/new_dunning_event_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*DunningEventNotificationNew); !ok {
		t.Fatalf("unexpected type: %T", result)
	} else if !reflect.DeepEqual(n, &DunningEventNotificationNew{
		Account: Account{
			XMLName: xml.Name{
				Space: "",
				Local: "account",
			},
			Code:        "1234",
			Username:    "",
			Email:       "",
			FirstName:   "",
			LastName:    "",
			CompanyName: "",
			Phone:       "",
		},
		Invoice: Invoice{
			XMLName: xml.Name{
				Space: "",
				Local: "invoice",
			},
			SubscriptionUUID:    "",
			UUID:                "424a9d4a2174b4f39bc776426aa19c32",
			State:               "past_due",
			InvoiceNumberPrefix: "",
			InvoiceNumber:       1813,
			PONumber:            "",
			VATNumber:           "",
			TotalInCents:        4500,
			Currency:            "USD",
			NetTerms: recurly.NullInt{
				Int:   30,
				Valid: true,
			},
			CollectionMethod: "manual",
		},
		Subscription: recurly.Subscription{
			XMLName: xml.Name{
				Space: "",
				Local: "subscription",
			},

			Plan: recurly.NestedPlan{
				Code: "gold",
				Name: "Gold",
			},
			AccountCode:            "",
			UUID:                   "4110792b3b01967d854f674b7282f542",
			State:                  "active",
			UnitAmountInCents:      0,
			Currency:               "",
			Quantity:               1,
			TotalAmountInCents:     4500,
			ActivatedAt:            recurly.NewTime(activatedAt),
			CurrentPeriodStartedAt: recurly.NewTime(currentPeriodStartedAt),
			CurrentPeriodEndsAt:    recurly.NewTime(currentPeriodEndsAt),
			TaxInCents:             0,
			TaxType:                "",
			TaxRegion:              "",
			TaxRate:                0,
			PONumber:               "",
			SubscriptionAddOns: []recurly.SubscriptionAddOn{
				recurly.SubscriptionAddOn{
					XMLName: xml.Name{
						Space: "",
						Local: "subscription_add_on",
					},
					Type:              "fixed",
					Code:              "training_classes",
					UnitAmountInCents: 700,
					Quantity:          1,
				},
				recurly.SubscriptionAddOn{
					XMLName: xml.Name{
						Space: "",
						Local: "subscription_add_on",
					},
					Type:              "fixed",
					Code:              "executive-brief",
					UnitAmountInCents: 2300,
					Quantity:          1,
				},
			},
			CollectionMethod: "",
		},
		Transaction: Transaction{
			XMLName: xml.Name{
				Space: "",
				Local: "transaction",
			},
			UUID:              "397083a9a871b53a3d5a4c469fa1216a",
			InvoiceNumber:     1002,
			SubscriptionUUID:  "396e4e17640ca516c2f3a84e47ae91dd",
			Action:            "purchase",
			AmountInCents:     2499,
			Status:            "declined",
			Message:           "Transaction Normal",
			GatewayErrorCodes: "00",
			FailureType:       "invalid_data",
			Reference:         "115948823",
			Source:            "subscription",
			Test: recurly.NullBool{
				Bool:  true,
				Valid: true,
			},
			Voidable: recurly.NullBool{
				Bool:  false,
				Valid: true,
			},
			Refundable: recurly.NullBool{
				Bool:  false,
				Valid: true,
			},
		},
	},
	) {
		t.Fatalf("unexpected notification: %#v", n)
	}

}
