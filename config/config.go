package config

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tedforv/goutil/filesystem"
	"os"
)

var settingFilePath string

// ViperConfig is global config
var ViperConfig *viper.Viper

func init() {
	setFlag()
	if len(settingFilePath) != 0 {
		configFilePath = settingFilePath
	}
	existed, err := filesystem.IsPathExisted(configFilePath)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	if !existed {
		logrus.Error(fmt.Sprintf("file (%s) is not existed", configFilePath))
		panic(nil)
	}
	file, err := os.Open(configFilePath)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	ViperConfig = viper.New()
	ViperConfig.SetConfigType("yaml")
	err = ViperConfig.ReadConfig(file)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
}

func setFlag() {
	flag.StringVar(&settingFilePath, "c", "", "set configuration file")

	flag.Parse()
}
