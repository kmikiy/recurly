package recurly

import "encoding/xml"

// NullFloat is used for properly handling float types that could be null.
type NullFloat struct {
	Float float64
	Valid bool
}

// NewFloat builds a new NewFloat struct.
func NewFloat(i float64) NullFloat {
	return NullFloat{Float: i, Valid: true}
}

// UnmarshalXML unmarshals an int properly, as well as marshaling an empty string to nil.
func (n *NullFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v float64
	err := d.DecodeElement(&v, &start)
	if err == nil {
		*n = NullFloat{Float: v, Valid: true}
	}

	return nil
}

// MarshalXML marshals NullFloats greater than zero to XML. Otherwise nothing is
// marshaled.
func (n NullFloat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if n.Valid {
		e.EncodeElement(n.Float, start)
	}

	return nil
}
