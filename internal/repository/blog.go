package repository

import (
	"aryavidyananta/Golang-Project/domain"
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type BlogRepository struct {
	db *goqu.Database
}

func NewBlog(con *sql.DB) domain.BlogRepository {
	return &BlogRepository{
		db: goqu.New("default", con),
	}
}

// FindAll implements domain.BlogRepository.
func (b *BlogRepository) FindAll(ctx context.Context) (result []domain.Blog, err error) {
	dataset := b.db.From("blogs").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

// FindById implements domain.BlogRepository.
func (b *BlogRepository) FindById(ctx context.Context, id string) (result domain.Blog, err error) {
	dataset := b.db.From("blogs").Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// FindByids implements domain.BlogRepository.
func (b *BlogRepository) FindByids(ctx context.Context, ids []string) (result []domain.Blog, err error) {
	dataset := b.db.From("blogs").Where(goqu.C("deleted_at").IsNull(), goqu.C("id").In(ids))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// Save implements domain.BlogRepository.
func (b *BlogRepository) Save(ctx context.Context, blog *domain.Blog) (domain.Blog, error) {
	executor := b.db.Insert("blogs").Rows(blog).Executor()
	_, err := executor.ExecContext(ctx)
	if err != nil {
		return domain.Blog{}, err
	}
	return *blog, nil
}

// Update implements domain.BlogRepository.
func (b *BlogRepository) Update(ctx context.Context, blog *domain.Blog) (domain.Blog, error) {
	executor := b.db.Update("blogs").Where(goqu.C("id").Eq(blog.Id)).Set(blog).Executor()
	_, err := executor.ExecContext(ctx)
	if err != nil {
		return domain.Blog{}, err
	}
	return *blog, nil
}

// Delete implements domain.BlogRepository.
func (b *BlogRepository) Delete(ctx context.Context, id string) error {
	executor := b.db.Update("blogs").
		Where(goqu.C("id").Eq(id)).
		Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
