package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type VisitActionService struct {
	repo repository.VisitRepo
}

func (s *VisitActionService) GetAllByTreatmentPlanId(treatmentPlanId int) ([]structures.Visit, error) {
	return s.repo.GetAllByTreatmentPlanId(treatmentPlanId)
}

func (s *VisitActionService) GetAllTodaysVisits() ([]structures.Visit, error) {
	return s.repo.GetVisitsForNextDays(1)
}

func (s *VisitActionService) GetAllWeeksVisits() ([]structures.Visit, error) {
	return s.repo.GetVisitsForNextDays(7)
}

func NewVisitActionService(repo repository.VisitRepo) *VisitActionService {
	return &VisitActionService{repo: repo}
}

func (s *VisitActionService) Create(visit structures.Visit) (int, error) {
	return s.repo.Create(visit)
}

func (s *VisitActionService) GetAll() ([]structures.Visit, error) {
	return s.repo.GetAll()
}

func (s *VisitActionService) Get(id int) (structures.Visit, error) {
	return s.repo.Get(id)
}

func (s *VisitActionService) Update(id int, input structures.Visit) error {
	return s.repo.Update(id, input)
}

func (s *VisitActionService) Delete(id int) error {
	return s.repo.Delete(id)
}
