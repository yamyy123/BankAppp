package interfaces

import (
	"bankDemo/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Iaccount interface {
	CreateAccount(*models.Account)(*mongo.InsertOneResult,error)
	GetAccountById(int64) (*models.Account, error)
	UpdateAccountById(int64, *models.Account) (*mongo.UpdateResult, error)
	DeleteAccountById(int64) (*mongo.DeleteResult, error)
}