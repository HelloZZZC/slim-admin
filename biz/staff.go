package biz

import (
	"slim-admin/model"
	"slim-admin/repository"
)

type StaffService struct {
	repository *repository.StaffRepository
}

func (s *StaffService) GetStaff(id uint) *model.StaffModel {
	return s.repository.Get(id)
}
