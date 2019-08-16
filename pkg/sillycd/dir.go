package sillycd

import (
	"os"
	"path/filepath"
	"strings"
)

// PickOutDirectory picks out one directory.
func PickOutDirectory(targetPath string) string {
	return PickOutDirectoryWithFunction(targetPath, findDirectories)
}

// PickOutDirectoryWithFunction picks out one directory.
func PickOutDirectoryWithFunction(targetPath string, doGetDirectories func(string, string) []string) string {
	var splittedPath = strings.Split(filepath.ToSlash(targetPath), "/")

	var firstDirectory string
	if filepath.IsAbs(targetPath) {
		firstDirectory = filepath.Join("/" + splittedPath[0])
	} else {
		absolutePath, err := filepath.Abs(splittedPath[0])
		if err != nil {
			panic(err)
		}
		firstDirectory = absolutePath
	}

	resultDirectory := traverseDirectoriesByTheEnd(firstDirectory, splittedPath[1:], doGetDirectories)
	return resultDirectory
}

func traverseDirectoriesByTheEnd(firstTargetDirectory string, wishTargets []string, doGetDirectories func(string, string) []string) string {
	var doFindDirectory func(string, int) string

	doFindDirectory = func(targetDirectory string, currentIndex int) string {
		if currentIndex >= len(wishTargets) {
			return targetDirectory
		}
		var candidate string
		directories := doGetDirectories(targetDirectory, wishTargets[currentIndex])
		for _, dir := range directories {
			candidate = doFindDirectory(filepath.Join(targetDirectory, dir), currentIndex+1)
			if candidate != "" {
				break
			}
		}
		return candidate
	}

	return doFindDirectory(firstTargetDirectory, 0)
}

func dirExist(path string) bool {
	fileInfo, err := os.Stat(path)
	return err == nil && fileInfo.IsDir()
}

// Shorten returns shortened names.
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

func findDirectories(targetDirectory string, targetName string) []string {
	patterns := generateFindDictionaryPatterns(targetDirectory, targetName)
	var directories []string
	for _, pattern := range patterns {
		foundDirectories := findDirectoriesByPattern(pattern)
		matchedDirectories := filterDirectoriesByName(foundDirectories, targetName)
		directories = append(directories, matchedDirectories...)
	}
	return directories
}

func filterDirectoriesByName(directories []string, targetName string) []string {
	var filteredDirectories []string
	for _, directory := range directories {
		shortNames := shorten(directory)
		for _, name := range shortNames {
			if name == targetName {
				foundDirectory := directory
				filteredDirectories = append(filteredDirectories, foundDirectory)
			}
		}
	}
	return filteredDirectories
}

func generateFindDictionaryPatterns(targetDirectory string, targetFile string) []string {
	var patterns []string
	if len(targetFile) == 0 {
		return patterns
	}
	var firstLetter = targetFile[0:1]
	var lowerLetter = strings.ToLower(firstLetter)
	var upperLetter = strings.ToUpper(firstLetter)

	if lowerLetter == firstLetter {
		patterns = []string{
			filepath.Join(targetDirectory, lowerLetter) + "*",
			filepath.Join(targetDirectory, upperLetter) + "*",
		}
	} else if lowerLetter == firstLetter {
		patterns = []string{
			filepath.Join(targetDirectory, upperLetter) + "*",
			filepath.Join(targetDirectory, lowerLetter) + "*",
		}
	}
	return patterns

}

func findDirectoriesByPattern(pattern string) []string {
	var directories []string
	entries, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		fileInfo, _ := os.Stat(entry)
		if fileInfo.IsDir() {
			relativePath := filepath.Base(entry)
			directories = append(directories, relativePath)
		}
	}
	return directories
}
