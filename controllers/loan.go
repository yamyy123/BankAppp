package controllers

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoanController struct {
	loanservice interfaces.Iloan
}

func InitLoanController(loanservice interfaces.Iloan) LoanController {
	return LoanController{loanservice}
}

func (l *LoanController) CreateLoan(ctx *gin.Context) {
	var loan *models.Loan
	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	newloan, err := l.loanservice.CreateLoan(loan)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newloan})
}

func (l *LoanController) GetLoanById(ctx * gin.Context){
	id := ctx.Param("id")
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	val, err := l.loanservice.GetLoanById(id1)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}
func (l *LoanController)UpdateLoanById(ctx *gin.Context){
	id := ctx.Param("id")
	loan := &models.Loan{}
	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	val, err := l.loanservice.UpdateLoanById(id1, loan)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}
func (l *LoanController)DeleteLoanById(ctx *gin.Context){
     id := ctx.Param("id")
	 id1, err := strconv.ParseInt(id, 10, 64)
	 if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	 }
	 val, err := l.loanservice.DeleteLoanById(id1)
	 if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	 }
	 ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}



