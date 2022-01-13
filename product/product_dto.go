package product

import "time"

type ProductDTO struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   *time.Time    `json:"deleted_at"`
}

type PaginationProductData struct {
	TotalData    int64
	FilteredData int
	Data         []ProductDTO
}
