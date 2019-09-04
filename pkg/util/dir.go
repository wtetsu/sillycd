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

package util

import (
	"os"
	"path/filepath"
	"strings"
)

// FindDirectories returns a directory list in targetDirectory
func FindDirectories(targetDirectory string, targetName string) []string {
	if strings.HasPrefix(targetName, ".") {
		return []string{
			targetName,
		}
	}
	patterns := generateFindDictionaryPatterns(targetDirectory, targetName)
	var directories []string
	for _, pattern := range patterns {
		foundDirectories := findDirectoriesByPattern(pattern)
		directories = append(directories, foundDirectories...)
	}
	return directories
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
