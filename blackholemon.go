package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Jarover/BlackHoleMon/readconfig"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.New()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

// Читаем флаги и окружение
func readFlag(configFlag *readconfig.Flag) {
	flag.StringVar(&configFlag.ConfigFile, "f", readconfig.GetEnv("CONFIGFILE", ""), "config file")
	flag.UintVar(&configFlag.Port, "p", uint(readconfig.GetEnvInt("PORT", 0)), "port")
	flag.Parse()

}

func main() {

	var configfile string

	var configFlag readconfig.Flag
	readFlag(&configFlag)

	fmt.Println(configFlag)

	if configFlag.ConfigFile == "" {
		// берем конфиг по умолчанию

		configfile = readconfig.GetDefaultConfigFile()

	} else {
		configfile = configFlag.ConfigFile
	}
	fmt.Println(configfile)

	Config, err := readconfig.ReadConfig(configfile)
	if configFlag.Port != 0 {
		Config.Port = configFlag.Port
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(Config)

	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := setupRouter()
	r.Run()
}
