package routes

import (
	"bankDemo/controllers"

	"github.com/gin-gonic/gin"
)

func BankRoute(r *gin.Engine, bankController controllers.BankController) {
	r.POST("/bank", bankController.CreateBank)
	r.GET("/bank/:id", bankController.GetBankById)
	r.PUT("/bank/:id", bankController.UpdateBankById)
	r.DELETE("/bank/:id", bankController.DeleteBankById)
	r.GET("/customerbank", bankController.GetAllCustomerBank)
	r.GET("/banktransaction/:id", bankController.GetAllBankTransaction)
	r.POST("/banktransaction", bankController.GetAllBankTransDate)
}