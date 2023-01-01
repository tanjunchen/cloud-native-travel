package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

var (
	port      string
	version   string
	headerNum int
)

func init() {
	port = getEnvDefault("port", "8000")
	version = getEnvDefault("SERVICE_VERSION", "v1")
	headerNum = *flag.Int("HEADER_NUM", 90, "")
}

func getEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}

func main() {
	flag.Parse()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, version")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "health")
	})

	e.GET("/hello", hello)

	e.GET("/headers", headers)
	e.Logger.Fatal(e.Start(":" + port))
}

func value() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	res := fmt.Sprintf("Hello version: %s, instance: %s\n", version, name)
	fmt.Printf(res)
	return res
}

func hello(c echo.Context) error {
	res := value()
	return c.String(http.StatusOK, res)
}

func headers(c echo.Context) error {
	for i := 0; i < headerNum; i++ {
		c.Response().Header().Add("X-", fmt.Sprintf("%v", i))
	}
	fmt.Printf("============>headers")
	res := value()
	return c.String(http.StatusOK, res)
}
