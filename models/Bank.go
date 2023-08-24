package models

type Bank struct {
	Bank_ID int64  `json:"bank_id" bson:"bank_id"`
	IFSC string `json:"ifsc" bson:"ifsc"`
	BankAddress string `json:"bank_address" bson:"bank_address"`
}