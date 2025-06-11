package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"

	"clinic/server/repository"
	"clinic/server/structures"
)

type AdminService struct {
	repo repository.AdminRepo
}

func NewAdminService(repo repository.AdminRepo) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) Create(admin structures.Admin) (int, error) {
	err := admin.Validate()
	if err != nil {
		return 0, fmt.Errorf("admin validation error: %w", err)
	}
	admin.PasswordHash = generatePasswordHash(admin.PasswordHash)
	return s.repo.Create(admin)
}

func (s *AdminService) GenerateToken(email, password string) (structures.UserToken, error) {
	admin, err := s.repo.GetByCreds(email, generatePasswordHash(password))
	fmt.Println(fmt.Sprintf("received admin: %v", admin))
	if err != nil {
		return structures.UserToken{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:       admin.Id,
		UserType: structures.AdminType,
	})

	signedToken, _ := token.SignedString([]byte(signingKey))

	return structures.UserToken{
		Token:  signedToken,
		UserId: admin.Id,
	}, nil
}

func (s *AdminService) GetById(adminId int) (structures.Admin, error) {
	return s.repo.GetById(adminId)
}

func (s *AdminService) GenerateTokenById(adminId int) (structures.UserToken, error) {
	admin, err := s.repo.GetById(adminId)
	if err != nil {
		return structures.UserToken{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:       admin.Id,
		UserType: structures.AdminType,
	})

	signedToken, _ := token.SignedString([]byte(signingKey))

	return structures.UserToken{
		Token:  signedToken,
		UserId: admin.Id,
	}, nil
}

func (s *AdminService) RefreshToken(refreshToken string) (structures.UserToken, structures.UserToken, error) {
	// Parse the refresh token
	tokenClaims, err := ParseToken(refreshToken)
	if err != nil {
		return structures.UserToken{}, structures.UserToken{}, err
	}

	// Generate a new access token for the user
	newAccessToken, err := s.GenerateTokenById(tokenClaims.Id)
	if err != nil {
		return structures.UserToken{}, structures.UserToken{}, err
	}

	// Generate a new refresh token for the user
	newRefreshToken, err := s.GenerateTokenById(tokenClaims.Id)
	if err != nil {
		return structures.UserToken{}, structures.UserToken{}, err
	}

	return newAccessToken, newRefreshToken, nil
}
