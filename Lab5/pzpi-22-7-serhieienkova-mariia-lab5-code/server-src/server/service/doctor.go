package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type DoctorService struct {
	repo repository.DoctorRepo
}

func NewDoctorService(repo repository.DoctorRepo) *DoctorService {
	return &DoctorService{repo: repo}
}

func (s *DoctorService) Create(doctor structures.Doctor) (int, error) {
	err := doctor.Validate()
	if err != nil {
		return 0, fmt.Errorf("doctor validation error: %w", err)
	}
	doctor.PasswordHash = generatePasswordHash(doctor.PasswordHash)
	return s.repo.Create(doctor)
}

func (s *DoctorService) GenerateToken(email, password string) (structures.UserToken, error) {
	doctor, err := s.repo.GetByCreds(email, generatePasswordHash(password))
	if err != nil {
		return structures.UserToken{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:       doctor.Id,
		UserType: structures.DoctorType,
	})

	signedToken, _ := token.SignedString([]byte(signingKey))

	return structures.UserToken{
		Token:  signedToken,
		UserId: doctor.Id,
	}, nil
}

func (s *DoctorService) GetById(doctorId int) (structures.Doctor, error) {
	return s.repo.GetById(doctorId)
}

func (s *DoctorService) GenerateTokenById(doctorId int) (structures.UserToken, error) {
	doctor, err := s.repo.GetById(doctorId)
	if err != nil {
		return structures.UserToken{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:       doctor.Id,
		UserType: structures.DoctorType,
	})

	signedToken, _ := token.SignedString([]byte(signingKey))

	return structures.UserToken{
		Token:  signedToken,
		UserId: doctor.Id,
	}, nil
}

func (s *DoctorService) RefreshToken(refreshToken string) (structures.UserToken, structures.UserToken, error) {
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
