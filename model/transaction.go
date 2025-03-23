package model

type Transaction struct {
	ID         string       `json:"id"`
	BillDate   string       `json:"billDate"`
	EntryDate  string       `json:"entryDate"`
	FinishDate string       `json:"finishDate"`
	EmployeeID string       `json:"employeeId"`
	CustomerID string       `json:"customerId"`
	BillDetails []BillDetail `json:"billDetails"`
	TotalBill  int          `json:"totalBill"`
}

type BillDetail struct {
	ID           string `json:"id"`
	BillID       string `json:"billId"`
	ProductID    string `json:"productId"`
	ProductPrice int    `json:"productPrice"`
	Qty          int    `json:"qty"`
}
