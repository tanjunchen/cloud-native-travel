package main

import (
	"fmt"
	"path"
)

func test1() {
	baseURL := "memfs://v1.10.0-download"
	architecture := "amd64"
	component := "kube-apiserver"

	tagURL := baseURL + "/bin/linux/" + architecture + "/" + component + ".docker_tag"
	fmt.Println(tagURL)

	fmt.Println("true url", tagURL)

	url := path.Join(baseURL, "/bin/linux/", architecture, component) + ".docker_tag"
	fmt.Println("false url", url)

	url2 := path.Join(baseURL, "bin/linux", architecture, component, ".docker_tag")
	fmt.Println("false url", url2)

	component = ""
	url3 := path.Join(baseURL, "bin/linux", architecture, component, ".docker_tag")
	fmt.Println("false url", url3)
}

func main() {
}
