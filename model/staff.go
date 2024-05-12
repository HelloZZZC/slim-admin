package model

import (
	"gorm.io/gorm"
	"slim-admin/enum"
)

type StaffModel struct {
	gorm.Model
	Nickname     *string `json:",omitempty"`
	RealName     *string `json:",omitempty"`
	Email        string
	Phone        string
	Avatar       *string `json:",omitempty"`
	Gender       enum.Gender
	Status       enum.StaffStatus
	RoleID       uint
	DepartmentID uint
	CreatorID    uint
	OperatorID   uint
}
