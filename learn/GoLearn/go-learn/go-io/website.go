package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Info struct {
	fileLocation string
	fileName     string
	enExist      map[string]string
	zhExist      map[string]string
}

const EN_LOCALTION = "D:\\opensource\\website\\content\\en"
const ZH_LOCALTION = "D:\\opensource\\website\\content\\zh"

func getAllFiles(path string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	pathSep := string(os.PathSeparator)

	for _, file := range dir {
		if file.IsDir() {
			dirs = append(dirs, path+pathSep+file.Name())
			getAllFiles(path + pathSep + file.Name())
		} else {
			if !strings.HasSuffix(file.Name(), ".md") {
				files = append(files, path+pathSep+file.Name())
			}
		}
	}

	for _, table := range dirs {
		temp, _ := getAllFiles(table)
		for _, tt := range temp {
			files = append(files, tt)
		}
	}
	return files, nil
}

func main() {
	enFiles, err := getAllFiles(EN_LOCALTION)
	if err != nil {
		fmt.Println(err)
	}
	zhFiles, err := getAllFiles(ZH_LOCALTION)
	if err != nil {
		fmt.Println(err)
	}

	for i, _ := range zhFiles {
		zhFiles[i] = strings.ReplaceAll(zhFiles[i], "\\content\\zh", "\\content\\en")
	}

	enMap := make(map[string]string, 550)
	for _, f := range enFiles {
		enMap[f] = f
	}
	for i := range zhFiles {
		diffString(zhFiles[i], enMap[zhFiles[i]])
	}
}

func diffString(a, b string) {
	a = strings.ReplaceAll(a, "\\content\\en\\", "\\content\\zh\\")
	contentA, _ := ioutil.ReadFile(a)
	contentB, _ := ioutil.ReadFile(b)
	if !strings.EqualFold(string(contentA), string(contentB)) && strings.Contains(a, ".yaml") {
		fmt.Println(strings.ReplaceAll(a, "D:\\opensource\\website\\", ""), " | ", strings.ReplaceAll(b, "D:\\opensource\\website\\", ""))
	}
}
