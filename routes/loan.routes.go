package routes

import (
	"bankDemo/controllers"

	"github.com/gin-gonic/gin"
)

func LoanRoute(r *gin.Engine, loanController controllers.LoanController) {
	r.POST("/loan", loanController.CreateLoan)
	r.GET("/loan/:id", loanController.GetLoanById)
	r.PUT("/loan/:id", loanController.UpdateLoanById)
	r.DELETE("/loan/:id", loanController.DeleteLoanById)
}