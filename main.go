package main

import (
	"fmt"
	"slim-admin/kernal"
)

func main() {
	fmt.Print("               _____ _ _                       _           _\n")
	fmt.Print("              / ____| (_)             /\\      | |         (_)\n")
	fmt.Print("             | (___ | |_ _ __ ___    /  \\   __| |_ __ ___  _ _ __\n")
	fmt.Print("              \\___ \\| | | '_ ` _ \\  / /\\ \\ / _` | '_ ` _ \\| | '_ \\\n")
	fmt.Print("              ____) | | | | | | | |/ ____ \\ (_| | | | | | | | | | |\n")
	fmt.Print("             |_____/|_|_|_| |_| |_/_/    \\_\\__,_|_| |_| |_|_|_| |_|\n")

	// 服务启动
	kernal.Bootstrap()
}
