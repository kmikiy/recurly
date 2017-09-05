package recurly

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"
)

func TestNullFloat(t *testing.T) {
	if !reflect.DeepEqual(NewFloat(1.5), NullFloat{Float: 1.5, Valid: true}) {
		t.Fatalf("unexpected value: %v", NewFloat(1.5))
	} else if !reflect.DeepEqual(NewFloat(0), NullFloat{Float: 0, Valid: true}) {
		t.Fatalf("unexpected value: %v", NewFloat(0))
	}

	type s struct {
		XMLName xml.Name  `xml:"s"`
		Name    string    `xml:"name"`
		Amount  NullFloat `xml:"amount,omitempty"`
	}

	tests := []struct {
		s        s
		expected string
	}{
		{s: s{XMLName: xml.Name{Local: "s"}, Name: "Bob", Amount: NewFloat(1.5)}, expected: "<s><name>Bob</name><amount>1.5</amount></s>"},
		{s: s{XMLName: xml.Name{Local: "s"}, Name: "Bob", Amount: NewFloat(0)}, expected: "<s><name>Bob</name><amount>0</amount></s>"},
		{s: s{XMLName: xml.Name{Local: "s"}, Name: "Bob"}, expected: "<s><name>Bob</name></s>"},
	}

	for i, tt := range tests {
		var given bytes.Buffer
		if err := xml.NewEncoder(&given).Encode(tt.s); err != nil {
			t.Errorf("(%d): unexpected error: %v", i, err)
		} else if tt.expected != given.String() {
			t.Errorf("(%d): unexpected value: %s", i, given.String())
		}

		var dst s
		if err := xml.NewDecoder(bytes.NewBufferString(tt.expected)).Decode(&dst); err != nil {
			t.Errorf("(%d) unexpected error: %s", i, err)
		} else if !reflect.DeepEqual(tt.s, dst) {
			t.Errorf("(%d): unexpected value: %v", i, dst)
		}
	}
}
