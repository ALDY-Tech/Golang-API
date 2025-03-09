package model

type Transaction struct {
	ID         string       `json:"id"`
	BillDate   string       `json:"billDate"`
	EntryDate  string       `json:"entryDate"`
	FinishDate string       `json:"finishDate"`
	Employee   Employee     `json:"employee"`
	Customer   Customers     `json:"customer"`
	BillDetails []BillDetail `json:"billDetails"`
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
