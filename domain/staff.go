package domain

import (
	"aryavidyananta/Golang-Project/dto"
	"context"
)

type Staff struct {
	Id      string `db:"id"`
	Nama    string `db:"nama"`
	NIP     string `db:"nip"`
	Jabatan string `db:"jabatan"`
	Gambar  string `db:"gambar"`
}

type StaffRepository interface {
	FindAll(ctx context.Context) ([]Staff, error)
	FindById(ctx context.Context, id string) (Staff, error)
	Save(ctx context.Context, staff *Staff) (Staff, error)
	Update(ctx context.Context, staff *Staff) (Staff, error)
	Delete(ctx context.Context, id string) error
}

type StaffService interface {
	Index(ctx context.Context) ([]dto.StaffData, error)
	Show(ctx context.Context, id string) (dto.StaffData, error)
	Create(ctx context.Context, req dto.CreateStaffRequest) (dto.StaffData, error)
	Update(ctx context.Context, req dto.UpdateStaffRequest) (dto.StaffData, error)
	Delete(ctx context.Context, id string) error
}
