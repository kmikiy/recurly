package recurly

import (
	"encoding/xml"
)

// Account represents an individual account on your site
type ExportDate struct {
	XMLName     xml.Name   `xml:"export_date"`
	ExportFiles ExportFile `xml:"export_files,omitempty"`
	Date        NullDate   `xml:"date,omitempty"`
}

type ExportFile struct {
	Href string `xml:"href,attr"`
}
