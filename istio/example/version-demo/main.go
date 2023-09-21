package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

var (
	portVal    string
	versionVal string
)

func init() {
	portVal = getEnvDefault("PORT", "9000")
	versionVal = getEnvDefault("VERSION", "dev")
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
	e.GET("/version", version)

	e.Logger.Fatal(e.Start(":" + portVal))
}

func version(c echo.Context) error {
	fmt.Println("version:", versionVal)
	return c.String(http.StatusOK, versionVal)
}
