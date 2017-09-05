package webhooks

import (
	"encoding/xml"
	"reflect"
	"testing"
)

func TestParse_AccountNotificationNewAccount(t *testing.T) {
	xmlFile := MustOpenFile("testdata/accounts/new_account_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*AccountNotificationNewAccount); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &AccountNotificationNewAccount{
		Account: Account{
			XMLName:   xml.Name{Local: "account"},
			Code:      "1",
			Email:     "verena@example.com",
			FirstName: "Verena",
			LastName:  "Example",
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_AccountNotificationUpdatedAccount(t *testing.T) {
	xmlFile := MustOpenFile("testdata/accounts/updated_account_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*AccountNotificationUpdatedAccount); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &AccountNotificationUpdatedAccount{
		Account: Account{
			XMLName:   xml.Name{Local: "account"},
			Code:      "1",
			Email:     "verena@example.com",
			FirstName: "Verena",
			LastName:  "Example",
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_AccountNotificationCanceledAccount(t *testing.T) {
	xmlFile := MustOpenFile("testdata/accounts/canceled_account_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*AccountNotificationCanceledAccount); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &AccountNotificationCanceledAccount{
		Account: Account{
			XMLName:   xml.Name{Local: "account"},
			Code:      "1",
			Email:     "verena@example.com",
			FirstName: "Verena",
			LastName:  "Example",
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_AccountNotificationBillingInfoUpdated(t *testing.T) {
	xmlFile := MustOpenFile("testdata/accounts/billing_info_updated_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*AccountNotificationBillingInfoUpdated); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &AccountNotificationBillingInfoUpdated{
		Account: Account{
			XMLName:   xml.Name{Local: "account"},
			Code:      "1",
			Email:     "verena@example.com",
			FirstName: "Verena",
			LastName:  "Example",
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_AccountNotificationBillingInfoUpdateFailed(t *testing.T) {
	xmlFile := MustOpenFile("testdata/accounts/billing_info_update_failed_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*AccountNotificationBillingInfoUpdateFailed); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &AccountNotificationBillingInfoUpdateFailed{
		Account: Account{
			XMLName:   xml.Name{Local: "account"},
			Code:      "1",
			Email:     "verena@example.com",
			FirstName: "Verena",
			LastName:  "Example",
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_AccountNotificationNewShippingAddress(t *testing.T) {
	xmlFile := MustOpenFile("testdata/accounts/new_shipping_address_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*AccountNotificationNewShippingAddress); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &AccountNotificationNewShippingAddress{
		Account: Account{
			XMLName:     xml.Name{Local: "account"},
			Code:        "SamSmith",
			FirstName:   "Sam",
			LastName:    "Smith",
			CompanyName: "Smith Co",
		},
		ShippingAddress: ShippingAddress{
			XMLName:   xml.Name{Local: "shipping_address"},
			ID:        2019760742762202549,
			Nickname:  "Steven",
			FirstName: "Steven",
			LastName:  "Smith",
			Street:    "231 Oregon Street",
			City:      "Portland",
			State:     "OR",
			ZIP:       "97201",
			Country:   "US",
			Email:     "stevensmith@example.com",
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_AccountNotificationUpdatedShippingAddress(t *testing.T) {
	xmlFile := MustOpenFile("testdata/accounts/updated_shipping_address_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*AccountNotificationUpdatedShippingAddress); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &AccountNotificationUpdatedShippingAddress{
		Account: Account{
			XMLName:     xml.Name{Local: "account"},
			Code:        "SamSmith",
			FirstName:   "Sam",
			LastName:    "Smith",
			CompanyName: "Smith Co",
		},
		ShippingAddress: ShippingAddress{
			XMLName:   xml.Name{Local: "shipping_address"},
			ID:        2019760742762202549,
			Nickname:  "Steven",
			FirstName: "Steven",
			LastName:  "Smith",
			Street:    "231 Oregon Street",
			City:      "Portland",
			State:     "OR",
			ZIP:       "97201",
			Country:   "US",
			Email:     "stevensmith@example.com",
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}

func TestParse_AccountNotificationDeletedShippingAddress(t *testing.T) {
	xmlFile := MustOpenFile("testdata/accounts/deleted_shipping_address_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*AccountNotificationDeletedShippingAddress); !ok {
		t.Fatalf("unexpected type: %T, result")
	} else if !reflect.DeepEqual(n, &AccountNotificationDeletedShippingAddress{
		Account: Account{
			XMLName:     xml.Name{Local: "account"},
			Code:        "SamSmith",
			FirstName:   "Sam",
			LastName:    "Smith",
			CompanyName: "Smith Co",
		},
		ShippingAddress: ShippingAddress{
			XMLName:   xml.Name{Local: "shipping_address"},
			ID:        2019760742762202549,
			Nickname:  "Steven",
			FirstName: "Steven",
			LastName:  "Smith",
			Street:    "231 Oregon Street",
			City:      "Portland",
			State:     "OR",
			ZIP:       "97201",
			Country:   "US",
			Email:     "stevensmith@example.com",
		},
	}) {
		t.Fatalf("unexpected notification: %#v", n)
	}
}
