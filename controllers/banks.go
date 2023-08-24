package controllers

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BankController struct {
	BankService interfaces.IBank
}

type Date struct{
	From string `json:"from"`
	To string `json:"to"`
}

func InitBankController(bankService interfaces.IBank) BankController {
	return BankController{bankService}
}

func (b *BankController) CreateBank(ctx *gin.Context) {
	var banks *models.Bank
	if err := ctx.ShouldBindJSON(&banks); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	newbank, err := b.BankService.CreateBank(banks)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newbank})
}

func (b *BankController) GetBankById(ctx *gin.Context){
	id := ctx.Param("id")
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	val, err := b.BankService.GetBankById(id1)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}

func (b *BankController) UpdateBankById(ctx *gin.Context) {
	id := ctx.Param("id")
	bank := &models.Bank{}
	if err := ctx.ShouldBindJSON(&bank); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	val, err := b.BankService.UpdateBankById(id1, bank)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}

func (b *BankController) DeleteBankById(ctx *gin.Context){
	id := ctx.Param("id")
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	val, err := b.BankService.DeleteBankById(id1)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}
func (b *BankController) GetAllCustomerBank(ctx *gin.Context){
    res,err := b.BankService.GetAllCustomerBank()
    if err!=nil{
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (b *BankController) GetAllBankTransaction(ctx *gin.Context){
	id := ctx.Param("id")
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	res,err := b.BankService.GetAllBankTransaction(id1)
	if err!=nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (b *BankController) GetAllBankTransDate(ctx *gin.Context){
	var date *Date
	if err := ctx.ShouldBindJSON(&date); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	res,err := b.BankService.GetAllBankTransDate(date.From,date.To)
	if err!=nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

