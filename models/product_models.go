package models

import "database/sql"

type ProductResponse struct {
	ID          int            `json:"id"`
	ProductName string         `json:"product_name"`
	Price       float64        `json:"price"`
	Category    sql.NullString `json:"category"`
	Description sql.NullString `json:"description"`
}

type ProductRequest struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Description *string `json:"description"`
}
