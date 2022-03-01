package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

var (
	port    string
	version string
)

func init() {
	port = getEnvDefault("port", "5000")
	version = getEnvDefault("SERVICE_VERSION", "v1")
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
	e.Logger.Fatal(e.Start(":" + port))
}

func hello(c echo.Context) error {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	res := fmt.Sprintf("Hello version: %s, instance: %s\n", version, name)
	fmt.Printf(res)
	return c.String(http.StatusOK, res)
}
