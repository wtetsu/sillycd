package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Wrong argument")
		os.Exit(1)
	}

	foundDirectory := pickOutDirectory(".", os.Args[1])

	if foundDirectory == "" {
		os.Exit(1)
	}
	fmt.Println(foundDirectory)
}

func pickOutDirectory(directory string, target string) string {
	var foundDirectory string

	directories := findDirectories(directory)
	for _, directory := range directories {
		names := shorten(directory.Name())
		for _, name := range names {
			if name == target {
				foundDirectory = directory.Name()
				break
			}
		}
	}
	return foundDirectory
}

func shorten(sourceString string) []string {
	var result []string

	strList := strings.Split(sourceString, "-")
	result = append(result, sourceString)
	result = append(result, strings.Join(strList, ""))

	shortStr1List := shortSplit(sourceString, "-", 1)
	result = append(result, strings.Join(shortStr1List, "-"))
	result = append(result, strings.Join(shortStr1List, ""))

	shortStr2List := shortSplit(sourceString, "-", 2)
	result = append(result, strings.Join(shortStr2List, "-"))
	result = append(result, strings.Join(shortStr2List, ""))

	return result
}

func shortSplit(sourceString string, separater string, length int) []string {
	var result []string
	strList := strings.Split(sourceString, separater)
	for _, str := range strList {
		if len(str) >= length {
			result = append(result, str[0:length])
		} else {
			result = append(result, str)
		}
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
