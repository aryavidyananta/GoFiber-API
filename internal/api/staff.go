package api

import (
	"aryavidyananta/Golang-Project/domain"
	"aryavidyananta/Golang-Project/dto"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type StaffApi struct {
	StaffService domain.StaffService
}

func NewStaff(app *fiber.App, StaffService domain.StaffService, autMid fiber.Handler) {
	sa := StaffApi{
		StaffService: StaffService,
	}
	staff := app.Group("/staff", autMid)
	staff.Get("/", sa.Index)
	staff.Get("/:id", sa.Show)
	staff.Post("/", sa.Create)
	staff.Put("/:id", sa.Update)
	staff.Delete("/:id", sa.Delete)
}

func (sa StaffApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	res, err := sa.StaffService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (sa StaffApi) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	id := ctx.Params("id")
	data, err := sa.StaffService.Show(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(data))
}

func (sa StaffApi) Create(ctx *fiber.Ctx) error {
	var req dto.CreateStaffRequest
	req.Nama = ctx.FormValue("nama")
	req.NIP = ctx.FormValue("nip")
	req.Jabatan = ctx.FormValue("jabatan")

	file, err := ctx.FormFile("gambar")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(err.Error()))
	}
	req.Gambar = file

	result, err := sa.StaffService.Create(ctx.Context(), req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(result))
}

func (sa StaffApi) Update(ctx *fiber.Ctx) error {
	var req dto.UpdateStaffRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(err.Error()))
	}
	req.Id = ctx.Params("id")
	result, err := sa.StaffService.Update(ctx.Context(), req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(result)
}

func (sa StaffApi) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := sa.StaffService.Delete(ctx.Context(), id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess("Staff deleted successfully"))
}
