package domain

import (
	"aryavidyananta/Golang-Project/dto"
	"context"
)

type Blog struct {
	Id      string `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Gambar  string `db:"gambar"`
}

type BlogRepository interface {
	FindAll(ctx context.Context) ([]Blog, error)
	FindById(ctx context.Context, id string) (Blog, error)
	FindByids(ctx context.Context, ids []string) ([]Blog, error)
	Save(ctx context.Context, blog *Blog) (Blog, error)
	Update(ctx context.Context, blog *Blog) (Blog, error)
	Delete(ctx context.Context, id string) error
}

type BlogService interface {
	Index(ctx context.Context) ([]dto.BlogData, error)
	Show(ctx context.Context, id string) (dto.BlogData, error)
	Create(ctx context.Context, req dto.CreateBlogRequest) (dto.BlogData, error) // ubah: kembalikan BlogData
	Update(ctx context.Context, req dto.UpdateBlogRequest) (dto.BlogData, error) // ubah: kembalikan BlogData
	Delete(ctx context.Context, id string) error
}
