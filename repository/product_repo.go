package repository

import (
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/utils"

	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Insert(product *model.Product) error
	FindAll(page int, totalRows int) ([]model.Product, error)
	FindById(id string) (model.Product, error)
	Update(product *model.Product) error
	Delete(id string) error
	Search(searchName string) ([]model.Product, error)
}

type productRepository struct {
	db *sqlx.DB	
}

func (p *productRepository) Insert(product *model.Product) error {
	_, err := p.db.NamedExec(utils.ProductQueries.Insert, product)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) FindAll(page int, totalRows int) ([]model.Product, error) {
	limit := totalRows
	offset := (page - 1) * limit
	var products []model.Product
	err := p.db.Select(&products, utils.ProductQueries.SelectAll, limit, offset)
	if err != nil {
		return nil, err
	}
	return products, nil
	}

func (p *productRepository) FindById(id string) (model.Product, error) {
	var product model.Product
	err := p.db.Get(&product, utils.ProductQueries.SelectByID, id)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (p *productRepository) Update(product *model.Product) error {
	_, err := p.db.NamedExec(utils.ProductQueries.Update, product)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) Delete(id string) error {
	_, err := p.db.Exec(utils.ProductQueries.Delete, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) Search(searchName string) ([]model.Product, error) {
	var products []model.Product
	err := p.db.Select(&products, utils.ProductQueries.SelectAll + utils.ProductQueries.Search, searchName)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}