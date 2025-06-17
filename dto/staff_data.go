package dto

import "mime/multipart"

type StaffData struct {
	Id      string `json:"id"`
	Nama    string `json:"nama"`
	NIP     string `json:"nip"`
	Jabatan string `json:"jabatan"`
	Gambar  string `json:"gambar"`
}

type CreateStaffRequest struct {
	Nama    string                `json:"nama" validate:"required"`
	NIP     string                `json:"nip" validate:"required"`
	Jabatan string                `json:"jabatan" validate:"required"`
	Gambar  *multipart.FileHeader `form:"gambar" validate:"required"`
}

type UpdateStaffRequest struct {
	Id      string                `json:"id" validate:"required"`
	Nama    string                `json:"nama" validate:"required"`
	NIP     string                `json:"nip" validate:"required"`
	Jabatan string                `json:"jabatan" validate:"required"`
	Gambar  *multipart.FileHeader `form:"gambar"`
}
