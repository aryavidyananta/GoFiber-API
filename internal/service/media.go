package service

import (
	"aryavidyananta/Golang-Project/domain"
	"aryavidyananta/Golang-Project/dto"
	"aryavidyananta/Golang-Project/internal/config"
	"context"
	"database/sql"
	"path"
	"time"

	"github.com/google/uuid"
)

type mediaService struct {
	conf            *config.Config
	MediaRepository domain.MediaRepository
}

func NewMedia(conf *config.Config, MediaReposiory domain.MediaRepository) domain.MediaService {
	return &mediaService{
		conf:            conf,
		MediaRepository: MediaReposiory,
	}
}

// Create implements domain.MediaService.
func (m *mediaService) Create(ctx context.Context, req dto.CreateMediaRequest) (dto.MediaData, error) {
	media := domain.Media{
		Id:        uuid.NewString(),
		Path:      req.Path,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	err := m.MediaRepository.Save(ctx, &media)
	if err != nil {
		return dto.MediaData{}, err
	}
	url := path.Join(m.conf.Server.Asset, media.Path)
	return dto.MediaData{
		Id:   media.Id,
		Path: media.Path,
		Url:  url,
	}, nil
}
