package kernal

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"path"
	"runtime"
	_const "slim-admin/const"
	"slim-admin/global"
)

// LoadConfigYaml 载入yaml配置文件
func loadConfigYaml(path ...string) {
	var configFile string
	if len(path) > 0 {
		configFile = path[0]
	} else {
		configFile = readConfigFileFromCMD()
	}

	global.Viper = viper.New()
	global.Viper.SetConfigFile(configFile)
	global.Viper.SetConfigType("yaml")

	if err := global.Viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	global.Viper.WatchConfig()

	global.Viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := global.Viper.Unmarshal(&global.ApplicationConfig); err != nil {
			fmt.Println(err)
		}
	})

	if err := global.Viper.Unmarshal(&global.ApplicationConfig); err != nil {
		panic(err)
	}
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
