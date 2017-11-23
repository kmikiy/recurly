package recurly

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestAutomatedExports_ListExportDates(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/export_dates", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(200)
		fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>
		<export_dates type="array">
			<export_date>
			  <date>2016-08-01</date>
			  <export_files href="https://your-subdomain.recurly.com/v2/export_dates/2016-08-01/export_files"/>
			</export_date>
		 </export_dates>`)
	})

	resp, exportDates, err := client.AutomatedExports.ListExportDates()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if resp.IsError() {
		t.Fatal("expected list automated exports to return OK")
	}

	ts, _ := time.Parse(DateDateFormat, "2016-08-01")
	if !reflect.DeepEqual(exportDates, []ExportDate{
		{
			XMLName: xml.Name{Local: "export_date"},
			Date:    NewDate(ts),
		},
	}) {
		t.Fatalf("unexpected export_dates: %v", exportDates)
	}
}

func TestAutomatedExports_ListExportFilesForDates(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/export_dates/2016-08-01/export_files", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(200)
		fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>
		<export_files href="https://your-subdomain.recurly.com/v2/export_dates/2016-08-01/export_files">
			<export_file href="https://https://your-subdomain.recurly.com/v2/export_dates/2016-08-01/export_files/revenue_schedules_full.csv">
			  <name>revenue_schedules_full.csv</name>
			  <md5sum>9aa55980167ae522b27410edcd5303b0</md5sum>
			</export_file>
		</export_files>`)
	})

	ts, _ := time.Parse(DateDateFormat, "2016-08-01")
	resp, exportFiles, err := client.AutomatedExports.ListExportFilesForDate(ts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if resp.IsError() {
		t.Fatal("expected list automated exports to return OK")
	}

	if !reflect.DeepEqual(exportFiles, []ExportFile{
		{
			XMLName: xml.Name{Local: "export_file"},
			Name:    "revenue_schedules_full.csv",
			MD5Sum:  "9aa55980167ae522b27410edcd5303b0",
		},
	}) {
		t.Fatalf("unexpected export_files: %v", exportFiles)
	}
}

func Test(t *testing.T) {
	date := time.Now()
	action := fmt.Sprintf("accounts/%s", date.Format("2006-01-02"))
	t.Logf("action: %#+v\n", action)
	log.Printf("action: %#+v\n", action)
}
