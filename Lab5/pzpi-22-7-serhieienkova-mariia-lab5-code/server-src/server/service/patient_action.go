package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type PatientActionService struct {
	patientRepo         repository.PatientRepo
	patientRelativeRepo repository.PatientRelativeRepo
}

func NewPatientActionService(patientRepo repository.PatientRepo, patientRelativeRepo repository.PatientRelativeRepo) *PatientActionService {
	return &PatientActionService{patientRepo: patientRepo, patientRelativeRepo: patientRelativeRepo}
}

func (s *PatientActionService) IsRelativeAllowedToSeeRecords(relativeId int, patientId int) (bool, error) {
	patientRelatives, err := s.patientRelativeRepo.GetAllByPatientId(patientId)
	if err != nil {
		return false, err
	}

	for _, patientRelative := range patientRelatives {
		if patientRelative.RelativeID == relativeId && patientRelative.AccessToRecords {
			return true, nil
		}
	}

	return false, nil
}

func (s *PatientActionService) Create(patient structures.Patient) (int, error) {
	err := patient.Validate()
	if err != nil {
		return 0, fmt.Errorf("patient validation error: %w", err)
	}
	patient.PasswordHash = generatePasswordHash(patient.PasswordHash)
	return s.patientRepo.Create(patient)
}

func (s *PatientActionService) AttachRelative(patientRelative structures.PatientRelative) error {
	return s.patientRepo.AttachRelative(patientRelative)
}

func (s *PatientActionService) GetAllByRelativeId(relativeId int) ([]structures.Patient, error) {
	return s.patientRepo.GetAllByRelativeId(relativeId)
}

func (s *PatientActionService) GetAllByDoctorId(doctorId int) ([]structures.Patient, error) {
	return s.patientRepo.GetAllByDoctorId(doctorId)
}

func (s *PatientActionService) GetAll() ([]structures.Patient, error) {
	return s.patientRepo.GetAll()
}

func (s *PatientActionService) GetById(id int) (structures.Patient, error) {
	return s.patientRepo.GetById(id)
}

func (s *PatientActionService) Delete(id int) error {
	return s.patientRepo.Delete(id)
}

func (s *PatientActionService) GetFullInfo(patientID int) (structures.PatientFullInfo, error) {
	return structures.PatientFullInfo{}, nil
}

func (s *PatientActionService) GenerateToken(email, password string) (structures.UserToken, error) {
	patient, err := s.patientRepo.GetByCreds(email, generatePasswordHash(password))
	fmt.Println(fmt.Sprintf("received patient: %v", patient))
	if err != nil {
		return structures.UserToken{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:       patient.Id,
		UserType: structures.PatientType,
	})

	signedToken, _ := token.SignedString([]byte(signingKey))

	return structures.UserToken{
		Token:  signedToken,
		UserId: patient.Id,
	}, nil
}

func (s *PatientActionService) GenerateTokenById(patientId int) (structures.UserToken, error) {
	admin, err := s.patientRepo.GetById(patientId)
	if err != nil {
		return structures.UserToken{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:       admin.Id,
		UserType: structures.PatientType,
	})

	signedToken, _ := token.SignedString([]byte(signingKey))

	return structures.UserToken{
		Token:  signedToken,
		UserId: admin.Id,
	}, nil
}

func (s *PatientActionService) RefreshToken(refreshToken string) (structures.UserToken, structures.UserToken, error) {
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

func (s *PatientActionService) Update(id string, input structures.Patient) error {
	return nil
}
