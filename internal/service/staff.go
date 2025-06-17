package service

import (
	"aryavidyananta/Golang-Project/domain"
	"aryavidyananta/Golang-Project/dto"
	"aryavidyananta/Golang-Project/internal/config"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/google/uuid"
)

type StaffService struct {
	config          *config.Config
	StaffRepository domain.StaffRepository
}

func NewStaff(conf *config.Config, StaffReposiory domain.StaffRepository) domain.StaffService {
	return &StaffService{
		config:          conf,
		StaffRepository: StaffReposiory,
	}
}

// Index implements domain.StaffService.
func (s *StaffService) Index(ctx context.Context) ([]dto.StaffData, error) {
	staff, err := s.StaffRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var data []dto.StaffData
	for _, v := range staff {
		data = append(data, dto.StaffData{
			Id:      v.Id,
			Nama:    v.Nama,
			NIP:     v.NIP,
			Jabatan: v.Jabatan,
			Gambar:  v.Gambar,
		})
	}
	return data, nil
}

// Show implements domain.StaffService.
func (s *StaffService) Show(ctx context.Context, id string) (dto.StaffData, error) {
	persisted, err := s.StaffRepository.FindById(ctx, id)
	if err != nil {
		return dto.StaffData{}, err
	}
	if persisted.Id == "" {
		return dto.StaffData{}, errors.New("staff not found")
	}
	return dto.StaffData{
		Id:      persisted.Id,
		Nama:    persisted.Nama,
		NIP:     persisted.NIP,
		Jabatan: persisted.Jabatan,
		Gambar:  persisted.Gambar,
	}, nil
}

// Create implements domain.StaffService.
func (s *StaffService) Create(ctx context.Context, req dto.CreateStaffRequest) (dto.StaffData, error) {
	ext := filepath.Ext(req.Gambar.Filename)
	filename := uuid.NewString() + ext

	relativePath := path.Join("staff", filename)
	absolutePath := filepath.Join(s.config.Storage.BasePath, relativePath)

	if err := os.MkdirAll(filepath.Dir(absolutePath), os.ModePerm); err != nil {
		return dto.StaffData{}, fmt.Errorf("failed to create directory: %w", err)
	}

	src, err := req.Gambar.Open()
	if err != nil {
		return dto.StaffData{}, fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(absolutePath)
	if err != nil {
		return dto.StaffData{}, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return dto.StaffData{}, fmt.Errorf("failed to save image: %w", err)
	}

	staff := domain.Staff{
		Id:      uuid.NewString(),
		Nama:    req.Nama,
		NIP:     req.NIP,
		Jabatan: req.Jabatan,
		Gambar:  relativePath,
	}

	saved, err := s.StaffRepository.Save(ctx, &staff)
	if err != nil {
		return dto.StaffData{}, err
	}

	url := path.Join(s.config.Server.Asset, saved.Gambar)

	return dto.StaffData{
		Id:      saved.Id,
		Nama:    saved.Nama,
		NIP:     saved.NIP,
		Jabatan: saved.Jabatan,
		Gambar:  url,
	}, nil

}

// Update implements domain.StaffService.
func (s *StaffService) Update(ctx context.Context, req dto.UpdateStaffRequest) (dto.StaffData, error) {
	persited, err := s.StaffRepository.FindById(ctx, req.Id)
	if err != nil {
		return dto.StaffData{}, fmt.Errorf("failed to find staff: %w", err)
	}

	relativePath := persited.Gambar
	if req.Gambar != nil {
		ext := filepath.Ext(req.Gambar.Filename)
		filename := uuid.NewString() + ext
		relativePath = path.Join("staff", filename)
		absolutePath := filepath.Join(s.config.Storage.BasePath, relativePath)

		file, err := req.Gambar.Open()
		if err != nil {
			return dto.StaffData{}, fmt.Errorf("failed to open uploaded file: %w", err)
		}
		defer file.Close()

		out, err := os.Create(absolutePath)
		if err != nil {
			return dto.StaffData{}, fmt.Errorf("failed to create file: %w", err)
		}
		defer out.Close()

		if _, err := io.Copy(out, file); err != nil {
			return dto.StaffData{}, fmt.Errorf("failed to save image: %w", err)
		}
	}

	persited.Nama = req.Nama
	persited.NIP = req.NIP
	persited.Jabatan = req.Jabatan
	persited.Gambar = relativePath

	updateStaff, err := s.StaffRepository.Update(ctx, &persited)
	if err != nil {
		return dto.StaffData{}, fmt.Errorf("failed to update staff: %w", err)
	}

	return dto.StaffData{
		Id:      updateStaff.Id,
		Nama:    updateStaff.Nama,
		NIP:     updateStaff.NIP,
		Jabatan: updateStaff.Jabatan,
		Gambar:  updateStaff.Gambar,
	}, nil

}

// Delete implements domain.StaffService.
func (s *StaffService) Delete(ctx context.Context, id string) error {
	exits, err := s.StaffRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if exits.Id == "" {
		return errors.New("staff not found")
	}
	return s.StaffRepository.Delete(ctx, exits.Id)
}
