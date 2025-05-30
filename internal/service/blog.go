package service

import (
	"aryavidyananta/Golang-Project/domain"
	"aryavidyananta/Golang-Project/dto"
	"aryavidyananta/Golang-Project/internal/config"
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/google/uuid"
)

type BlogService struct {
	config         *config.Config
	BlogRepository domain.BlogRepository
}

func NewBlog(conf *config.Config, BlogReposiory domain.BlogRepository) domain.BlogService {
	return &BlogService{
		config:         conf,
		BlogRepository: BlogReposiory,
	}
}

// Index implements domain.BlogService.
func (b *BlogService) Index(ctx context.Context) ([]dto.BlogData, error) {
	// TODO: implement actual index logic
	return []dto.BlogData{}, nil
}

// Show implements domain.BlogService.
func (b *BlogService) Show(ctx context.Context, id string) (dto.BlogData, error) {
	// TODO: implement actual show logic
	return dto.BlogData{}, nil
}

func (b *BlogService) Create(ctx context.Context, req dto.CreateBlogRequest) (dto.BlogData, error) {
	// Ekstensi dan generate filename unik
	ext := filepath.Ext(req.Gambar.Filename)
	filename := uuid.NewString() + ext

	// Path relatif dan absolut
	relativePath := path.Join("blog", filename)
	absolutePath := filepath.Join(b.config.Storage.BasePath, relativePath)

	// Pastikan direktori tujuan ada
	if err := os.MkdirAll(filepath.Dir(absolutePath), os.ModePerm); err != nil {
		return dto.BlogData{}, fmt.Errorf("failed to create directory: %w", err)
	}

	// Simpan file
	src, err := req.Gambar.Open()
	if err != nil {
		return dto.BlogData{}, fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(absolutePath)
	if err != nil {
		return dto.BlogData{}, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return dto.BlogData{}, fmt.Errorf("failed to save image: %w", err)
	}

	// Simpan ke DB
	blog := domain.Blog{
		Id:      uuid.NewString(),
		Title:   req.Title,
		Content: req.Content,
		Gambar:  relativePath,
	}
	saved, err := b.BlogRepository.Save(ctx, &blog)
	if err != nil {
		return dto.BlogData{}, err
	}

	// URL lengkap gambar
	url := path.Join(b.config.Server.Asset, saved.Gambar)

	return dto.BlogData{
		Id:      saved.Id,
		Title:   saved.Title,
		Content: saved.Content,
		Gambar:  url,
	}, nil
}

// Update implements domain.BlogService.
func (b *BlogService) Update(ctx context.Context, req dto.UpdateBlogRequest) (dto.BlogData, error) {
	// Ambil data blog lama
	persisted, err := b.BlogRepository.FindById(ctx, req.Id)
	if err != nil {
		return dto.BlogData{}, fmt.Errorf("failed to find blog: %w", err)
	}

	// Proses gambar jika ada yang baru di-upload
	relativePath := persisted.Gambar
	if req.Gambar != nil {
		ext := filepath.Ext(req.Gambar.Filename)
		filename := uuid.NewString() + ext
		relativePath = path.Join("blog", filename)
		absolutePath := filepath.Join(b.config.Storage.BasePath, relativePath)

		file, err := req.Gambar.Open()
		if err != nil {
			return dto.BlogData{}, fmt.Errorf("failed to open image: %w", err)
		}
		defer file.Close()

		out, err := os.Create(absolutePath)
		if err != nil {
			return dto.BlogData{}, fmt.Errorf("failed to create file: %w", err)
		}
		defer out.Close()

		if _, err := io.Copy(out, file); err != nil {
			return dto.BlogData{}, fmt.Errorf("failed to save image: %w", err)
		}

	}
	// Update data blog
	persisted.Title = req.Title
	persisted.Content = req.Content
	persisted.Gambar = relativePath

	updatedBlog, err := b.BlogRepository.Update(ctx, &persisted)
	if err != nil {
		return dto.BlogData{}, fmt.Errorf("failed to update blog: %w", err)
	}

	// Kembalikan data dalam bentuk DTO
	return dto.BlogData{
		Id:      updatedBlog.Id,
		Title:   updatedBlog.Title,
		Content: updatedBlog.Content,
		Gambar:  updatedBlog.Gambar,
	}, nil
}

// Delete implements domain.BlogService.
func (b *BlogService) Delete(ctx context.Context, id string) error {
	// TODO: implement actual delete logic
	return nil
}
