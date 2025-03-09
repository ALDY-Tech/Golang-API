package usecase

import (
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/repository"
	"submission-project-enigma-laundry/utils"
)

type CustomerUseCase interface {
	CreateNewCustomer(customer *model.People) error
	GetAllCustomer(page int, totalRows int) ([]model.People, error)
	GetCustomerById(id string) (model.People, error)
	UpdateCustomer(customer *model.People) error
	DeleteCustomer(id string) error
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func (c *customerUseCase) CreateNewCustomer(customer *model.People) error {
	newCustomer.id = utils.GenerateUUID()
	return c.repo.Insert(newCustomer)
}

func (c *customerUseCase) GetAllCustomer(page int, totalRows int) ([]model.People, error) {
	return c.repo.FindAll(page, totalRows)
}

func (c *customerUseCase) GetCustomerById(id string) (model.People, error) {
	return c.repo.FindById(id)
}

func (c *customerUseCase) UpdateCustomer(customer *model.People) error {
	return c.repo.Update(customer)
}

func (c *customerUseCase) DeleteCustomer(id string) error {
	return c.repo.Delete(id)
}	

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	pc := new(customerUseCase)
	pc.repo = repo
	return pc
}