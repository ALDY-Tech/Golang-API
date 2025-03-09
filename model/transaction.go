package model

type Transaction struct {
	ID         string       `json:"id"`
	BillDate   string       `json:"billDate"`
	EntryDate  string       `json:"entryDate"`
	FinishDate string       `json:"finishDate"`
	Employee   People     `json:"employee"`
	Customer   People     `json:"customer"`
	BillDetails []BillDetailCreate `json:"billDetails"`
	TotalBill  int          `json:"totalBill"`
}

type TransactionCreate struct {
	BillDate   string       `json:"billDate"`
	EntryDate  string       `json:"entryDate"`
	FinishDate string       `json:"finishDate"`
	EmployeeId   string     `json:"employeeId"`
	CustomerId   string     `json:"customerId"`
	BillDetails []BillDetailCreate `json:"billDetails"`
}
