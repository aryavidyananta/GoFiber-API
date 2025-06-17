package repository

import (
	"aryavidyananta/Golang-Project/domain"
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type StaffRepository struct {
	db *goqu.Database
}

func NewStaff(con *sql.DB) domain.StaffRepository {
	return &StaffRepository{
		db: goqu.New("default", con),
	}
}

// FindAll implements domain.StaffRepository.
func (s *StaffRepository) FindAll(ctx context.Context) (result []domain.Staff, err error) {
	dataset := s.db.From("staff").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

// FindById implements domain.StaffRepository.
func (s *StaffRepository) FindById(ctx context.Context, id string) (result domain.Staff, err error) {
	dataset := s.db.From("staff").Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// Save implements domain.StaffRepository.
func (s *StaffRepository) Save(ctx context.Context, staff *domain.Staff) (domain.Staff, error) {
	dataset := s.db.Insert("staff").Rows(staff).Executor()
	_, err := dataset.ExecContext(ctx)
	if err != nil {
		return domain.Staff{}, err
	}
	return *staff, nil
}

// Update implements domain.StaffRepository.
func (s *StaffRepository) Update(ctx context.Context, staff *domain.Staff) (domain.Staff, error) {
	executor := s.db.Update("staff").Where(goqu.C("id").Eq(staff.Id)).Set(staff).Executor()
	_, err := executor.ExecContext(ctx)
	if err != nil {
		return domain.Staff{}, err
	}
	return *staff, nil
}

// Delete implements domain.StaffRepository.
func (s *StaffRepository) Delete(ctx context.Context, id string) error {
	executor := s.db.Update("staff").Where(goqu.C("id").Eq(id)).Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
