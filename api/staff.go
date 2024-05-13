package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slim-admin/biz"
	"slim-admin/util"
	"strconv"
)

type StaffController struct {
	Service *biz.StaffService `inject:""`
}

// GetStaff 获取职员详情
func (c *StaffController) GetStaff(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	ctx.JSON(http.StatusOK, new(util.Response).Success(c.Service.GetStaff(uint(id))))
}
