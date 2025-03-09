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

func (pc *CustomerController) CreateNewCustomer(c *gin.Context) {
	