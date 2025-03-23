package repository

import (
	"database/sql"
	"errors"
	"submission-project-enigma-laundry/model"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	Insert(transaction *model.Transaction) error
	// FindAll(page int, totalRows int) ([]model.Transaction, error)
	FindById(id string) (*model.Transaction, error)
}

type transactionRepository struct {
	db *sqlx.DB
}

func (t *transactionRepository) Insert(transaction *model.Transaction) error {
	tx, err := t.db.Beginx()
	if err != nil {
		return err
	}

	// Insert transaksi tanpa RETURNING id
	queryTransaction := `INSERT INTO transactions (id, bill_date, entry_date, finish_date, employee_id, customer_id, total_bill) 
		VALUES ($1, $2, $3, $4, $5, $6, $7);`
	_, err = tx.Exec(queryTransaction,
		transaction.ID,
		transaction.BillDate,
		transaction.EntryDate,
		transaction.FinishDate,
		transaction.EmployeeID,
		transaction.CustomerID,
		0, // totalBill awalnya 0, diperbarui nanti
	)

	if err != nil {
		tx.Rollback()
		return errors.New("Failed to insert transaction: " + err.Error())
	}

	// Insert bill details
	var totalBill int
	for i, detail := range transaction.BillDetails {
		var productPrice int

		// Ambil harga produk dari database
		err := tx.QueryRow("SELECT price FROM products WHERE id = $1", detail.ProductID).Scan(&productPrice)
		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {
				return errors.New("Product not found: " + detail.ProductID)
			} else {
				return errors.New("Failed to retrieve product price: " + err.Error())
			}
		}

		// Insert ke bill_details tanpa RETURNING id
		queryBillDetail := `INSERT INTO bill_details (bill_id, product_id, qty, product_price) VALUES ($1, $2, $3, $4);`
		_, err = tx.Exec(queryBillDetail, transaction.ID, detail.ProductID, detail.Qty, productPrice)
		if err != nil {
			tx.Rollback()
			return errors.New("Failed to insert bill detail: " + err.Error())
		}

		// Simpan harga produk di struct agar bisa dikembalikan ke client
		transaction.BillDetails[i].ProductPrice = productPrice

		// Hitung total bill
		totalBill += productPrice * detail.Qty
	}

	// Update total bill setelah semua detail dimasukkan
	queryUpdateTotalBill := `UPDATE transactions SET total_bill = $1 WHERE id = $2`
	_, err = tx.Exec(queryUpdateTotalBill, totalBill, transaction.ID)
	if err != nil {
		tx.Rollback()
		return errors.New("Failed to update total bill: " + err.Error())
	}
	transaction.TotalBill = totalBill

	// Commit transaksi
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (t *transactionRepository) FindById(id string) (*model.Transaction, error) {
	var transaction model.TransactionDetail
	query := `
	SELECT t.id, t.bill_date, t.entry_date, t.finish_date, t.total_bill,
		   e.id AS employee_id, e.name AS employee_name, e.phone_number AS employee_phone, e.address AS employee_address,
		   c.id AS customer_id, c.name AS customer_name, c.phone_number AS customer_phone, c.address AS customer_address,
		   bd.id AS bill_detail_id, bd.bill_id, 
		   p.id AS product_id, p.name AS product_name, p.price AS product_price, p.unit AS product_unit, 
		   bd.product_price, bd.qty
	FROM transactions t
	JOIN employees e ON t.employee_id = e.id
	JOIN customers c ON t.customer_id = c.id
	JOIN bill_details bd ON t.id = bd.bill_id
	JOIN products p ON bd.product_id = p.id
	WHERE t.id = $1
	`
	rows, err := t.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var billDetails []model.BillDetail
	for rows.Next() {
		var billDetail model.Bill
		var product model.Product

		err := rows.Scan( &transaction.ID, &transaction.BillDate, &transaction.EntryDate, &transaction.FinishDate, &transaction.TotalBill,
			&transaction.Employee.Id, &transaction.Employee.Name, &transaction.Employee.PhoneNumber, &transaction.Employee.Address,
			&transaction.Customer.Id, &transaction.Customer.Name, &transaction.Customer.PhoneNumber, &transaction.Customer.Address,
			&billDetail.ID, &billDetail.BillID,
			&product.Id, &product.Name, &product.Price, &product.Unit,
			&billDetail.ProductPrice, &billDetail.Qty,
		)
		if err != nil {
			return nil, err
		}
		billDetail.Product = product
		billDetails = append(billDetails, billDetail)
	}
	transaction.BillDetails = billDetails

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	repo := new(transactionRepository)
	repo.db = db
	return repo
}
