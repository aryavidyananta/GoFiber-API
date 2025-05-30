package repository

import (
	"aryavidyananta/Golang-Project/domain"
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type BookRepository struct {
	db *goqu.Database
}

func NewBook(con *sql.DB) domain.BookRepository {
	return &BookRepository{
		db: goqu.New("default", con),
	}
}

// FindAll implements domain.BookRepository.
func (b *BookRepository) FindAll(ctx context.Context) (result []domain.Book, err error) {
	dataset := b.db.From("book").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

// FindById implements domain.BookRepository.
func (b *BookRepository) FindById(ctx context.Context, id string) (result domain.Book, err error) {
	dataset := b.db.From("book").Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// Save implements domain.BookRepository.
func (b *BookRepository) Save(ctx context.Context, book *domain.Book) (domain.Book, error) {
	executor := b.db.Insert("book").Rows(book).Executor()
	_, err := executor.ExecContext(ctx)
	if err != nil {
		return domain.Book{}, err

	}
	return *book, nil
}

// Update implements domain.BookRepository.
func (b *BookRepository) Update(ctx context.Context, book *domain.Book) (domain.Book, error) {
	executor := b.db.Update("book").Where(goqu.C("id").Eq(book.Id)).Set(book).Executor()
	_, err := executor.ExecContext(ctx)
	if err != nil {
		return domain.Book{}, err
	}
	return *book, nil
}

// Delete implements domain.BookRepository.
func (b *BookRepository) Delete(ctx context.Context, id string) error {
	executor := b.db.Update("book").
		Where(goqu.C("id").Eq(id)).
		Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).
		Executor()

	_, err := executor.ExecContext(ctx)
	return err
}
