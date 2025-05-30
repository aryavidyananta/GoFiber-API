package api

import (
	"aryavidyananta/Golang-Project/domain"
	"aryavidyananta/Golang-Project/dto"
	"aryavidyananta/Golang-Project/internal/util"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BookApi struct {
	BookService domain.BookService
}

func NewBook(app *fiber.App,
	BookService domain.BookService,
	autMid fiber.Handler) {

	ba := BookApi{
		BookService: BookService,
	}

	app.Get("/books", autMid, ba.Index)
	app.Post("/books", autMid, ba.Create)
	app.Put("/books/:id", autMid, ba.Update)
	app.Delete("/books/:id", autMid, ba.Delete)
	app.Get("/books/:id", autMid, ba.Show)

}

func (ba BookApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	res, err := ba.BookService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ba BookApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("Validation Failed", fails))
	}

	err := ba.BookService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).
		JSON(dto.CreateResponseSuccess(""))
}

func (ba BookApi) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("Validate Failed", fails))
	}
	req.Id = ctx.Params("id")
	err := ba.BookService.Update(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess(""))
}

func (ba BookApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := ba.BookService.Delete(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (ba BookApi) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	data, err := ba.BookService.Show(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess(data))
}
