package api

import (
	"aryavidyananta/Golang-Project/domain"
	"aryavidyananta/Golang-Project/dto"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type BlogApi struct {
	BlogService domain.BlogService
}

// NewBlog registers the blog routes into Fiber app
func NewBlog(app *fiber.App, blogService domain.BlogService, autMid fiber.Handler) {
	ba := BlogApi{
		BlogService: blogService,
	}

	blog := app.Group("/blogs", autMid)

	blog.Get("/", ba.Index)
	blog.Get("/:id", ba.Show)
	blog.Post("/", ba.Create)
	blog.Put("/:id", ba.Update)
	blog.Delete("/:id", ba.Delete)
}

// GET /blogs
func (ba BlogApi) Index(ctx *fiber.Ctx) error {
	result, err := ba.BlogService.Index(context.Background())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(result)
}

// GET /blogs/:id
func (ba BlogApi) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	result, err := ba.BlogService.Show(context.Background(), id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(result)
}

// POST /blogs
func (ba BlogApi) Create(ctx *fiber.Ctx) error {
	var req dto.CreateBlogRequest
	req.Title = ctx.FormValue("title")
	req.Content = ctx.FormValue("content")

	file, err := ctx.FormFile("gambar")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "gambar is required",
		})
	}
	req.Gambar = file

	result, err := ba.BlogService.Create(context.Background(), req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusCreated).JSON(result)
}

// PUT /blogs/:id
func (ba BlogApi) Update(ctx *fiber.Ctx) error {
	var req dto.UpdateBlogRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	req.Id = ctx.Params("id")

	result, err := ba.BlogService.Update(context.Background(), req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(result)
}

// DELETE /blogs/:id
func (ba BlogApi) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := ba.BlogService.Delete(context.Background(), id); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.SendStatus(http.StatusNoContent)
}
