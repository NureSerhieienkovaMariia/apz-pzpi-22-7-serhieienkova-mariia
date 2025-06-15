package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type PrescriptionActionService struct {
	repo repository.PrescriptionRepo
}

func (s *PrescriptionActionService) GetAllByTreatmentPlanId(treatmentPlanId int) ([]structures.Prescription, error) {
	return s.repo.GetAllByTreatmentPlanId(treatmentPlanId)
}

func NewPrescriptionActionService(repo repository.PrescriptionRepo) *PrescriptionActionService {
	return &PrescriptionActionService{repo: repo}
}

func (s *PrescriptionActionService) Create(prescription structures.Prescription) (int, error) {
	return s.repo.Create(prescription)
}

func (s *PrescriptionActionService) GetAll() ([]structures.Prescription, error) {
	return s.repo.GetAll()
}

func (s *PrescriptionActionService) Get(id int) (structures.Prescription, error) {
	return s.repo.Get(id)
}

func (s *PrescriptionActionService) Update(id int, input structures.Prescription) error {
	return s.repo.Update(id, input)
}

func (s *PrescriptionActionService) Delete(id int) error {
	return s.repo.Delete(id)
}
