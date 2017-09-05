package webhooks

import (
	"encoding/xml"
	"reflect"
	"testing"
	"time"

	"github.com/kmikiy/recurly"
)

func TestParse_InvoiceNotificationNew(t *testing.T) {
	xmlFile := MustOpenFile("testdata/invoices/new_invoice_notification.xml")
	createdAt := time.Date(2014, 1, 1, 20, 21, 44, 0, time.UTC)
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*InvoiceNotificationNew); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &InvoiceNotificationNew{
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
	xmlFile := MustOpenFile("testdata/invoices/past_due_invoice_notification.xml")
	createdAt := time.Date(2014, 1, 1, 20, 21, 44, 0, time.UTC)
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*InvoiceNotificationPastDue); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &InvoiceNotificationPastDue{
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
