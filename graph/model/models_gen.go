// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CheckoutInput struct {
	Products []*CheckoutProductInput `json:"products"`
}

type CheckoutProductInput struct {
	Sku string `json:"sku"`
	Qty int    `json:"qty"`
}

type Product struct {
	Sku   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
}

type ProductInput struct {
	Sku   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
}

type Transaction struct {
	ID       string               `json:"id"`
	SubTotal float64              `json:"subTotal"`
	Total    float64              `json:"total"`
	Discount float64              `json:"discount"`
	Details  []*TransactionDetail `json:"details"`
}

type TransactionDetail struct {
	Product *Product `json:"product"`
	Qty     int      `json:"qty"`
	Price   float64  `json:"price"`
}