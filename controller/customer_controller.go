package controller

import (
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	router 		*gin.Engine
	customerUseCase usecase.CustomerUseCase
}

func (pc *CustomerController) CreateNewCustomer(ctx *gin.Context) {
	var newCustomer *model.People
	err := ctx.ShouldBindJSON(&newCustomer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		pc.customerUseCase.CreateNewCustomer(newCustomer)
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Customer created successfully",
		})
	}
}

func (pc *CustomerController) GetCustomerById(ctx *gin.Context) {
	customerId := ctx.Param("id")
	responseUc, err := pc.customerUseCase.GetCustomerById(customerId)
	if (responseUc == model.People{}) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data": responseUc,
		})
	}
}

func (pc *CustomerController) GetAllCustomers(ctx *gin.Context) {
	page,_ := strconv.Atoi(ctx.Query("page"))
	totalRows,_ := strconv.Atoi(ctx.Query("totalRows"))
	if page == 0 || totalRows == 0 {
		page = 1
		totalRows = 5
	}
	customers, err := pc.customerUseCase.GetAllCustomer(page, totalRows)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data": customers,
			"page": page,
			"totalRows": totalRows,
		})
	}
}

func (pc *CustomerController) UpdateCustomer(ctx *gin.Context) {
	var updateCustomer *model.People
	err := ctx.BindJSON(&updateCustomer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		err := pc.customerUseCase.UpdateCustomer(*updateCustomer)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Customer updated successfully",
				"data": updateCustomer,
			})
		}
	}
}

func (pc *CustomerController) DeleteCustomer(ctx *gin.Context) {
	customerId := ctx.Param("id")
	err := pc.customerUseCase.DeleteCustomer(customerId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Customer deleted successfully",
		})
	}
}

func NewCustomerController(router *gin.Engine, customerUseCase usecase.CustomerUseCase) *CustomerController {
	newCustController := CustomerController{
		router: router,
		customerUseCase: customerUseCase,
	}

	customer := router.Group("/customers")
	customer.POST("", newCustController.CreateNewCustomer)
	customer.GET("/:id", newCustController.GetCustomerById)
	customer.GET("", newCustController.GetAllCustomers)
	customer.PUT("/:id", newCustController.UpdateCustomer)
	customer.DELETE("/:id", newCustController.DeleteCustomer)

	return &newCustController
}