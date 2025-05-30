package service

import (
	"aryavidyananta/Golang-Project/domain"
	"aryavidyananta/Golang-Project/dto"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type BookService struct {
	BookRepository      domain.BookRepository
	BookStockRepository domain.BookStockRepository
}

func NewBook(BookRepository domain.BookRepository, BookStockRepository domain.BookStockRepository) domain.BookService {
	return &BookService{
		BookRepository:      BookRepository,
		BookStockRepository: BookStockRepository,
	}
}

// Index implements domain.BookService.
func (b *BookService) Index(ctx context.Context) ([]dto.BookData, error) {
	book, err := b.BookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var data []dto.BookData
	for _, v := range book {
		data = append(data, dto.BookData{
			Id:        v.Id,
			Judul:     v.Judul,
			Deskripsi: v.Deskripsi,
		})
	}
	return data, nil
}

// Create implements domain.BookService.
func (b *BookService) Create(ctx context.Context, req dto.CreateBookRequest) error {
	coverId := sql.NullString{Valid: false, String: req.CoverId}
	if req.CoverId != "" {
		coverId.Valid = true
	}

	book := domain.Book{
		Id:        uuid.NewString(),
		Judul:     req.Judul,
		Deskripsi: req.Deskripsi,
		CoverId:   coverId,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	_, err := b.BookRepository.Save(ctx, &book)
	return err
}

// Update implements domain.BookService.
func (b *BookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
	persited, err := b.BookRepository.FindById(ctx, req.Id)
	if err != nil {
		return err
	}
	if persited.Id == "" {
		return errors.New("book not found")
	}
	coverId := sql.NullString{Valid: false, String: req.CoverId}
	if req.CoverId != "" {
		coverId.Valid = true
	}

	persited.Judul = req.Judul
	persited.Deskripsi = req.Deskripsi
	persited.CoverId = coverId
	persited.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	_, err = b.BookRepository.Update(ctx, &persited)
	return err
}

// Delete implements domain.BookService.
func (b *BookService) Delete(ctx context.Context, id string) error {
	exits, err := b.BookRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if exits.Id == "" {
		return errors.New("book not found")
	}
	return b.BookRepository.Delete(ctx, exits.Id)
}

// Show implements domain.BookService.
func (b *BookService) Show(ctx context.Context, id string) (dto.BookData, error) {
	persited, err := b.BookRepository.FindById(ctx, id)
	if err != nil {
		return dto.BookData{}, err
	}
	if persited.Id == "" {
		return dto.BookData{}, errors.New("book not found")
	}
	return dto.BookData{
		Id:        persited.Id,
		Judul:     persited.Judul,
		Deskripsi: persited.Deskripsi,
		CoverId:   persited.CoverId.String,
	}, nil
}
