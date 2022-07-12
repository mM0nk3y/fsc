package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io"
	"os"
	"runtime"
)

func CheckFile(path string, data []byte, perm os.FileMode) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_APPEND, perm)
	defer func() { f.Close() }()
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

func API_init() {
	// Check OS is Linux or Windows
	sysType := runtime.GOOS
	if sysType == "linux" {
		path := "./config/config.yaml"
		ApiProfile := []byte("FofaEmail: \nFofaToken: \nFofaUrl: \nFofaApi: \nShodanApi: \n")
		CheckFile(path, ApiProfile, 0777)
	}
	if sysType == "windows" {
		path := "config\\config.yaml"
		ApiProfile := []byte("FofaEmail:  \r\nFofaToken:  \r\nFofaUrl:  \r\nFofaApi:  \r\nShodanApi:  \r\n")
		CheckFile(path, ApiProfile, 0777)
	}
	viper.SetConfigFile("./config/config.yaml")
	viper.AddConfigPath(".")
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed: ", in.Name)
	})
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file:%s \n", err))
	}
}
