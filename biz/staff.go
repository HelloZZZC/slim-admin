package biz

import (
	"encoding/json"
	"fmt"
	_const "slim-admin/const"
	"slim-admin/model"
	"slim-admin/repository"
	"slim-admin/util"
	"time"
)

type StaffService struct {
	Repository     *repository.StaffRepository `inject:""`
	RedisCacheUtil *util.RedisCacheUtil        `inject:""`
}

// GetStaff 根据职工ID获取职工详情
func (svc *StaffService) GetStaff(id uint) *model.Staff {
	return svc.Repository.Get(id)
}

// GetStaffFromCache 根据ID从缓存中获取职工详情
func (svc *StaffService) GetStaffFromCache(id uint) *model.Staff {
	var staff *model.Staff
	cacheKey := fmt.Sprintf(_const.StaffDetail, id)
	result := svc.RedisCacheUtil.Get(cacheKey)
	if result != "" {
		err := json.Unmarshal([]byte(result), &staff)
		if err != nil {
			panic(err)
		}
		return staff
	}

	staff = svc.GetStaff(id)
	if staff != nil {
		staffJson, err := json.Marshal(&staff)
		if err != nil {
			panic(err)
		}
		svc.RedisCacheUtil.Set(cacheKey, string(staffJson), 60*time.Minute)
	}

	return staff
}
