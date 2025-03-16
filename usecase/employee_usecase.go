package usecase

import (
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/repository"
	"submission-project-enigma-laundry/utils"
)

type EmployeeUseCase interface {
	CreateNewEmployee(newEmployee *model.People) error
	GetAllEmployee(page int, totalRows int) ([]model.People, error)
	GetEmployeeById(id string) (model.People, error)
	UpdateEmployee(employee model.People) error
	DeleteEmployee(id string) error
}

type employeeUseCase struct {
	repo repository.EmployeeRepository
}

func (e *employeeUseCase) CreateNewEmployee(newEmployee *model.People) error {
	newEmployee.Id = utils.GenerateID()
	return e.repo.Insert(newEmployee)
}

func (e *employeeUseCase) GetAllEmployee(page int, totalRows int) ([]model.People, error) {
	return e.repo.FindAll(page, totalRows)
}

func (e *employeeUseCase) GetEmployeeById(id string) (model.People, error) {
	return e.repo.FindById(id)
}

func (e *employeeUseCase) UpdateEmployee(employee model.People) error {
	return e.repo.Update(&employee)
}

func (e *employeeUseCase) DeleteEmployee(id string) error {
	return e.repo.Delete(id)
}

func NewEmployeeUseCase(repo repository.EmployeeRepository) EmployeeUseCase {
	em := new(employeeUseCase)
	em.repo = repo
	return em 
}