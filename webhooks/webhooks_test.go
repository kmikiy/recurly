package webhooks

import (
	"encoding/xml"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/kmikiy/recurly"
)

func TestParse_NewInvoiceNotification(t *testing.T) {
	xmlFile := MustOpenFile("testdata/new_invoice_notification.xml")
	createdAt := time.Date(2014, 1, 1, 20, 21, 44, 0, time.UTC)
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*NewInvoiceNotification); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &NewInvoiceNotification{
		Account: Account{
			XMLName:   xml.Name{Local: "account"},
			Code:      "1",
			Email:     "verena@example.com",
			FirstName: "Verena",
			LastName:  "Example",
		},
		Invoice: Invoice{
			XMLName:          xml.Name{Local: "invoice"},
			UUID:             "ffc64d71d4b5404e93f13aac9c63b007",
			State:            "open",
			Currency:         "USD",
			CreatedAt:        recurly.NullTime{Time: &createdAt},
			InvoiceNumber:    1000,
			TotalInCents:     1000,
			NetTerms:         recurly.NullInt{Valid: true, Int: 0},
			CollectionMethod: recurly.CollectionMethodManual,
		},
	}) {
		t.Fatalf("unexpected notification: %v", n)
	}
}

func TestParse_PastDueInvoiceNotification(t *testing.T) {
	xmlFile := MustOpenFile("testdata/past_due_invoice_notification.xml")
	createdAt := time.Date(2014, 1, 1, 20, 21, 44, 0, time.UTC)
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*PastDueInvoiceNotification); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &PastDueInvoiceNotification{
		Account: Account{
			XMLName:     xml.Name{Local: "account"},
			Code:        "1",
			Username:    "verena",
			Email:       "verena@example.com",
			FirstName:   "Verena",
			LastName:    "Example",
			CompanyName: "Company, Inc.",
		},
		Invoice: Invoice{
			XMLName:       xml.Name{Local: "invoice"},
			UUID:          "ffc64d71d4b5404e93f13aac9c63b007",
			State:         "past_due",
			CreatedAt:     recurly.NullTime{Time: &createdAt},
			InvoiceNumber: 1000,
			TotalInCents:  1100,
		},
	}) {
		t.Fatalf("unexpected notification: %v", n)
	}
}

func TestParse_SuccessfulPaymentNotification(t *testing.T) {
	xmlFile := MustOpenFile("testdata/successful_payment_notification.xml")
	if result, err := Parse(xmlFile); err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*SuccessfulPaymentNotification); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &SuccessfulPaymentNotification{
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

func TestParse_FailedPaymentNotification(t *testing.T) {
	xmlFile := MustOpenFile("testdata/failed_payment_notification.xml")
	if result, err := Parse(xmlFile); err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*FailedPaymentNotification); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &FailedPaymentNotification{
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

func TestParse_VoidPaymentNotification(t *testing.T) {
	xmlFile := MustOpenFile("testdata/void_payment_notification.xml")
	if result, err := Parse(xmlFile); err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*VoidPaymentNotification); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &VoidPaymentNotification{
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

func TestParse_SuccessfulRefundNotification(t *testing.T) {
	xmlFile := MustOpenFile("testdata/successful_refund_notification.xml")
	if result, err := Parse(xmlFile); err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*SuccessfulRefundNotification); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &SuccessfulRefundNotification{
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

func TestParse_ErrUnknownNotification(t *testing.T) {
	xmlFile := MustOpenFile("testdata/unknown_notification.xml")
	result, err := Parse(xmlFile)
	if result != nil {
		t.Fatalf("unexpected notification: %#v", result)
	} else if e, ok := err.(ErrUnknownNotification); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if err.Error() != "unknown notification: unknown_notification" {
		t.Fatalf("unexpected error string: %s", err.Error())
	} else if e.Name() != "unknown_notification" {
		t.Fatalf("unexpected notification name: %s", e.Name())
	}
}

func MustOpenFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return file
}
