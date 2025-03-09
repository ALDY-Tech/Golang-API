package repository

import (
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/utils"

	"github.com/jmoiron/sqlx"
)

type CustomerRepository interface {
	Insert(customer *model.People) error
	FindAll(page int, totalRows int) ([]model.People, error)
	FindById(id string) (model.People, error)
	Update(customer *model.People) error
	Delete(id string) error
}

type customerRepository struct {
	db *sqlx.DB
}

func (c *customerRepository) Insert(customer *model.People) error {
	_, err := c.db.NamedExec(utils.INSERT_CUSTOMER, customer)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) FindAll(page int, totalRows int) ([]model.People, error) {
	limit := totalRows
	offset := (page - 1) * limit
	var customers []model.People

	err := c.db.Select(&customers, utils.SELECT_ALL_CUSTOMER, limit, offset)
	if err != nil {
		return nil, err
	}
	return customers, nil
}		

func (c *customerRepository) FindById(id string) (model.People, error) {
	var customer model.People
	err := c.db.Get(&customer, utils.SELECT_CUSTOMER_BY_ID, id)
	if err != nil {
		return model.People{}, err
	}
	return customer, nil
}

func (c *customerRepository) Update(customer *model.People) error {
	_, err := c.db.NamedExec(utils.UPDATE_CUSTOMER, customer)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) Delete(id string) error {
	_, err := c.db.Exec(utils.DELETE_CUSTOMER, id)
	if err != nil {
		return err
	}
	return nil
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo
}