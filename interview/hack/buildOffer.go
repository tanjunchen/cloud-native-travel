package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "https://leetcode-cn.com/api/problems/lcof/"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer res.Body.Close()
	fmt.Println(string(body))
}

