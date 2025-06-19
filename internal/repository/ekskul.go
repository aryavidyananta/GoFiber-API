package repository

import (
	"aryavidyananta/Golang-Project/domain"
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type EkskulRepository struct {
	db *goqu.Database
}

func NewEkskul(con *sql.DB) domain.EkskulRepository {
	return &EkskulRepository{
		db: goqu.New("default", con),
	}
}

// FindAll implements domain.EkskulRepository.
func (e *EkskulRepository) FindAll(ctx context.Context) (result []domain.Ekskul, err error) {
	dataset := e.db.From("ekskul").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

// FindById implements domain.EkskulRepository.
func (e *EkskulRepository) FindById(ctx context.Context, id string) (result domain.Ekskul, err error) {
	dataset := e.db.From("ekskul").Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// Save implements domain.EkskulRepository.
func (e *EkskulRepository) Save(ctx context.Context, ekskul *domain.Ekskul) (domain.Ekskul, error) {
	dataset := e.db.Insert("ekskul").Rows(ekskul).Executor()
	_, err := dataset.ExecContext(ctx)
	if err != nil {
		return domain.Ekskul{}, err
	}
	return *ekskul, nil

}

// Update implements domain.EkskulRepository.
func (e *EkskulRepository) Update(ctx context.Context, ekskul *domain.Ekskul) (domain.Ekskul, error) {
	executor := e.db.Update("ekskul").Where(goqu.C("id").Eq(ekskul.Id)).Set(ekskul).Executor()
	_, err := executor.ExecContext(ctx)
	if err != nil {
		return domain.Ekskul{}, err
	}
	return *ekskul, nil
}

// Delete implements domain.EkskulRepository.
func (e *EkskulRepository) Delete(ctx context.Context, id string) error {
	executor := e.db.Update("ekskul").Where(goqu.C("id").Eq(id)).Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
