package biz

import "slim-admin/model"

type StaffService interface {
	getStaff(id uint) *model.StaffModel
}

type StaffServiceImpl struct{}

func (impl StaffServiceImpl) getStaff(id uint) *model.StaffModel {
	return nil
}
