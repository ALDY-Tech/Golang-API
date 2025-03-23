package repository

import (
	"submission-project-enigma-laundry/model"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	Insert(transaction *model.Transaction) error
	// FindAll(page int, totalRows int) ([]model.Transaction, error)
	// FindById(id string) (model.Transaction, error)
}

type transactionRepository struct {
	db *sqlx.DB
}

func (t *transactionRepository) Insert(transaction *model.Transaction) error {
	tx := t.db.MustBegin()

	// Insert transaction
	queryTransaction := `INSERT INTO transactions (bill_date, entry_date, finish_date, employee_id, customer_id) 
		VALUES (:bill_date, :entry_date, :finish_date, :employee_id, :customer_id) RETURNING id`
	err := tx.QueryRow(queryTransaction,
		transaction.BillDate, transaction.EntryDate, transaction.FinishDate,
		transaction.EmployeeID, transaction.CustomerID).Scan(&transaction.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert bill details
	for _, detail := range transaction.BillDetails {
		queryBillDetail := `INSERT INTO bill_details (bill_id, product_id, qty, product_price) 
			VALUES ($1, $2, $3, $4) RETURNING id`
		err = tx.QueryRow(queryBillDetail, transaction.ID, detail.ProductID, detail.Qty, detail.ProductPrice).Scan(&detail.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	repo := new(transactionRepository)
	repo.db = db
	return repo
}
