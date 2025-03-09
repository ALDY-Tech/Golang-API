package model

type BillDetailCreate struct {
	ID           string `json:"id"`
	ProductId    string `json:"productId"`
	ProductPrice int    `json:"productPrice"`
	Qty          int    `json:"qty"`
}