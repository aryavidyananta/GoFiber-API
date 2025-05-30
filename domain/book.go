package domain

import (
	"aryavidyananta/Golang-Project/dto"
	"context"
	"database/sql"
)

type Book struct {
	Id        string         `db:"id"`
	Judul     string         `db:"judul"`
	Deskripsi string         `db:"deskripsi"`
	CoverId   sql.NullString `db:"cover_id"`
	CreatedAt sql.NullTime   `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
	DeletedAt sql.NullTime   `db:"deleted_at"`
}

type BookRepository interface {
	FindAll(ctx context.Context) ([]Book, error)
	FindById(ctx context.Context, id string) (Book, error)
	Save(ctx context.Context, book *Book) (Book, error)
	Update(ctx context.Context, book *Book) (Book, error)
	Delete(ctx context.Context, id string) error
}

type BookService interface {
	Index(ctx context.Context) ([]dto.BookData, error)
	Show(ctx context.Context, id string) (dto.BookData, error)
	Create(ctx context.Context, req dto.CreateBookRequest) error
	Update(ctx context.Context, req dto.UpdateBookRequest) error
	Delete(ctx context.Context, id string) error
}
