package router

import "github.com/gin-gonic/gin"

type RouteRegister interface {
	register()
}

func Register(r *gin.Engine) *gin.Engine {
	r.Group("")
	{
		HealthRouter{r}.register()
	}

	return r
}
