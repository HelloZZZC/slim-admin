package repository

import (
	"gorm.io/gorm"
	"slim-admin/model"
)

type StaffRepository struct {
	GormDB *gorm.DB `inject:""`
}

func (r *StaffRepository) Get(id uint) *model.Staff {
	staffModel := new(model.Staff)
	r.GormDB.First(&staffModel, id)
	return staffModel
}
