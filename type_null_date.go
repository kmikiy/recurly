package recurly

import (
	"encoding/xml"
	"time"
)

// DateTimeFormat is the format Recurly uses to represent datetimes.
const DateDateFormat = "2006-01-02"

// NullTime is used for properly handling time.Time types that could be null.
type NullDate struct {
	*time.Time
	Raw string `xml:",innerxml"`
}

// NewTime generates a new NullTime.
func NewDate(t time.Time) NullDate {
	t = t.UTC()
	return NullDate{Time: &t}
}

// NewTimeFromString generates a new NullTime based on a
// time string in the DateTimeFormat format.
// This is primarily used in unit testing.
func NewDateFromString(str string) NullDate {
	t, _ := time.Parse(DateDateFormat, str)
	return NullDate{Time: &t}
}

// UnmarshalXML unmarshals an int properly, as well as marshaling an empty string to nil.
func (t *NullDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err == nil && v != "" {
		parsed, err := time.Parse(DateDateFormat, v)
		if err != nil {
			return err
		}

		*t = NewDate(parsed)
	}

	return nil
}

// MarshalXML marshals times into their proper format. Otherwise nothing is
// marshaled. All times are sent in UTC.
func (t NullDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.Time != nil {
		e.EncodeElement(t.String(), start)
	}

	return nil
}

// String returns a string representation of the time in UTC using the
// DateTimeFormat constant as the format.
func (t NullDate) String() string {
	if t.Time != nil {
		return t.Time.UTC().Format(DateDateFormat)
	}

	return ""
}
