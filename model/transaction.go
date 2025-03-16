package model

type Transaction struct {
	ID         string       `json:"id"`
	BillDate   string       `json:"billDate"`
	EntryDate  string       `json:"entryDate"`
	FinishDate string       `json:"finishDate"`
	Employee   People     `json:"employee"`
	Customer   People     `json:"customer"`
	BillDetails []BillDetail `json:"billDetails"`
	TotalBill  int          `json:"totalBill"`
}

type BillDetail struct {
	ID           string `json:"id"`
	BillID string `json:"billId"`
	Product Product `json:"product"`
	ProductPrice int `json:"productPrice"`
	Qty          int    `json:"qty"`
}
