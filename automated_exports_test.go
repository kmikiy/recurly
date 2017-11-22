package recurly

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestAutomatedExports_List(t *testing.T) {
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

	resp, exportDates, err := client.AutomatedExports.List()
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
			ExportFiles: ExportFile{
				Href: "https://your-subdomain.recurly.com/v2/export_dates/2016-08-01/export_files",
			},
		},
	}) {
		t.Fatalf("unexpected adjustments: %v", exportDates)
	}
}
