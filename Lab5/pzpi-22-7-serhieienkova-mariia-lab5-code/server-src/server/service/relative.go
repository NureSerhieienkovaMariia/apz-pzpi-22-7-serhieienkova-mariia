package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type RelativeService struct {
	repo repository.RelativeRepo
}

func NewRelativeService(repo repository.RelativeRepo) *RelativeService {
	return &RelativeService{repo: repo}
}

func (s *RelativeService) Create(relative structures.Relative) (int, error) {
	err := relative.Validate()
	if err != nil {
		return 0, fmt.Errorf("relative validation error: %w", err)
	}
	relative.PasswordHash = generatePasswordHash(relative.PasswordHash)
	return s.repo.Create(relative)
}

func (s *RelativeService) GenerateToken(email, password string) (structures.UserToken, error) {
	relative, err := s.repo.GetByCreds(email, generatePasswordHash(password))
	if err != nil {
		return structures.UserToken{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:       relative.Id,
		UserType: structures.RelativeType,
	})

	signedToken, _ := token.SignedString([]byte(signingKey))

	return structures.UserToken{
		Token:  signedToken,
		UserId: relative.Id,
	}, nil
}

func (s *RelativeService) GetById(relativeId int) (structures.Relative, error) {
	return s.repo.GetById(relativeId)
}

func (s *RelativeService) GenerateTokenById(relativeId int) (structures.UserToken, error) {
	relative, err := s.repo.GetById(relativeId)
	if err != nil {
		return structures.UserToken{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:       relative.Id,
		UserType: structures.RelativeType,
	})

	signedToken, _ := token.SignedString([]byte(signingKey))

	return structures.UserToken{
		Token:  signedToken,
		UserId: relative.Id,
	}, nil
}

func (s *RelativeService) RefreshToken(refreshToken string) (structures.UserToken, structures.UserToken, error) {
	tokenClaims, err := ParseToken(refreshToken)
	if err != nil {
		return structures.UserToken{}, structures.UserToken{}, err
	}

	newAccessToken, err := s.GenerateTokenById(tokenClaims.Id)
	if err != nil {
		return structures.UserToken{}, structures.UserToken{}, err
	}

	newRefreshToken, err := s.GenerateTokenById(tokenClaims.Id)
	if err != nil {
		return structures.UserToken{}, structures.UserToken{}, err
	}

	return newAccessToken, newRefreshToken, nil
}

func (s *RelativeService) GetAllByPatientId(patientId int) ([]structures.Relative, error) {
	relatives, err := s.repo.GetAllByPatientId(patientId)
	if err != nil {
		return nil, fmt.Errorf("failed to get relatives by patient id: %w", err)
	}
	return relatives, nil
}

func (s *RelativeService) Update(id string, input structures.Relative) error {
	return nil
}

func (s *RelativeService) Delete(id string) error {
	return nil
}
