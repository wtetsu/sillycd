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
	"path/filepath"
	"sort"

	"github.com/wtetsu/sillycd/pkg/util"
)

// PickOutDirectory picks out one directory.
func PickOutDirectory(targetPath string) string {
	return PickOutDirectoryWithFunction(targetPath, util.FindDirectories)
}

// PickOutDirectoryWithFunction picks out one directory.
func PickOutDirectoryWithFunction(orgTargetDirectory string, doGetDirectories func(string, string) []string) string {
	if len(orgTargetDirectory) == 0 {
		return ""
	}

	var targetPath string
	firstLetter := orgTargetDirectory[0]
	if firstLetter == '/' || firstLetter == '\\' {
		targetPath = util.GetVolumeName() + filepath.ToSlash(orgTargetDirectory)
	} else {
		targetPath = filepath.ToSlash(orgTargetDirectory)
	}

	firstDirectory := util.GetFirstDirectory(targetPath)
	splittedPath := util.SplitPath(targetPath, "/")

	if len(splittedPath) <= 0 {
		return filepath.Join(firstDirectory)
	}

	resultDirectory := traverseDirectoriesByTheEnd(firstDirectory, splittedPath, doGetDirectories)
	return resultDirectory
}

func traverseDirectoriesByTheEnd(firstTargetDirectory string, wishTargets []string, doGetDirectories func(string, string) []string) string {
	var doFindDirectory func(string, int) string

	doFindDirectory = func(targetDirectory string, currentIndex int) string {
		if currentIndex >= len(wishTargets) {
			return targetDirectory
		}
		name := wishTargets[currentIndex]
		var candidate string
		directories := doGetDirectories(targetDirectory, name)
		sortedDirectories := sortDirectoriesByScore(directories, name)
		for _, dir := range sortedDirectories {
			candidate = doFindDirectory(filepath.Join(targetDirectory, dir), currentIndex+1)
			if candidate != "" {
				break
			}
		}
		return candidate
	}

	return doFindDirectory(firstTargetDirectory, 0)
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
