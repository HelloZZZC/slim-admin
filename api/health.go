package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slim-admin/util"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, new(util.Response).Success("Slim Admin"))
}
