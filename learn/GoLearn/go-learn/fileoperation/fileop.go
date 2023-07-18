package fileoperation

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadAllFiles() {
	pwd, _ := os.Getwd()
	//获取当前目录下的所有文件或目录信息
	filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		//fmt.Println(path)        //打印path信息
		fmt.Println(info.Name()) //打印文件或目录名
		return nil
	})
}

func ReadAllFiles2(basePath string) {
	files, err := ioutil.ReadDir(basePath) //读取目录下文件
	if err != nil {
		log.Print(err)
	}
	for _, file := range files {
		if file.IsDir() {
			ReadAllFiles2(basePath + "/" + file.Name())
		} else if strings.HasPrefix(file.Name(), "Go") && !strings.HasSuffix(file.Name(), ".md") {
			fmt.Println(file.Name())
		}
	}
}

func ReadAllFiles3(basePath string) {
	paths := make([]string, 0)
	//获取当前目录下的所有文件或目录信息
	filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(info.Name(), "Go") && !strings.HasSuffix(info.Name(), ".md") {
			paths = append(paths, path)
		}
		return nil
	})
	for i, path := range paths {
		if i == 0 || i == 1 {
			continue
		}
		newPath := path + ".md"
		fmt.Println(path, newPath)
		os.Rename(path, newPath)
	}
}
