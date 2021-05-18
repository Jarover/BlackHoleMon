package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Jarover/BlackHoleMon/readconfig"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.New()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		fmt.Println(path)
		fmt.Println(method)
		c.HTML(http.StatusOK, "noroute.html", gin.H{
			"title": "Opanki",
		})
	})
	return r
}

func setupRouter2() *gin.Engine {
	r := gin.New()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		fmt.Println(path)
		fmt.Println(method)
		c.HTML(http.StatusOK, "noroute.html", gin.H{
			"title": "Opanki",
		})
	})
	return r
}

// Читаем флаги и окружение
func readFlag(configFlag *readconfig.Flag) {
	flag.StringVar(&configFlag.ConfigFile, "f", readconfig.GetEnv("CONFIGFILE", readconfig.GetDefaultConfigFile()), "config file")
	flag.StringVar(&configFlag.Host, "h", readconfig.GetEnv("HOST", ""), "host")
	flag.UintVar(&configFlag.Port, "p", uint(readconfig.GetEnvInt("PORT", 0)), "port")
	flag.UintVar(&configFlag.Port2, "p2", uint(readconfig.GetEnvInt("PORT2", 0)), "port2")
	flag.Parse()

}

func main() {
	var configFlag readconfig.Flag
	readFlag(&configFlag)

	fmt.Println(configFlag)

	Config, err := readconfig.ReadConfig(configFlag.ConfigFile)
	if configFlag.Port != 0 {
		Config.Port = configFlag.Port
	}

	if configFlag.Host != "" {
		Config.Host = configFlag.Host
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(Config)

	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create(readconfig.GetBaseFile() + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := setupRouter()
	go r.Run(Config.Host + ":" + strconv.FormatUint(uint64(Config.Port), 10))

	r2 := setupRouter2()
	r2.Run(Config.Host + ":" + strconv.FormatUint(uint64(Config.Port2), 10))

}
