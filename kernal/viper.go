package kernal

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"path"
	"runtime"
	"slim-admin/config"
	_const "slim-admin/const"
)

var ApplicationConfig config.Application

func Viper(path ...string) *viper.Viper {
	var configFile string
	if len(path) > 0 {
		configFile = path[0]
	} else {
		configFile = readConfigFileFromCMD()
	}

	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&ApplicationConfig); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&ApplicationConfig); err != nil {
		panic(err)
	}

	return v
}

func readConfigFileFromCMD() (configFile string) {
	flag.StringVar(&configFile, "c", "", "choose config file.")
	flag.Parse()
	if configFile == "" {
		_, filename, _, _ := runtime.Caller(0)
		root := path.Dir(path.Dir(filename))
		configFile = root + "/" + _const.DefaultConfigFile
		fmt.Printf("您正在使用系统默认的配置文件,configFile的路径为%s\n", configFile)
	} else {
		fmt.Printf("您正在使用命令行的-c参数传递的值,configFile的路径为%s\n", configFile)
	}

	return
}
