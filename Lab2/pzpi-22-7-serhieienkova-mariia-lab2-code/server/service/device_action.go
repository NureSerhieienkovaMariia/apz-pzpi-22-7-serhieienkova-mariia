package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type DeviceActionService struct {
	repo repository.DeviceRepo
}

func NewDeviceActionService(repo repository.DeviceRepo) *DeviceActionService {
	return &DeviceActionService{repo: repo}
}

func (s *DeviceActionService) Create(device structures.Device) (int, error) {
	device.Password = generatePasswordHash(device.Password)
	return s.repo.Create(device)
}

func (s *DeviceActionService) GetAll() ([]structures.Device, error) {
	return s.repo.GetAll()
}

func (s *DeviceActionService) Get(id int) (structures.Device, error) {
	return s.repo.Get(id)
}

func (s *DeviceActionService) Update(id int, input structures.Device) error {
	if input.Password != "" {
		input.Password = generatePasswordHash(input.Password)
	}
	return s.repo.Update(id, input)
}

func (s *DeviceActionService) Delete(id int) error {
	return s.repo.Delete(id)
}
