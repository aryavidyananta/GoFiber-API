package dto

type BookData struct {
	Id        string `json:"id"`
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
	CoverId   string `json:"cover_id"`
}

type CreateBookRequest struct {
	Judul     string `json:"judul" validate:"required"`
	Deskripsi string `json:"deskripsi" validate:"required"`
	CoverId   string `json:"cover_id"`
}

type UpdateBookRequest struct {
	Id        string `json:"-"`
	Judul     string `json:"judul" validate:"required"`
	Deskripsi string `json:"deskripsi" validate:"required"`
	CoverId   string `json:"cover_id"`
}
