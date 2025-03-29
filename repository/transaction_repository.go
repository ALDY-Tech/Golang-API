package repository

import (
	"database/sql"
	"errors"
	"submission-project-enigma-laundry/model"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	Insert(transaction *model.Transaction) error
	FindAll(startDate string, endDate string, productName string) ([]model.TransactionDetail, error)
	FindById(id string) (model.TransactionDetail, error)
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

func (t *transactionRepository) FindById(id string) (model.TransactionDetail, error) {
	var transaction model.TransactionDetail
	query := `
	SELECT t.id, t.bill_date, t.entry_date, t.finish_date, t.total_bill,
		   e.id, e.name, e.phonenumber, e.address,
		   c.id, c.name, c.phonenumber, c.address,
		   bd.id, bd.bill_id, 
		   p.id, p.name, p.price, p.unit, 
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
		return transaction, err
	}
	defer rows.Close()

	var billDetails []model.Bill
	for rows.Next() {
		var billDetail model.Bill
		var product model.Product

		err := rows.Scan(
			&transaction.ID, &transaction.BillDate, &transaction.EntryDate, &transaction.FinishDate, &transaction.TotalBill,
			&transaction.Employee.Id, &transaction.Employee.Name, &transaction.Employee.PhoneNumber, &transaction.Employee.Address,
			&transaction.Customer.Id, &transaction.Customer.Name, &transaction.Customer.PhoneNumber, &transaction.Customer.Address,
			&billDetail.ID, &billDetail.BillID,
			&product.Id, &product.Name, &product.Price, &product.Unit,
			&billDetail.ProductPrice, &billDetail.Qty,
		)
		if err != nil {
			return transaction, err
		}
		billDetail.Product = product
		billDetails = append(billDetails, billDetail)
	}
	transaction.BillDetails = billDetails

	if err := rows.Err(); err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (t *transactionRepository) FindAll(startDate string, endDate string, productName string) ([]model.TransactionDetail, error) {
	var transactions []model.TransactionDetail

	query := `
		SELECT t.id, t.bill_date, t.entry_date, t.finish_date, t.employee_id, t.customer_id,t.total_bill
		FROM transactions t
		LEFT JOIN bill_details bd ON t.id = bd.bill_id
		LEFT JOIN products p ON bd.product_id = p.id
		WHERE 1=1`

	var params []interface{}

	if startDate != "" {
		query += " AND t.bill_date >= $1"
		params = append(params, startDate)
	}

	if endDate != "" {
		query += " AND t.bill_date <= $2"
		params = append(params, endDate)
	}

	if productName != "" {
		query += " AND p.name ILIKE '%' || $3 || '%'"
		params = append(params, productName)
	}

	rows, err := t.db.Query(query, params...)
	if err != nil {
		return transactions, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction model.TransactionDetail
		err := rows.Scan(&transaction.ID, &transaction.BillDate, &transaction.EntryDate, &transaction.FinishDate, &transaction.Employee.Id, &transaction.Customer.Id, &transaction.TotalBill)
		if err != nil {
			return transactions, err
		}
		err = t.db.QueryRow(`
			SELECT id, name, phonenumber, address
			FROM employees WHERE id = $1`, transaction.Employee.Id).Scan(
			&transaction.Employee.Id,
			&transaction.Employee.Name,
			&transaction.Employee.PhoneNumber,
			&transaction.Employee.Address,
		)
		if err != nil {
			return transactions, err
		}
		err = t.db.QueryRow(`
			SELECT id, name, phonenumber, address
			FROM customers WHERE id = $1`, transaction.Customer.Id).Scan(
			&transaction.Customer.Id,
			&transaction.Customer.Name,
			&transaction.Customer.PhoneNumber,
			&transaction.Customer.Address,
		)
		if err != nil {
			return transactions, err
		}

		detailsRows, err := t.db.Query(`
			SELECT bd.id, bd.bill_id, p.id, p.name, p.price, p.unit, bd.product_price, bd.qty
			FROM bill_details bd
			JOIN products p ON bd.product_id = p.id
			WHERE bd.bill_id = $1`, transaction.ID)
		if err != nil {
			return transactions, err
		}
		defer detailsRows.Close()

		for detailsRows.Next() {
			var billDetail model.Bill
			var product model.Product

			err := detailsRows.Scan(&billDetail.ID, &billDetail.BillID, &product.Id, &product.Name, &product.Price, &product.Unit, &billDetail.ProductPrice, &billDetail.Qty)
			if err != nil {
				return transactions, err
			}
			billDetail.Product = product
			transaction.BillDetails = append(transaction.BillDetails, billDetail)
		}
		transactions = append(transactions, transaction)

	}
	return transactions, nil
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	repo := new(transactionRepository)
	repo.db = db
	return repo
}
