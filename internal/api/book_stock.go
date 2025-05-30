package api

import (
	"aryavidyananta/Golang-Project/domain"
	"aryavidyananta/Golang-Project/dto"
	"aryavidyananta/Golang-Project/internal/util"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type bookStockApi struct {
	bookStockService domain.BookStockService
}

func NewBookStock(app *fiber.App,
	bookStockService domain.BookStockService,
	autMid fiber.Handler) {
	bsa := bookStockApi{
		bookStockService: bookStockService,
	}

	app.Post("/bookstocks", autMid, bsa.Create)
	app.Delete("/bookstocks", autMid, bsa.Delete)
}

func (ba bookStockApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateBookStockRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("Validation Failed", fails))
	}
	err := ba.bookStockService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(""))
}

func (ba bookStockApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	codes := strings.Split(ctx.Query("code"), ",")
	if len(codes) < 1 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseError("Parameter Code Wajib Diisi"))
	}
	err := ba.bookStockService.Delete(c, dto.DeleteBookStockRequest{Codes: codes})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.SendStatus(http.StatusNoContent)

}
