package recurly

import (
	"encoding/xml"
)

// Account represents an individual account on your site
type ExportDate struct {
	XMLName xml.Name `xml:"export_date"`
	Date    NullDate `xml:"date,omitempty"`
}

type ExportFile struct {
	XMLName xml.Name `xml:"export_file"`
	Name    string   `xml:"name"`
	MD5Sum  string   `xml:"md5sum"`
}
