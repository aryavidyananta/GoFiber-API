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

type CustomerApi struct {
	CustomerService domain.CustomerService
}

func NewCustomer(app *fiber.App,
	customerService domain.CustomerService,
	auzMidd fiber.Handler) {
	ca := CustomerApi{
		CustomerService: customerService,
	}
	app.Get("/customers", auzMidd, ca.Index)
	app.Post("/customers", auzMidd, ca.Create)
	app.Put("/customers/:id", auzMidd, ca.Update)
	app.Delete("/customers/:id", auzMidd, ca.Delete)
	app.Get("/customers/:id", auzMidd, ca.Show)

}

func (ca CustomerApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	res, err := ca.CustomerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ca CustomerApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	err := ca.CustomerService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).
		JSON(dto.CreateResponseSuccess(""))
}

func (ca CustomerApi) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("validation failed", fails))
	}
	req.ID = ctx.Params("id")
	err := ca.CustomerService.Update(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess(""))

}

func (ca CustomerApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := ca.CustomerService.Delete(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (ca CustomerApi) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	id := ctx.Params("id")
	data, err := ca.CustomerService.Show(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess(data))
}
