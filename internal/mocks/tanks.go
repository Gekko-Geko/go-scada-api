package mocks

import (
	"time"
	"scada-api/internal/models"
)

func GetMockTanks() []models.Device {
	return []models.Device{
		{ID: "abc-def", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Percentage", DataValue: "34", DataUnit: "%"},
		{ID: "gjh-khl", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Percentage", DataValue: "71", DataUnit: "%"},
		{ID: "poi-asa", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Percentage", DataValue: "31", DataUnit: "%"},
		{ID: "szx-daa", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Percentage", DataValue: "57", DataUnit: "%"},
		{ID: "yic-yux", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Percentage", DataValue: "78", DataUnit: "%"},
		{ID: "aos-gyx", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Percentage", DataValue: "15", DataUnit: "%"},
		{ID: "oiz-azx", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Percentage", DataValue: "58", DataUnit: "%"},
		{ID: "gnb-gyl", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Percentage", DataValue: "78", DataUnit: "%"},
	}
}
