package models

type UpdateModel struct {
	Topic string `json:"topic"`
	FinalValue  interface{} `json:"final_value"`
}