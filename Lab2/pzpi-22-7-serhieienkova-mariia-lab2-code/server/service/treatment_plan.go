package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type TreatmentPlanActionService struct {
	repo repository.TreatmentPlanRepo
}

func (s *TreatmentPlanActionService) GetByPatientId(patientId int) ([]structures.TreatmentPlan, error) {
	return s.repo.GetByPatientId(patientId)
}

func (s *TreatmentPlanActionService) GetAllByDoctorId(doctorId int) ([]structures.TreatmentPlan, error) {
	return s.repo.GetAllByDoctorId(doctorId)
}

func NewTreatmentPlanActionService(repo repository.TreatmentPlanRepo) *TreatmentPlanActionService {
	return &TreatmentPlanActionService{repo: repo}
}

func (s *TreatmentPlanActionService) Create(treatmentPlan structures.TreatmentPlan) (int, error) {
	return s.repo.Create(treatmentPlan)
}

func (s *TreatmentPlanActionService) GetAll() ([]structures.TreatmentPlan, error) {
	return s.repo.GetAll()
}

func (s *TreatmentPlanActionService) Get(id int) (structures.TreatmentPlan, error) {
	return s.repo.Get(id)
}

func (s *TreatmentPlanActionService) Update(id int, input structures.TreatmentPlan) error {
	return s.repo.Update(id, input)
}

func (s *TreatmentPlanActionService) Delete(id int) error {
	return s.repo.Delete(id)
}
