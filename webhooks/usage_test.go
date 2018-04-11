package webhooks

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/kmikiy/recurly"
	"github.com/stretchr/testify/assert"
)

func TestParse_UsageNotificationNewUsage(t *testing.T) {
	recordingTs, _ := time.Parse(recurly.DateTimeFormat, "2016-04-28T21:57:53+00:00")
	usageTs, _ := time.Parse(recurly.DateTimeFormat, "2016-04-28T21:57:53+00:00")
	createdTs, _ := time.Parse(recurly.DateTimeFormat, "2016-04-28T21:57:54+00:00")
	billedTs, _ := time.Parse(recurly.DateTimeFormat, "2016-04-28T21:57:54+00:00")

	xmlFile := MustOpenFile("testdata/usage/new_usage_notification.xml")
	result, err := Parse(xmlFile)
	if err != nil {
		t.Fatal(err)
	} else if n, ok := result.(*UsageNotificationNewUsage); !ok {
		t.Fatalf("unexpected type: %T", result)
	} else {
		assert.Equal(t, n, &UsageNotificationNewUsage{
			Account: Account{
				XMLName: xml.Name{Local: "account"},
				Code:    "923845792374",
			},
			Usage: Usage{
				XMLName:            xml.Name{Local: "usage"},
				ID:                 394729929104688227,
				SubscriptionUUID:   "35cda8d4ae0a214f69779e4ddbbc2ebd",
				AddOnCode:          "video_storage",
				MeasuredUnitID:     394681920153192422,
				Amount:             -40,
				RecordingTimestamp: recurly.NewTime(recordingTs),
				UsageTimestamp:     recurly.NewTime(usageTs),
				CreatedAt:          recurly.NewTime(createdTs),
				BilledAt:           recurly.NewTime(billedTs),
				UsageType:          "PRICE",
				UnitAmountInCents:  recurly.NewInt(50),
				UsagePercentage:    recurly.NewFloat(0),
			},
		})
	}
}
