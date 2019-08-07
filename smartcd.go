package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	directories := findDirectories(".")
	for _, directory := range directories {
		if directory.IsDir() {
			fmt.Println(directory.Name())
		}
	}
}

func shorten(sourceString string) []string {
	var result []string

	strList := strings.Split(sourceString, "-")
	result = append(result, sourceString)
	result = append(result, strings.Join(strList, ""))

	shortStrList := shortSplit(sourceString, "-", 1)
	result = append(result, strings.Join(shortStrList, "-"))
	result = append(result, strings.Join(shortStrList, ""))

	return result
}

func shortSplit(sourceString string, separater string, len int) []string {
	var result []string
	strList := strings.Split(sourceString, separater)
	for _, str := range strList {
		result = append(result, str[0:len])
	}
	return result
}

func findDirectories(target string) []os.FileInfo {
	filesAndDirectories, err := ioutil.ReadDir(target)

	if err != nil {
		panic(err)
	}

	var directories []os.FileInfo
	for _, file := range filesAndDirectories {
		if file.IsDir() {
			directories = append(directories, file)
		}
	}
	return directories

}
