package mocks

import (
	"time"
	"scada-api/internal/models"
)

func GetMockMotors() []models.Device {
	return []models.Device{
		{ID: "abc-def", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Temp.", DataValue: "35", DataUnit: "C"},
		{ID: "qwe-asd", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Temp.", DataValue: "41", DataUnit: "C"},
		{ID: "jlk-uio", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Temp.", DataValue: "52", DataUnit: "C"},
		{ID: "gjh-oiu", Timestamp: time.Now().Format(time.RFC3339), Status: "running", DataType: "Temp.", DataValue: "39", DataUnit: "C"},
	}
}
