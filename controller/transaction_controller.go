package controller

import (
	"log"
	"net/http"
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	router           *gin.Engine
	transactionUseCase usecase.TransactionUseCase
}

func (tc *TransactionController) CreateTransaction(ctx *gin.Context) {
	var transaction model.Transaction

	// Bind JSON
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	// Insert transaction
	err := tc.transactionUseCase.Insert(&transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction", "details": err.Error()})
		log.Fatal(err)
		return
	}

	// Return response
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Transaction created",
		"data":    transaction,
	})
}

func (tc *TransactionController) FindTransactionById(ctx *gin.Context) {
	id := ctx.Param("id")
	transaction, err := tc.transactionUseCase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    transaction,
	})
}

func (tc *TransactionController) FindAllTransactions(ctx *gin.Context) {
	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")
	productName := ctx.Query("productName")

	transactions, err := tc.transactionUseCase.FindAll(startDate, endDate, productName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    transactions,
	})
	}
func NewTransactionController(router *gin.Engine, transactionUseCase usecase.TransactionUseCase) *TransactionController {
	newTransactionController := TransactionController{
		router:           router,
		transactionUseCase: transactionUseCase,
	}

	transaction := router.Group("/transactions")
	transaction.POST("/", newTransactionController.CreateTransaction)
	transaction.GET("/:id", newTransactionController.FindTransactionById)
	transaction.GET("/", newTransactionController.FindAllTransactions)

	return &newTransactionController
}

