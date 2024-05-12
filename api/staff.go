package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slim-admin/biz"
	"slim-admin/util"
)

type StaffController struct {
	service *biz.StaffService
}

func (c *StaffController) getStaff(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, new(util.Response).Success(c.service.GetStaff(1)))
}
