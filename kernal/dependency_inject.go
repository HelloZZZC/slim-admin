package kernal

import (
	"github.com/facebookgo/inject"
	"slim-admin/global"
)

// ExecuteDI 执行控制器依赖注入
func executeDI() {
	var g inject.Graph
	err := g.Provide(
		&inject.Object{Value: &global.Controllers},
		&inject.Object{Value: global.GormDB})
	if err != nil {
		panic(err)
	}
	err = g.Populate()
	if err != nil {
		panic(err)
	}
}
