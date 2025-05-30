package dto

type CreateBookStockRequest struct {
	BookId string `json:"book_id" validate:"required"`
	// Codes represents a list of book stock codes to be processed or validated.
	Codes []string `json:"codes" validate:"required"`
}

type DeleteBookStockRequest struct {
	Codes []string
}
