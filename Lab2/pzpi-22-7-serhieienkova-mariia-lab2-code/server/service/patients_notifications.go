package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type PatientNotificationService struct {
	repo repository.PatientNotificationRepo
}

func NewPatientNotificationService(repo repository.PatientNotificationRepo) *PatientNotificationService {
	return &PatientNotificationService{repo: repo}
}

func (s *PatientNotificationService) Create(notification structures.PatientNotification) (int, error) {
	return s.repo.Create(notification)
}

func (s *PatientNotificationService) GetAll() ([]structures.PatientNotification, error) {
	return s.repo.GetAll()
}

func (s *PatientNotificationService) Get(id int) (structures.PatientNotification, error) {
	return s.repo.Get(id)
}

func (s *PatientNotificationService) Update(id int, input structures.PatientNotification) error {
	return s.repo.Update(id, input)
}

func (s *PatientNotificationService) Delete(id int) error {
	return s.repo.Delete(id)
}
