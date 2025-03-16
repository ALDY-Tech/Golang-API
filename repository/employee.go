package repository

import (
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/utils"

	"github.com/jmoiron/sqlx"	
)

type EmployeeRepository interface { 
	Insert(employee *model.People) error
	FindAll(page int, totalRows int) ([]model.People, error)
	FindById(id string) (model.People, error)
	Update(employee *model.People) error
	Delete(id string) error
}

type employeeRepository struct {
	db *sqlx.DB
}

func (e *employeeRepository) Insert(employee *model.People) error {
	_, err := e.db.NamedExec(utils.EmployeeQueries.Insert, employee)
	if err != nil {
		return err
	}
	return nil
}

func (e *employeeRepository) FindAll(page int, totalRows int) ([]model.People, error) {
	limit := totalRows
	offset := (page - 1) * limit
	var employees []model.People
	err := e.db.Select(&employees, utils.EmployeeQueries.SelectAll, limit, offset)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (e *employeeRepository) FindById(id string) (model.People, error) {
	var employee model.People
	err := e.db.Get(&employee, utils.EmployeeQueries.SelectByID, id)
	if err != nil {
		return model.People{}, err
	}
	return employee, nil
}

func (e *employeeRepository) Update(employee *model.People) error {	
	_, err := e.db.NamedExec(utils.EmployeeQueries.Update, employee)
	if err != nil {
		return err
	}
	return nil
	}

func (e *employeeRepository) Delete(id string) error {
	_, err := e.db.Exec(utils.EmployeeQueries.Delete, id)
	if err != nil {
		return err
	}
	return nil
}	

func NewEmployeeRepository(db *sqlx.DB) EmployeeRepository {
	repo := new(employeeRepository)
	repo.db = db
	return repo
}
