package recurly

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestAccountEncoding ensures structs are encoded to XML properly.
// Because Recurly supports partial updates, it's important that only defined
// fields are handled properly -- including types like booleans and integers which
// have zero values that we want to send.
func TestAccounts_Encoding(t *testing.T) {
	tests := []struct {
		v        interface{}
		expected string
	}{
		{v: Account{}, expected: "<account></account>"},
		{v: Account{Code: "abc"}, expected: "<account><account_code>abc</account_code></account>"},
		{v: Account{State: "active"}, expected: "<account><state>active</state></account>"},
		{v: Account{Email: "me@example.com"}, expected: "<account><email>me@example.com</email></account>"},
		{v: Account{FirstName: "Larry"}, expected: "<account><first_name>Larry</first_name></account>"},
		{v: Account{LastName: "Larrison"}, expected: "<account><last_name>Larrison</last_name></account>"},
		{v: Account{FirstName: "Larry", LastName: "Larrison"}, expected: "<account><first_name>Larry</first_name><last_name>Larrison</last_name></account>"},
		{v: Account{CompanyName: "Acme, Inc"}, expected: "<account><company_name>Acme, Inc</company_name></account>"},
		{v: Account{VATNumber: "123456789"}, expected: "<account><vat_number>123456789</vat_number></account>"},
		{v: Account{TaxExempt: NewBool(true)}, expected: "<account><tax_exempt>true</tax_exempt></account>"},
		{v: Account{TaxExempt: NewBool(false)}, expected: "<account><tax_exempt>false</tax_exempt></account>"},
		{v: Account{AcceptLanguage: "en_US"}, expected: "<account><accept_language>en_US</accept_language></account>"},
		{v: Account{FirstName: "Larry", Address: Address{Address: "123 Main St.", City: "San Francisco", State: "CA", Zip: "94105", Country: "US"}}, expected: "<account><first_name>Larry</first_name><address><address1>123 Main St.</address1><city>San Francisco</city><state>CA</state><zip>94105</zip><country>US</country></address></account>"},
		{v: Account{Code: "test@example.com", BillingInfo: &Billing{Token: "507c7f79bcf86cd7994f6c0e"}}, expected: "<account><account_code>test@example.com</account_code><billing_info><token_id>507c7f79bcf86cd7994f6c0e</token_id></billing_info></account>"},
		{v: Address{}, expected: ""},
		{v: Address{Address: "123 Main St."}, expected: "<address><address1>123 Main St.</address1></address>"},
		{v: Address{Address2: "Unit A"}, expected: "<address><address2>Unit A</address2></address>"},
		{v: Address{City: "San Francisco"}, expected: "<address><city>San Francisco</city></address>"},
		{v: Address{State: "CA"}, expected: "<address><state>CA</state></address>"},
		{v: Address{Zip: "94105"}, expected: "<address><zip>94105</zip></address>"},
		{v: Address{Country: "US"}, expected: "<address><country>US</country></address>"},
		{v: Address{Phone: "555-555-5555"}, expected: "<address><phone>555-555-5555</phone></address>"},
	}

	for i, tt := range tests {
		var buf bytes.Buffer
		if err := xml.NewEncoder(&buf).Encode(tt.v); err != nil {
			t.Fatalf("TestAccountEncoding Error: %s", err)
		} else if buf.String() != tt.expected {
			t.Fatalf("(%d) unexpected value: %s", i, buf.String())
		}
	}
}

func TestAccounts_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/accounts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Link", `<https://your-subdomain.recurly.com/v2/accounts?cursor=1304958672>; rel="next"`)
		w.WriteHeader(200)
		fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>
		<accounts>
			<account href="https://your-subdomain.recurly.com/v2/accounts/1">
			  <adjustments href="https://your-subdomain.recurly.com/v2/accounts/1/adjustments"/>
			  <billing_info href="https://your-subdomain.recurly.com/v2/accounts/1/billing_info"/>
			  <invoices href="https://your-subdomain.recurly.com/v2/accounts/1/invoices"/>
			  <redemption href="https://your-subdomain.recurly.com/v2/accounts/1/redemption"/>
			  <subscriptions href="https://your-subdomain.recurly.com/v2/accounts/1/subscriptions"/>
			  <transactions href="https://your-subdomain.recurly.com/v2/accounts/1/transactions"/>
			  <account_code>1</account_code>
			  <state>active</state>
			  <username nil="nil"></username>
			  <email>verena@example.com</email>
			  <first_name>Verena</first_name>
			  <last_name>Example</last_name>
			  <company_name></company_name>
			  <vat_number nil="nil"></vat_number>
			  <tax_exempt type="boolean">false</tax_exempt>
			  <address>
			    <address1>123 Main St.</address1>
			    <address2 nil="nil"></address2>
			    <city>San Francisco</city>
			    <state>CA</state>
			    <zip>94105</zip>
			    <country>US</country>
			    <phone nil="nil"></phone>
			  </address>
			  <accept_language nil="nil"></accept_language>
			  <hosted_login_token>a92468579e9c4231a6c0031c4716c01d</hosted_login_token>
			  <created_at type="datetime">2011-10-25T12:00:00Z</created_at>
			</account>
		</accounts>`)
	})

	resp, accounts, err := client.Accounts.List(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if resp.IsError() {
		t.Fatal("expected list accounts to return OK")
	} else if resp.Prev() != "" {
		t.Fatalf("unexpected cursor: %s", resp.Prev())
	} else if resp.Next() != "1304958672" {
		t.Fatalf("unexpected cursor: %s", resp.Next())
	}

	ts, _ := time.Parse(DateTimeFormat, "2011-10-25T12:00:00Z")
	assert.Equal(t, accounts, []Account{Account{
		XMLName:   xml.Name{Local: "account"},
		Code:      "1",
		State:     "active",
		Email:     "verena@example.com",
		FirstName: "Verena",
		LastName:  "Example",
		TaxExempt: NewBool(false),
		Address: Address{
			Address: "123 Main St.",
			City:    "San Francisco",
			State:   "CA",
			Zip:     "94105",
			Country: "US",
		},
		HostedLoginToken: "a92468579e9c4231a6c0031c4716c01d",
		CreatedAt:        NewTime(ts),
	},
	})
}

func TestAccounts_List_Pagination(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/accounts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Link", `<https://your-subdomain.recurly.com/v2/transactions>; rel="start",
  <https://your-subdomain.recurly.com/v2/transactions?cursor=-1318344434>; rel="prev",
<https://your-subdomain.recurly.com/v2/transactions?cursor=1318388868>; rel="next"`)
		w.WriteHeader(200)
		fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?><accounts></accounts>`)
	})

	resp, _, err := client.Accounts.List(Params{"cursor": "12345"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if resp.IsError() {
		t.Fatal("expected list accounts to return OK")
	} else if resp.Prev() != "-1318344434" {
		t.Fatalf("unexpected cursor: %s", resp.Prev())
	} else if resp.Next() != "1318388868" {
		t.Fatalf("unexpected cursor: %s", resp.Next())
	}
}

func TestAccounts_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/accounts/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(200)
		fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>
			<account href="https://your-subdomain.recurly.com/v2/accounts/1">
			  <adjustments href="https://your-subdomain.recurly.com/v2/accounts/1/adjustments"/>
			  <billing_info href="https://your-subdomain.recurly.com/v2/accounts/1/billing_info"/>
			  <invoices href="https://your-subdomain.recurly.com/v2/accounts/1/invoices"/>
			  <redemption href="https://your-subdomain.recurly.com/v2/accounts/1/redemption"/>
			  <subscriptions href="https://your-subdomain.recurly.com/v2/accounts/1/subscriptions"/>
			  <transactions href="https://your-subdomain.recurly.com/v2/accounts/1/transactions"/>
			  <account_code>1</account_code>
			  <state>active</state>
			  <username nil="nil"></username>
			  <email>verena@example.com</email>
			  <first_name>Verena</first_name>
			  <last_name>Example</last_name>
			  <company_name></company_name>
			  <vat_number nil="nil"></vat_number>
			  <tax_exempt type="boolean">false</tax_exempt>
			  <address>
			    <address1>123 Main St.</address1>
			    <address2 nil="nil"></address2>
			    <city>San Francisco</city>
			    <state>CA</state>
			    <zip>94105</zip>
			    <country>US</country>
			    <phone nil="nil"></phone>
			  </address>
			  <accept_language nil="nil"></accept_language>
			  <hosted_login_token>a92468579e9c4231a6c0031c4716c01d</hosted_login_token>
			  <has_live_subscription type="boolean">true</has_live_subscription>
			  <has_active_subscription type="boolean">true</has_active_subscription>
			  <has_future_subscription type="boolean">false</has_future_subscription>
			  <has_canceled_subscription type="boolean">false</has_canceled_subscription>
			  <has_past_due_invoice type="boolean">false</has_past_due_invoice>
			  <created_at type="datetime">2011-10-25T12:00:00Z</created_at>
			  <updated_at type="datetime">2016-07-11T17:56:24Z</updated_at>
			  <closed_at nil="nil"/> 
			</account>`)
	})

	resp, a, err := client.Accounts.Get("1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if resp.IsError() {
		t.Fatal("expected get accounts to return OK")
	}

	tsCreated, _ := time.Parse(DateTimeFormat, "2011-10-25T12:00:00Z")
	tsUpdated, _ := time.Parse(DateTimeFormat, "2016-07-11T17:56:24Z")
	if !reflect.DeepEqual(a, &Account{
		XMLName:   xml.Name{Local: "account"},
		Code:      "1",
		State:     "active",
		Email:     "verena@example.com",
		FirstName: "Verena",
		LastName:  "Example",
		TaxExempt: NewBool(false),
		Address: Address{
			Address: "123 Main St.",
			City:    "San Francisco",
			State:   "CA",
			Zip:     "94105",
			Country: "US",
		},
		HostedLoginToken:        "a92468579e9c4231a6c0031c4716c01d",
		CreatedAt:               NewTime(tsCreated),
		UpdatedAt:               NewTime(tsUpdated),
		HasLiveSubscription:     NewBool(true),
		HasActiveSubscription:   NewBool(true),
		HasFutureSubscription:   NewBool(false),
		HasCanceledSubscription: NewBool(false),
		HasPastDueInvoice:       NewBool(false),
	}) {
		t.Fatalf("unexpected value: %v", a)
	}

}

func TestAccounts_Get_ErrNotFound(t *testing.T) {
	setup()
	defer teardown()

	var invoked bool
	mux.HandleFunc("/v2/accounts/1", func(w http.ResponseWriter, r *http.Request) {
		invoked = true
		w.WriteHeader(http.StatusNotFound)
	})

	_, account, err := client.Accounts.Get("1")
	if !invoked {
		t.Fatal("handler not invoked")
	} else if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if account != nil {
		t.Fatalf("expected account to be nil: %#v", account)
	}
}

func TestAccounts_LookupAccountBalance(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/accounts/1/balance", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(200)
		fmt.Fprint(w, `<account_balance href="https://your-subdomain.recurly.com/v2/accounts/1/balance">
						  <account href="https://your-subdomain.recurly.com/v2/accounts/1"/>
						  <past_due type="boolean">false</past_due>
						  <balance_in_cents>
						    <USD type="integer">3000</USD>
						    <EUR type="integer">0</EUR>
						  </balance_in_cents>
						</account_balance>`)
	})

	resp, b, err := client.Accounts.LookupAccountBalance("1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if resp.IsError() {
		t.Fatal("expected get account balance to return OK")
	}

	if !reflect.DeepEqual(b, &AccountBalance{
		XMLName:     xml.Name{Local: "account_balance"},
		AccountCode: "1",
		PastDue:     false,
		Balance:     3000,
	}) {
		t.Fatalf("unexpected value: \n%+v \n%+v", b, &AccountBalance{
			XMLName:     xml.Name{Local: "account_balance"},
			AccountCode: "1",
			PastDue:     false,
			Balance:     3000,
		})
	}
}

func TestAccounts_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/accounts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(201)
		fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?><account></account>`)
	})

	resp, _, err := client.Accounts.Create(Account{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if resp.IsError() {
		t.Fatal("expected create account to return OK")
	}
}

func TestAccounts_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/accounts/245", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(200)
		fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?><account></account>`)
	})

	resp, _, err := client.Accounts.Update("245", Account{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if resp.IsError() {
		t.Fatal("expected update account to return OK")
	}
}

func TestAccounts_Close(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/accounts/5322", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(204)
	})

	resp, err := client.Accounts.Close("5322")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if resp.IsError() {
		t.Fatal("expected close account to return OK")
	}
}

func TestAccounts_Reopen(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/accounts/5322/reopen", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(204)
	})

	resp, err := client.Accounts.Reopen("5322")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if resp.IsError() {
		t.Fatal("expected reopen account to return OK")
	}
}

func TestAccounts_ListNotes(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/accounts/abcd@example.com/notes", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		rw.WriteHeader(200)
		fmt.Fprint(rw, `<?xml version="1.0" encoding="UTF-8"?>
			<notes type="array">
			  <note>
			    <account href="https://your-subdomain.recurly.com/v2/accounts/abcd@example.com"/>
			    <message>This is my second note</message>
			    <created_at type="datetime">2013-05-14T18:53:04Z</created_at>
			  </note>
			  <note>
			    <account href="https://your-subdomain.recurly.com/v2/accounts/abcd@example.com"/>
			    <message>This is my first note</message>
			    <created_at type="datetime">2013-05-14T18:52:50Z</created_at>
			  </note>
			</notes>`)
	})

	resp, notes, err := client.Accounts.ListNotes("abcd@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if resp.IsError() {
		t.Fatal("expected list notes to return OK")
	} else if !reflect.DeepEqual(notes, []Note{
		{
			XMLName:   xml.Name{Local: "note"},
			Message:   "This is my second note",
			CreatedAt: time.Date(2013, time.May, 14, 18, 53, 4, 0, time.UTC),
		},
		{
			XMLName:   xml.Name{Local: "note"},
			Message:   "This is my first note",
			CreatedAt: time.Date(2013, time.May, 14, 18, 52, 50, 0, time.UTC),
		},
	}) {
		t.Fatalf("unexpected notes: %v", notes)
	}
}
