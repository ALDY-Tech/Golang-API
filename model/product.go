package model

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
	Unit  string `json:"unit" binding:"required"`
}