package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"slim-admin/router"
)

func main() {
	r := gin.Default()
	router.Register(r)
	err := endless.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
