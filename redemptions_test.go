package recurly

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestRedemptions_GetForAccount(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/accounts/1/redemption", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(200)
		fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>
        <redemption href="https://your-subdomain.recurly.com/v2/accounts/1/redemption">
            <coupon href="https://your-subdomain.recurly.com/v2/coupons/special"/>
            <account href="https://your-subdomain.recurly.com/v2/accounts/1"/>
            <single_use type="boolean">false</single_use>
            <total_discounted_in_cents type="integer">0</total_discounted_in_cents>
            <currency>USD</currency>
            <state>active</state>
            <created_at type="datetime">2011-06-27T12:34:56Z</created_at>
        </redemption>`)
	})

	r, redemption, err := client.Redemptions.GetForAccount("1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if r.IsError() {
		t.Fatal("expected get redemption to return OK")
	}

	ts, _ := time.Parse(DateTimeFormat, "2011-06-27T12:34:56Z")
	if !reflect.DeepEqual(redemption, &Redemption{
		CouponCode:             "special",
		AccountCode:            "1",
		SingleUse:              NewBool(false),
		TotalDiscountedInCents: 0,
		Currency:               "USD",
		State:                  "active",
		CreatedAt:              NewTime(ts),
	}) {
		t.Fatalf("unexpected redemption: %v", redemption)
	}
}

func TestRedemptions_GetForAccount_ErrNotFound(t *testing.T) {
	setup()
	defer teardown()

	var invoked bool
	mux.HandleFunc("/v2/accounts/1/redemption", func(w http.ResponseWriter, r *http.Request) {
		invoked = true
		w.WriteHeader(http.StatusNotFound)
	})

	_, redemption, err := client.Redemptions.GetForAccount("1")
	if !invoked {
		t.Fatal("handler not invoked")
	} else if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if redemption != nil {
		t.Fatalf("expected redemption to be nil: %#v", redemption)
	}
}

func TestRedemptions_GetForInvoice(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/invoices/1108/redemption", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("expected %s request, given %s", "GET", r.Method)
		}
		w.WriteHeader(200)
		fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>
        <redemption href="https://your-subdomain.recurly.com/v2/accounts/1/redemption">
            <coupon href="https://your-subdomain.recurly.com/v2/coupons/special"/>
            <account href="https://your-subdomain.recurly.com/v2/accounts/1"/>
            <single_use type="boolean">true</single_use>
            <total_discounted_in_cents type="integer">0</total_discounted_in_cents>
            <currency>USD</currency>
            <state>inactive</state>
            <created_at type="datetime">2011-06-27T12:34:56Z</created_at>
        </redemption>`)
	})

	r, redemption, err := client.Redemptions.GetForInvoice("1108")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if r.IsError() {
		t.Fatal("expected get redemption to return OK")
	}

	ts, _ := time.Parse(DateTimeFormat, "2011-06-27T12:34:56Z")
	if !reflect.DeepEqual(redemption, &Redemption{
		CouponCode:             "special",
		AccountCode:            "1",
		SingleUse:              NewBool(true),
		TotalDiscountedInCents: 0,
		Currency:               "USD",
		State:                  "inactive",
		CreatedAt:              NewTime(ts),
	}) {
		t.Fatalf("unexpected redemption: %v", redemption)
	}
}

func TestRedemptions_GetForInvoice_ErrNotFound(t *testing.T) {
	setup()
	defer teardown()

	var invoked bool
	mux.HandleFunc("/v2/invoices/1108/redemption", func(w http.ResponseWriter, r *http.Request) {
		invoked = true
		w.WriteHeader(http.StatusNotFound)
	})

	_, redemption, err := client.Redemptions.GetForInvoice("1108")
	if !invoked {
		t.Fatal("handler not invoked")
	} else if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if redemption != nil {
		t.Fatalf("expected redemption to be nil: %#v", redemption)
	}
}

func TestRedemptions_RedeemCoupon(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/coupons/special/redeem", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		var given bytes.Buffer
		given.ReadFrom(r.Body)
		expected := "<redemption><account_code>1</account_code><currency>USD</currency></redemption>"
		if expected != given.String() {
			t.Fatalf("unexpected input: %s", given.String())
		}

		w.WriteHeader(201)
		fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?><redemption></redemption>`)
	})

	r, _, err := client.Redemptions.Redeem("special", "1", "USD")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if r.IsError() {
		t.Fatal("expected redeeming add on to return OK")
	}
}

func TestRedemptions_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/accounts/27/redemption", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(204)
	})

	r, err := client.Redemptions.Delete("27")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if r.IsError() {
		t.Fatal("expected delete add on to return OK")
	}
}
