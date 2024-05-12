package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slim-admin/util"
)

type HealthController struct{}

func (c *HealthController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, new(util.Response).Success("Slim Admin"))
}
