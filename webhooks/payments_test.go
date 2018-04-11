package webhooks

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/kmikiy/recurly"
)

func TestParse_PaymentNotificationSuccessful(t *testing.T) {
	xmlFile := MustOpenFile("testdata/payments/successful_payment_notification.xml")
	if result, err := Parse(xmlFile); err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*PaymentNotificationSuccessful); !ok {
		t.Fatalf("unexpected type: %T", result)
	} else if !reflect.DeepEqual(n, &PaymentNotificationSuccessful{
		Account: Account{
			XMLName:     xml.Name{Local: "account"},
			Code:        "1",
			Username:    "verena",
			Email:       "verena@example.com",
			FirstName:   "Verena",
			LastName:    "Example",
			CompanyName: "Company, Inc.",
		},
		Transaction: Transaction{
			XMLName:       xml.Name{Local: "transaction"},
			UUID:          "a5143c1d3a6f4a8287d0e2cc1d4c0427",
			InvoiceNumber: 2059,
			Action:        "purchase",
			AmountInCents: 1000,
			Status:        "success",
			Message:       "Bogus Gateway: Forced success",
			Reference:     "reference",
			Source:        "subscription",
			Test:          recurly.NullBool{Valid: true, Bool: true},
			Voidable:      recurly.NullBool{Valid: true, Bool: true},
			Refundable:    recurly.NullBool{Valid: true, Bool: true},
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_PaymentNotificationFailed(t *testing.T) {
	xmlFile := MustOpenFile("testdata/payments/failed_payment_notification.xml")
	if result, err := Parse(xmlFile); err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*PaymentNotificationFailed); !ok {
		t.Fatalf("unexpected type: %T", result)
	} else if !reflect.DeepEqual(n, &PaymentNotificationFailed{
		Account: Account{
			XMLName:     xml.Name{Local: "account"},
			Code:        "1",
			Username:    "verena",
			Email:       "verena@example.com",
			FirstName:   "Verena",
			LastName:    "Example",
			CompanyName: "Company, Inc.",
		},
		Transaction: Transaction{
			XMLName:          xml.Name{Local: "transaction"},
			UUID:             "a5143c1d3a6f4a8287d0e2cc1d4c0427",
			InvoiceNumber:    2059,
			SubscriptionUUID: "1974a098jhlkjasdfljkha898326881c",
			Action:           "purchase",
			AmountInCents:    1000,
			Status:           "Declined",
			Message:          "This transaction has been declined",
			FailureType:      "Declined by the gateway",
			Reference:        "reference",
			Source:           "subscription",
			Test:             recurly.NullBool{Valid: true, Bool: true},
			Voidable:         recurly.NullBool{Valid: true, Bool: false},
			Refundable:       recurly.NullBool{Valid: true, Bool: false},
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_PaymentNotificationVoid(t *testing.T) {
	xmlFile := MustOpenFile("testdata/payments/void_payment_notification.xml")
	if result, err := Parse(xmlFile); err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*PaymentNotificationVoid); !ok {
		t.Fatalf("unexpected type: %T", result)
	} else if !reflect.DeepEqual(n, &PaymentNotificationVoid{
		Account: Account{
			XMLName:     xml.Name{Local: "account"},
			Code:        "1",
			Username:    "verena",
			Email:       "verena@example.com",
			FirstName:   "Verena",
			LastName:    "Example",
			CompanyName: "Company, Inc.",
		},
		Transaction: Transaction{
			XMLName:          xml.Name{Local: "transaction"},
			UUID:             "a5143c1d3a6f4a8287d0e2cc1d4c0427",
			InvoiceNumber:    2059,
			SubscriptionUUID: "1974a098jhlkjasdfljkha898326881c",
			Action:           "purchase",
			AmountInCents:    1000,
			Status:           "void",
			Message:          "Test Gateway: Successful test transaction",
			Reference:        "reference",
			Source:           "subscription",
			Test:             recurly.NullBool{Valid: true, Bool: true},
			Voidable:         recurly.NullBool{Valid: true, Bool: true},
			Refundable:       recurly.NullBool{Valid: true, Bool: true},
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_PaymentNotificationSuccessfulRefund(t *testing.T) {
	xmlFile := MustOpenFile("testdata/payments/successful_refund_notification.xml")
	if result, err := Parse(xmlFile); err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*PaymentNotificationSuccessfulRefund); !ok {
		t.Fatalf("unexpected type: %T", result)
	} else if !reflect.DeepEqual(n, &PaymentNotificationSuccessfulRefund{
		Account: Account{
			XMLName:     xml.Name{Local: "account"},
			Code:        "1",
			Username:    "verena",
			Email:       "verena@example.com",
			FirstName:   "Verena",
			LastName:    "Example",
			CompanyName: "Company, Inc.",
		},
		Transaction: Transaction{
			XMLName:          xml.Name{Local: "transaction"},
			UUID:             "a5143c1d3a6f4a8287d0e2cc1d4c0427",
			InvoiceNumber:    2059,
			SubscriptionUUID: "1974a098jhlkjasdfljkha898326881c",
			Action:           "credit",
			AmountInCents:    1000,
			Status:           "success",
			Message:          "Bogus Gateway: Forced success",
			Reference:        "reference",
			Source:           "subscription",
			Test:             recurly.NullBool{Valid: true, Bool: true},
			Voidable:         recurly.NullBool{Valid: true, Bool: true},
			Refundable:       recurly.NullBool{Valid: true, Bool: true},
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}
