package dto

import "mime/multipart"

type EkskulData struct {
	id     string `json:"id"`
	Name   string `json:"name"`
	Gambar string `json:"gambar"`
}

type CreateEkskulRequest struct {
	Name   string                `json:"name" validate:"required"`
	Gambar *multipart.FileHeader `form:"gambar" validate:"required"`
}

type UpdateEkskulRequest struct {
	Id     string                `json:"id" validate:"required"`
	Name   string                `json:"name" validate:"required"`
	Gambar *multipart.FileHeader `form:"gambar"`
}
