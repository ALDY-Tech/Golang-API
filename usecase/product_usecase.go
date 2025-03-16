package usecase

import (
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/repository"
	"submission-project-enigma-laundry/utils"
)

type ProductUseCase interface {
	CreateNewProduct(newProduct *model.Product) error
	GetAllProduct(searchName string, page int, totalRows int) ([]model.Product, error)
	GetProductById(id string) (model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(id string) error
}

type productUseCase struct {
	productRepository repository.ProductRepository
}

func (p *productUseCase) CreateNewProduct(newProduct *model.Product) error {
	newProduct.Id = utils.GenerateID()
	return p.productRepository.Insert(newProduct)
}

func (p *productUseCase) GetAllProduct( searchName string, page int, totalRows int) ([]model.Product, error) {
	return p.productRepository.FindAll(searchName,page, totalRows)
}

func (p *productUseCase) GetProductById(id string) (model.Product, error) {
	return p.productRepository.FindById(id)
}

func (p *productUseCase) UpdateProduct(product *model.Product) error {
	return p.productRepository.Update(product)
}

func (p *productUseCase) DeleteProduct(id string) error {
	return p.productRepository.Delete(id)
}

func NewProductUseCase(productRepository repository.ProductRepository) ProductUseCase {
	pu := new(productUseCase)
	pu.productRepository = productRepository
	return pu
}