/*
 * Copyright 2019 wtetsu
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sillycd

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// PickOutDirectory picks out one directory.
func PickOutDirectory(targetPath string) string {
	return PickOutDirectoryWithFunction(targetPath, findDirectories)
}

// PickOutDirectoryWithFunction picks out one directory.
func PickOutDirectoryWithFunction(orgTargetDirectory string, doGetDirectories func(string, string) []string) string {
	if len(orgTargetDirectory) == 0 {
		return ""
	}

	var targetPath string

	firstLetter := orgTargetDirectory[0]
	if firstLetter == '/' || firstLetter == '\\' {
		targetPath = getVolumeName() + filepath.ToSlash(orgTargetDirectory)
	} else {
		targetPath = filepath.ToSlash(orgTargetDirectory)
	}

	var splittedPath = strings.Split(targetPath, "/")
	var firstDirectory string

	if IsAbs(targetPath) {
		firstDirectory = getFirstDirectory(targetPath)
	} else {
		absolutePath, err := filepath.Abs(splittedPath[0])
		if err != nil {
			return ""
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

func findDirectories(targetDirectory string, targetName string) []string {
	patterns := generateFindDictionaryPatterns(targetDirectory, targetName)
	var directories []string
	for _, pattern := range patterns {
		foundDirectories := findDirectoriesByPattern(pattern)
		directories = append(directories, foundDirectories...)
	}
	sortedDirectories := sortDirectoriesByScore(directories, targetName)
	return sortedDirectories
}

type directory struct {
	Name  string
	Score int
}

func sortDirectoriesByScore(directoryNames []string, targetName string) []string {
	var directories []directory
	for _, dirName := range directoryNames {
		score := computeDirectoryScore(dirName, targetName)
		if score > 0 {
			directories = append(directories, directory{dirName, score})
		}
	}

	sort.SliceStable(directories, func(i, j int) bool {
		return directories[i].Score > directories[j].Score
	})

	var sortedDirectoryNames []string
	for _, dirName := range directories {
		sortedDirectoryNames = append(sortedDirectoryNames, dirName.Name)
	}
	return sortedDirectoryNames
}

// "foo-bar-baz"
//   "f":   10
//   "fo":  20
//   "fb":  20
//   "fbb": 30
//   "foo-bar-baz": 999999999
func computeDirectoryScore(directoryName string, specifiedName string) int {
	if directoryName == specifiedName {
		return 999999999
	}
	var score int

	names := splitDirectoryName(directoryName)
	var restSpecifiedName = specifiedName

	for i := 0; ; i++ {
		if restSpecifiedName == "" {
			break
		}
		if i >= len(names) {
			score = -1
			break
		}
		name := names[i]

		len, rate := matchedPrefixLength(name, restSpecifiedName)
		if len == 0 {
			score = -1
			break
		}
		score += rate * 5
		restSpecifiedName = restSpecifiedName[len:]
	}

	return score
}

func matchedPrefixLength(orgString string, orgPrefix string) (int, int) {
	str := strings.ToLower(orgString)
	prefix := strings.ToLower(orgPrefix)

	var length int
	var rate int
	for i := 0; ; i++ {
		if i >= len(str) || i >= len(prefix) {
			break
		}
		if orgString[i] == orgPrefix[i] {
			length++
			rate += 2
		} else if str[i] == prefix[i] {
			length++
			rate++
		} else {
			break
		}
	}
	return length, rate
}

func splitDirectoryName(directoryName string) []string {
	var splitNames []string
	var lastIndex = -1
	for i, ch := range directoryName {
		if ch == '-' || ch == '_' || ch == ' ' || ch == '.' {
			if lastIndex >= 0 {
				splitNames = append(splitNames, directoryName[lastIndex:i])
				lastIndex = -1
			}
		} else {
			if lastIndex == -1 {
				lastIndex = i
			}
		}
	}

	if lastIndex >= 0 {
		splitNames = append(splitNames, directoryName[lastIndex:])
	}

	return splitNames
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
	} else if upperLetter == firstLetter {
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
		return directories
	}
	for _, entry := range entries {
		fileInfo, err := os.Stat(entry)
		if err == nil && fileInfo.IsDir() {
			relativePath := filepath.Base(entry)
			directories = append(directories, relativePath)
		}
	}
	return directories
}

// path delimiter must be "/" in this function.
func getFirstDirectory(path string) string {
	if len(path) == 0 {
		return ""
	}
	if len(path) >= 2 && path[1] == ':' {
		f := path[0]
		if f >= 'A' || f <= 'Z' || f >= 'a' || f <= 'z' {
			return path[0:2] + "/"
		}
	}
	if path[0] == '/' {
		return "/"
	}
	return ""
}

// IsAbs reports whether the path is absolute.
func IsAbs(path string) bool {
	if len(path) == 0 {
		return false
	}
	if len(path) >= 2 {
		f := path[0]
		if path[1] == ':' && (f >= 'A' || f <= 'Z' || f >= 'a' || f <= 'z') {
			return true
		}
	}
	return path[0] == '/'
}

func getVolumeName() string {
	path, err := filepath.Abs(".")
	if err != nil {
		return ""
	}
	return filepath.VolumeName(path)
}
