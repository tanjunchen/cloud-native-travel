package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
)

var (
	consumerDemoURL string
	port            string
)

func getEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}

func init() {
	port = getEnvDefault("port", "9999")
	//consumerDemoURL = getEnvDefault("consumerDemoURL", "http://localhost:10001/echo-rest")
	consumerDemoURL = getEnvDefault("consumerDemoURL", "http://provider-demo/echo")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Consumer Demo")
	})
	e.GET("/echo-rest/:msg", getMsg)
	e.Logger.Fatal(e.Start(":" + port))
}

func getMsg(c echo.Context) error {
	msg := c.Param("msg")
	if msg == "" {
		return c.String(http.StatusBadRequest, "no parameter")
	}
	headers := getForwardHeaders(c.Request())
	res, err := getResponseStr(consumerDemoURL+"/"+msg, headers)
	if err != nil {
		panic(err)
	}
	return c.String(http.StatusOK, res)
}

var client = &http.Client{Timeout: 10 * time.Second}

func getResponseStr(url string, headers map[string]string) (string, error) {
	request, err := http.NewRequest("GET", url, nil)
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	if err != nil {
		return "", err
	}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	msgStr, err := parseResponseString(response)
	return msgStr, nil
}

func parseResponseString(response *http.Response) (string, error) {
	body, err := ioutil.ReadAll(response.Body)
	return string(body), err
}

func getForwardHeaders(r *http.Request) map[string]string {
	headers := make(map[string]string)
	forwardHeaders := []string{
		"user",
		"x-request-id",
		"x-b3-traceid",
		"x-b3-spanid",
		"x-b3-parentspanid",
		"x-b3-sampled",
		"x-b3-flags",
		"x-ot-span-context",
	}

	for _, h := range forwardHeaders {
		if v := r.Header.Get(h); v != "" {
			headers[h] = v
		}
	}

	return headers
}
