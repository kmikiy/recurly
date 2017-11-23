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

type DownloadExportFile struct {
	XMLName     xml.Name `xml:"export_file"`
	ExpiresAt   NullTime `xml:"expires_at"`
	DownloadURL string   `xml:"download_url"`
}
