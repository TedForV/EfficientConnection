package main

import (
	"fmt"
	"github.com/TedForV/EfficientConnection/config"
	"github.com/sirupsen/logrus"
	"github.com/tedforv/gin-util/base"
	"strconv"
)

func main() {

	r, err := base.CreateDefaultGin(
		config.ViperConfig.GetString("mode") == "product",
		config.ViperConfig.GetBool("cors"),
		config.ViperConfig.GetString("logFolderPath"),
		nil,
		[]string{"Authorization"},
		nil)

	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	port := config.ViperConfig.GetInt("port")

	r.Run(fmt.Sprintf(":%s", strconv.Itoa(port)))
}
