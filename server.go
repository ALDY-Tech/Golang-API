package main 

import (
	"submission-project-enigma-laundry/config"
	"submission-project-enigma-laundry/controller"
	"submission-project-enigma-laundry/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func Server() *appServer {
	ginEngine := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUseCaseManager(repo)
	return &appServer{
		useCaseManager: usecase,
		engine:         ginEngine,
		host:           "localhost:8080",
	}
}

func (a *appServer) initHandlers() {
	controller.NewCustomerController(a.engine, a.useCaseManager.CustomerUseCase())
	controller.NewEmployeeController(a.engine, a.useCaseManager.EmployeeUseCase())
	controller.NewProductController(a.engine, a.useCaseManager.ProductUsecase())
	controller.NewTransactionController(a.engine, a.useCaseManager.TransactionUsecase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
