package dto

import (
	"mime/multipart"
)

type BlogData struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Gambar  string `json:"gambar"`
}

type CreateBlogRequest struct {
	Title   string                `form:"title" validate:"required"`
	Content string                `form:"content" validate:"required"`
	Gambar  *multipart.FileHeader `form:"gambar" validate:"required"`
}

type UpdateBlogRequest struct {
	Id      string                `json:"id"`
	Title   string                `json:"title" validate:"required"`
	Content string                `json:"content" validate:"required"`
	Gambar  *multipart.FileHeader `form:"gambar"`
}
