package api

import (
	"aryavidyananta/Golang-Project/domain"
	"aryavidyananta/Golang-Project/dto"
	"aryavidyananta/Golang-Project/internal/config"
	"context"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type mediaAPI struct {
	cnf          *config.Config
	mediaService domain.MediaService
}

func NewMedia(app *fiber.App,
	cnf *config.Config,
	MediaService domain.MediaService,
	autMid fiber.Handler) {
	ma := mediaAPI{
		cnf:          cnf,
		mediaService: MediaService,
	}

	app.Post("/media", autMid, ma.Create)
	app.Static("/media", cnf.Storage.BasePath)

}

func (ma mediaAPI) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	file, err := ctx.FormFile("media")
	if err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}
	filename := uuid.NewString() + file.Filename
	path := filepath.Join(ma.cnf.Storage.BasePath, filename)
	err = ctx.SaveFile(file, path)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	res, err := ma.mediaService.Create(c, dto.CreateMediaRequest{
		Path: filename,
	})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).
		JSON(dto.CreateResponseSuccess(res))
}
