package service

import (
	"aryavidyananta/Golang-Project/domain"
	"aryavidyananta/Golang-Project/dto"
	"aryavidyananta/Golang-Project/internal/config"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

func NewAuth(cnf *config.Config,
	userRepository domain.UserRepository) domain.AuthService {
	return authService{
		conf:           cnf,
		userRepository: userRepository,
	}
}

// Login implements domain.AuthService.
func (a authService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}
	if user.Id == "" {
		return dto.AuthResponse{}, errors.New("auth gagal")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthResponse{}, errors.New("auth gagal")
	}
	claims := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Duration(a.conf.Jwt.Exp) * time.Minute).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := jwtToken.SignedString([]byte(a.conf.Jwt.Key))
	if err != nil {
		return dto.AuthResponse{}, errors.New("auth gagal")
	}

	return dto.AuthResponse{
		Token: tokenStr,
	}, nil
}
