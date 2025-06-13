package mocks

import (
	"time"
	"scada-api/internal/models"
)

func GetMockPumps() []models.Device {
	return []models.Device{
		{ID: "abc-def", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Temp.", DataValue: "21", DataUnit: "C"},
		{ID: "kjl-xcy", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Temp.", DataValue: "28", DataUnit: "C"},
		{ID: "gha-dca", Timestamp: time.Now().Format(time.RFC3339), Status: "stopped", DataType: "Temp.", DataValue: "23", DataUnit: "C"},
		{ID: "gyi-ssw", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Temp.", DataValue: "20", DataUnit: "C"},
	}
}
