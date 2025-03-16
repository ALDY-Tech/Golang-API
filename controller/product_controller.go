package controller

import (
	"net/http"
	"strconv"
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/usecase"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router 		*gin.Engine
	productUseCase usecase.ProductUseCase
}

func (pc *ProductController) CreateNewProduct(ctx *gin.Context) {
	var newProduct *model.Product
	err := ctx.ShouldBindJSON(&newProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		pc.productUseCase.CreateNewProduct(newProduct)
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Product created successfully",
		})
	}
}

func (pc *ProductController) GetAllProduct(ctx *gin.Context) {
	searchName := ctx.Query("name")
	page, _ := strconv.Atoi(ctx.Query("page"))
	totalRows, _ := strconv.Atoi(ctx.Query("totalRows"))
	if page == 0 || totalRows == 0 {
		page = 1
		totalRows = 5
	}
	responseUc, err := pc.productUseCase.GetAllProduct(searchName, page, totalRows)
	if err != nil {
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

func (pc *ProductController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	responseUc, err := pc.productUseCase.GetProductById(id)
	if (responseUc == model.Product{}) {
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

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	var updateProduct model.Product 

	 id := ctx.Param("id")
	 if id == "" {
		 ctx.JSON(http.StatusBadRequest, gin.H{
			 "message": "ID is required",
		 })
		 	return
	} 

	if err := ctx.BindJSON(&updateProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON data: " + err.Error(),
		})
		return
	}

	if updateProduct.Id != "" && updateProduct.Id != id {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID in URL and body must be same",
		})
		return
	}

	updateProduct.Id = id

	if err := pc.productUseCase.UpdateProduct(&updateProduct); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update product: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"data":    updateProduct,
	})
}

func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	err := pc.productUseCase.DeleteProduct(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Product deleted successfully",
		})
	}
}

func NewProductController(router *gin.Engine, productUseCase usecase.ProductUseCase) *ProductController {
	newProductController := ProductController{
		router: router,
		productUseCase: productUseCase,
	}

	product := router.Group("/products")
	product.POST("/", newProductController.CreateNewProduct)
	product.GET("", newProductController.GetAllProduct)
	product.GET("/:id", newProductController.GetProductById)	
	product.PUT("/:id", newProductController.UpdateProduct)
	product.DELETE("/:id", newProductController.DeleteProduct)

	return &newProductController
}