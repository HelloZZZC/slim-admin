package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"slim-admin/router"
)

func main() {
	fmt.Print("               _____ _ _                       _           _\n")
	fmt.Print("              / ____| (_)             /\\      | |         (_)\n")
	fmt.Print("             | (___ | |_ _ __ ___    /  \\   __| |_ __ ___  _ _ __\n")
	fmt.Print("              \\___ \\| | | '_ ` _ \\  / /\\ \\ / _` | '_ ` _ \\| | '_ \\\n")
	fmt.Print("              ____) | | | | | | | |/ ____ \\ (_| | | | | | | | | | |\n")
	fmt.Print("             |_____/|_|_|_| |_| |_/_/    \\_\\__,_|_| |_| |_|_|_| |_|\n")

	r := gin.Default()
	router.Register(r)
	err := endless.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
