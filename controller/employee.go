package controller

import (
	"net/http"
	"strconv"
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/usecase"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	router          *gin.Engine
	employeeUseCase usecase.EmployeeUseCase
}

func (e *EmployeeController) CreateNewEmployee(ctx *gin.Context) {
	var newEmployee *model.People
	err := ctx.ShouldBindJSON(&newEmployee)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		e.employeeUseCase.CreateNewEmployee(newEmployee)
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Employee created successfully",
		})
	}
}

func (e *EmployeeController) GetAllEmployee(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	totalRows, _ := strconv.Atoi(ctx.Query("totalRows"))
	if page == 0 || totalRows == 0 {
		page = 1
		totalRows = 5
	}
	responseUc, err := e.employeeUseCase.GetAllEmployee(page, totalRows)
	if responseUc == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message":   "OK",
			"data":      responseUc,
			"page":      page,
			"totalRows": totalRows,
		})
	}
}

func (e *EmployeeController) GetEmployeeById(ctx *gin.Context) {
	id := ctx.Param("id")
	responseUc, err := e.employeeUseCase.GetEmployeeById(id)
	if (responseUc == model.People{}) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    responseUc,
		})
	}
}

func (e *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	var updateEmployee model.People

	 id := ctx.Param("id")
    if id == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "ID is required in URL",
        })
        return
    }
	
	if err := ctx.BindJSON(&updateEmployee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON data: " + err.Error(),
		})
		return
	}

	if updateEmployee.Id != "" && updateEmployee.Id != id {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID in body does not match ID in URL",
		})
		return
	}

	updateEmployee.Id = id

	if err := e.employeeUseCase.UpdateEmployee(updateEmployee); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update employee:" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Employee updated successfully",
		"data":    updateEmployee,
	})
}

func (e *EmployeeController) DeleteEmployee(ctx *gin.Context) {
	id := ctx.Param("id")
	err := e.employeeUseCase.DeleteEmployee(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Employee deleted successfully",
		})
	}
}

func NewEmployeeController(router *gin.Engine, employeeUseCase usecase.EmployeeUseCase) *EmployeeController {
	newEmployeController := EmployeeController{
		router:          router,
		employeeUseCase: employeeUseCase,
	}

	employee := router.Group("/employees")
	employee.POST("", newEmployeController.CreateNewEmployee)
	employee.GET("", newEmployeController.GetAllEmployee)
	employee.GET("/:id", newEmployeController.GetEmployeeById)
	employee.PUT("/:id", newEmployeController.UpdateEmployee)
	employee.DELETE("/:id", newEmployeController.DeleteEmployee)

	return &newEmployeController
}
