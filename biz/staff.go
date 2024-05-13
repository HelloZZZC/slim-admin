package biz

import (
	"slim-admin/model"
	"slim-admin/repository"
)

type StaffService struct {
	Repository *repository.StaffRepository `inject:""`
}

// GetStaff 根据职工ID获取职工详情
func (svc *StaffService) GetStaff(id uint) *model.Staff {
	return svc.Repository.Get(id)
}
