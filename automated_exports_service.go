package recurly

import (
	"encoding/xml"
	"fmt"
	"time"
)

var _ AutomatedExportsService = &automatedExportsImpl{}

// accountsImpl handles communication with the accounts related methods
// of the recurly API.
type automatedExportsImpl struct {
	client *Client
}

// NewAccountsImpl returns a new instance of accountsImpl.
func NewAutomatedExportsImpl(client *Client) *automatedExportsImpl {
	return &automatedExportsImpl{client: client}
}

// List returns a list of the accounts on your site.
// https://docs.recurly.com/api/accounts#list-accounts
func (s *automatedExportsImpl) ListExportDates() (*Response, []ExportDate, error) {
	req, err := s.client.newRequest("GET", "export_dates", nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var a struct {
		XMLName     xml.Name     `xml:"export_dates"`
		ExportDates []ExportDate `xml:"export_date"`
	}
	resp, err := s.client.do(req, &a)

	return resp, a.ExportDates, err
}

func (s *automatedExportsImpl) ListExportFilesForDate(date time.Time) (*Response, []ExportFile, error) {
	action := fmt.Sprintf("export_dates/%s/export_files", date.Format(DateDateFormat))
	req, err := s.client.newRequest("GET", action, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var a struct {
		XMLName     xml.Name     `xml:"export_files"`
		ExportFiles []ExportFile `xml:"export_file"`
	}
	resp, err := s.client.do(req, &a)

	return resp, a.ExportFiles, err
}
