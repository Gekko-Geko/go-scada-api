package models

type Device struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
	DataType  string `json:"data_type"`
	DataValue string `json:"data_value"`
	DataUnit  string `json:"data_unit"`
}
