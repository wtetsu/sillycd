package stupidcd

import (
	"os"
	"path/filepath"
	"strings"
)

func PickOutDirectory(targetPath string) string {
	var foundDirectory string
	var splittedPath = strings.Split(filepath.ToSlash(targetPath), "/")

	var targetDirectory = splittedPath[0]
	var targetName = splittedPath[1]

	var fullPath = targetDirectory + "/" + targetName
	if dirExist(fullPath) {
		return fullPath
	}

	directories := findDirectories(targetDirectory, targetName)
	for _, directory := range directories {
		names := shorten(directory)
		for _, name := range names {
			if name == targetName {
				foundDirectory = directory
				break
			}
		}
	}
	return foundDirectory
}

func dirExist(path string) bool {
	fileInfo, err := os.Stat(path)
	return err == nil && fileInfo.IsDir()
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

func findDirectories(targetDirectory string, targetFile string) []string {
	var firstLetter = targetFile[0:1]
	var lowerLetter = strings.ToLower(firstLetter)
	var upperLetter = strings.ToUpper(firstLetter)

	entries1, err := filepath.Glob(targetDirectory + "/" + strings.ToLower(firstLetter) + "*")
	if err != nil {
		panic(err)
	}
	entries2, err := filepath.Glob(targetDirectory + "/" + strings.ToUpper(firstLetter) + "*")
	if err != nil {
		panic(err)
	}

	var entries []string

	if upperLetter == firstLetter {
		entries = append(entries1, entries2...)
	} else if lowerLetter == firstLetter {
		entries = append(entries2, entries1...)
	}

	var directories []string
	for _, entry := range entries {
		fileInfo, _ := os.Stat(entry)
		if fileInfo.IsDir() {
			directories = append(directories, entry)
		}
	}

	return directories
}
