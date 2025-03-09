package model

type TransactionCreate struct {
	BillDate   string       `json:"billDate"`
	EntryDate  string       `json:"entryDate"`
	FinishDate string       `json:"finishDate"`
	EmployeeId   string     `json:"employeeId"`
	CustomerId   string     `json:"customerId"`
	BillDetails []BillDetailCreate `json:"billDetails"`
}

type BillDetailCreate struct {
	ID           string `json:"id"`
	ProductId    string `json:"productId"`
	ProductPrice int    `json:"productPrice"`
	Qty          int    `json:"qty"`
}