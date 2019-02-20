package models

// Product definition
type Product struct {
	Model
	Code  string `json:"code"`
	Price uint   `json:"price"`
}
