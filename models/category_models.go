package models

type CategoryResponse struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
}

type CategoryRequest struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
}
