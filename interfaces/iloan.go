package interfaces

import (
	"bankDemo/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Iloan interface {
	CreateLoan(*models.Loan) (*mongo.InsertOneResult, error)
	GetLoanById(int64) (*models.Loan, error)
	UpdateLoanById(int64, *models.Loan) (*mongo.UpdateResult, error)
	DeleteLoanById(int64) (*mongo.DeleteResult, error)
}