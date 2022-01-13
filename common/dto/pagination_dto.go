package dto

type PaginationDTO struct {
	Sort   string `url:"sort"`
	Order  string `url:"order"`
	Offset string `url:"offset"`
	Limit  string `url:"limit"`
	Search string `url:"search"`
}
