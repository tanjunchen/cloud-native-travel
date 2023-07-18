package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
)

var (
	port string
)

func init() {
	port = getEnvDefault("port", "10001")
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
		return c.String(http.StatusOK, "Hello, Provider Demo")
	})
	e.GET("/echo/:msg", getMsg)
	e.Logger.Fatal(e.Start(":" + port))
}

func getMsg(c echo.Context) error {
	msg := c.Param("msg")
	if msg == "" {
		return c.String(http.StatusBadRequest, "no parameter")
	}
	ip, err := getIP(c.Request())
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	res := "echo() -> ip [ " + ip + " ] param [ " + msg + " ]"
	fmt.Println(res)
	return c.String(http.StatusOK, res)
}

// getIP returns request real ip.
func getIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}
