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

func (c *CustomerController) CreateNewCustomer(ctx *gin.Context) {
	var newCustomer *model.People
	err := ctx.ShouldBindJSON(&newCustomer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.customerUseCase.CreateNewCustomer(newCustomer)
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Customer created successfully",
		})
	}
}

func (c *CustomerController) GetCustomerById(ctx *gin.Context) {
	customerId := ctx.Param("id")
	responseUc, err := c.customerUseCase.GetCustomerById(customerId)
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

func (c *CustomerController) GetAllCustomers(ctx *gin.Context) {
	page,_ := strconv.Atoi(ctx.Query("page"))
	totalRows,_ := strconv.Atoi(ctx.Query("totalRows"))
	if page == 0 || totalRows == 0 {
		page = 1
		totalRows = 5
	}
	customers, err := c.customerUseCase.GetAllCustomer(page, totalRows)
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

func (c *CustomerController) UpdateCustomer(ctx *gin.Context) {
    var updateCustomer model.People

    // Ambil ID dari URL
    id := ctx.Param("id")
    if id == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "ID is required in URL",
        })
        return
    }

    // Bind JSON ke struct
    if err := ctx.BindJSON(&updateCustomer); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid JSON data: " + err.Error(),
        })
        return
    }

    // Overwrite ID dari URL ke dalam struct
    if updateCustomer.Id != "" && updateCustomer.Id != id {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "ID in body does not match ID in URL",
        })
        return
    }

    // Pastikan ID dari URL digunakan
    updateCustomer.Id = id

    // Lanjutkan proses update
    if err := c.customerUseCase.UpdateCustomer(updateCustomer); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "message": "Failed to update customer: " + err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Customer updated successfully",
        "data": updateCustomer,
    })
}

func (c *CustomerController) DeleteCustomer(ctx *gin.Context) {
	customerId := ctx.Param("id")
	err := c.customerUseCase.DeleteCustomer(customerId)
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
	customer.POST("/:id", newCustController.CreateNewCustomer)
	customer.GET("/:id", newCustController.GetCustomerById)
	customer.GET("", newCustController.GetAllCustomers)
	customer.PUT("/:id", newCustController.UpdateCustomer)
	customer.DELETE("/:id", newCustController.DeleteCustomer)

	return &newCustController
}