package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type IndicatorsNotificationActionService struct {
	repo repository.IndicatorsNotificationRepo
}

func NewIndicatorsNotificationActionService(repo repository.IndicatorsNotificationRepo) *IndicatorsNotificationActionService {
	return &IndicatorsNotificationActionService{repo: repo}
}

func (s *IndicatorsNotificationActionService) Create(indicatorsNotification structures.IndicatorsNotification) (int, error) {
	return s.repo.Create(indicatorsNotification)
}

func (s *IndicatorsNotificationActionService) GetAll() ([]structures.IndicatorsNotification, error) {
	return s.repo.GetAll()
}

func (s *IndicatorsNotificationActionService) Get(id int) (structures.IndicatorsNotification, error) {
	return s.repo.Get(id)
}

func (s *IndicatorsNotificationActionService) GetAllByPatientID(patientID int) ([]structures.IndicatorsNotification, error) {
	return s.repo.GetAllByPatientID(patientID)
}
