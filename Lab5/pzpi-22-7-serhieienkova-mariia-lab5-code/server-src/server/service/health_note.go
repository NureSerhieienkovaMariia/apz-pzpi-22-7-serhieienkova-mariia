package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
	"time"
)

type HealthNoteActionService struct {
	repo repository.HealthNoteRepo
}

func (s *HealthNoteActionService) GetAllByPatientId(patientId int) ([]structures.HealthNote, error) {
	healthNotes, err := s.repo.GetAllByPatientId(patientId)
	if err != nil {
		return nil, err
	}
	return healthNotes, nil
}

func NewHealthNoteActionService(repo repository.HealthNoteRepo) *HealthNoteActionService {
	return &HealthNoteActionService{repo: repo}
}

func (s *HealthNoteActionService) Create(healthNote structures.HealthNote) (int, error) {
	healthNote.Timestamp = time.Now()
	return s.repo.Create(healthNote)
}

func (s *HealthNoteActionService) GetAll() ([]structures.HealthNote, error) {
	return s.repo.GetAll()
}

func (s *HealthNoteActionService) Get(id int) (structures.HealthNote, error) {
	return s.repo.Get(id)
}

func (s *HealthNoteActionService) Update(id int, input structures.HealthNote) error {
	return s.repo.Update(id, input)
}

func (s *HealthNoteActionService) Delete(id int) error {
	return s.repo.Delete(id)
}
