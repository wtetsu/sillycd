package stupidcd

import (
	"io/ioutil"
	"os"
	"strings"
)

func PickOutDirectory(directory string, target string) string {
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

func Shorten(sourceString string) []string {
	return shorten(sourceString)
}

func shorten(sourceString string) []string {
	var result []string

	{
		strList := strings.Split(sourceString, "-")
		result = append(result, sourceString)
		result = append(result, strings.Join(strList, ""))

		shortStr1List := shortSplit(sourceString, "-", 1)
		result = append(result, strings.Join(shortStr1List, "-"))
		result = append(result, strings.Join(shortStr1List, ""))

		shortStr2List := shortSplit(sourceString, "-", 2)
		result = append(result, strings.Join(shortStr2List, "-"))
		result = append(result, strings.Join(shortStr2List, ""))
	}

	{
		strList := strings.Split(sourceString, "_")
		result = append(result, sourceString)
		result = append(result, strings.Join(strList, ""))

		shortStr1List := shortSplit(sourceString, "_", 1)
		result = append(result, strings.Join(shortStr1List, "_"))
		result = append(result, strings.Join(shortStr1List, ""))

		shortStr2List := shortSplit(sourceString, "_", 2)
		result = append(result, strings.Join(shortStr2List, "_"))
		result = append(result, strings.Join(shortStr2List, ""))
	}

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