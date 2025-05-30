package service

import (
	"aryavidyananta/Golang-Project/domain"
	"aryavidyananta/Golang-Project/dto"
	"context"
)

type bookStockService struct {
	BookRepository      domain.BookRepository
	BookStockRepository domain.BookStockRepository
}

func NewBookStock(bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository) domain.BookStockService {
	return &bookStockService{
		BookRepository:      bookRepository,
		BookStockRepository: bookStockRepository,
	}
}

// Create implements domain.BookStockService.
func (b *bookStockService) Create(ctx context.Context, req dto.CreateBookStockRequest) error {
	book, err := b.BookRepository.FindById(ctx, req.BookId)
	if err != nil {
		return err
	}
	if book.Id == "" {
		return domain.BookNotFound
	}

	stocks := make([]domain.BookStock, 0)
	for _, v := range req.Codes {
		stocks = append(stocks, domain.BookStock{
			Code:   v,
			BookId: req.BookId,
			Status: domain.BookStockStatusAvailable,
		})
	}
	return b.BookStockRepository.Save(ctx, stocks)
}

// Delete implements domain.BookStockService.
func (b *bookStockService) Delete(ctx context.Context, req dto.DeleteBookStockRequest) error {
	return b.BookStockRepository.DeleteByCodes(ctx, req.Codes)
}
