package repository

import "slim-admin/model"

type StaffRepository interface {
	get(id uint) *model.StaffModel
}

type StaffRepositoryImpl struct{}

func (impl StaffRepositoryImpl) get(id uint) *model.StaffModel {
	return nil
}
