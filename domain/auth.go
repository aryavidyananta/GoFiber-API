package domain

import (
	"aryavidyananta/Golang-Project/dto"
	"context"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}
