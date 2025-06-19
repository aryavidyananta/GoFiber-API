package domain

import (
	"aryavidyananta/Golang-Project/dto"
	"context"
)

type Ekskul struct {
	Id     string `db:"id"`
	Nama   string `db:"nama"`
	Gambar string `db:"gambar"`
}

type EkskulRepository interface {
	FindAll(ctx context.Context) ([]Ekskul, error)
	FindById(ctx context.Context, id string) (Ekskul, error)
	Save(ctx context.Context, ekskul *Ekskul) (Ekskul, error)
	Update(ctx context.Context, ekskul *Ekskul) (Ekskul, error)
	Delete(ctx context.Context, id string) error
}

type EkskulService interface {
	Index(ctx context.Context) ([]dto.EkskulData, error)
	Show(ctx context.Context, id string) (dto.EkskulData, error)
	Create(ctx context.Context, req dto.CreateEkskulRequest) (dto.EkskulData, error)
	Update(ctx context.Context, req dto.UpdateEkskulRequest) (dto.EkskulData, error)
	Delete(ctx context.Context, id string) error
}
