package model

type Transaction struct {
	ID         string       `db:"id" json:"id"`
	BillDate   string       `db:"bill_date" json:"billDate"`
	EntryDate  string       `db:"entry_date" json:"entryDate"`
	FinishDate string       `db:"finish_date" json:"finishDate"`
	EmployeeID string       `db:"employee_id" json:"employeeId"`
	CustomerID string       `db:"customer_id" json:"customerId"`
	BillDetails []BillDetail `json:"billDetails"`
	TotalBill  int          `db:"total_bill" json:"totalBill"`
}


type BillDetail struct {
	ID           string `db:"id" json:"id"`
	BillID       string `db:"bill_id" json:"billId"`
	ProductID    string `db:"product_id" json:"productId"`
	ProductPrice int    `db:"product_price" json:"productPrice"`
	Qty          int    `db:"qty" json:"qty"`
}
//  list transaction
type Bill struct {
	ID         string `db:"id" json:"id"`
	BillID	 string `db:"bill_id" json:"billId"`
	Product Product `json:"product"`
	ProductPrice int `db:"product_price" json:"productPrice"`
	Qty 	 int    `db:"qty" json:"qty"`
}

type TransactionDetail struct {
	ID         string       `db:"id" json:"id"`
	BillDate   string       `db:"bill_date" json:"billDate"`
	EntryDate  string       `db:"entry_date" json:"entryDate"`
	FinishDate string       `db:"finish_date" json:"finishDate"`
	Employee People       `json:"employee"`
	Customer People       `json:"customer"`
	BillDetails []Bill `json:"billDetails"`
	TotalBill  int          `db:"total_bill" json:"totalBill"`
}
