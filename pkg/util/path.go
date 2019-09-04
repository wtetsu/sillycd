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

import "path/filepath"

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

// GetVolumeName returns the volume name of the current path.
func GetVolumeName() string {
	path, err := filepath.Abs(".")
	if err != nil {
		return ""
	}
	return filepath.VolumeName(path)
}

// SplitPath splits a full path string.
//
// - "/usr/local/bin" -> ["usr", "local", "bin"]
// - "c:/Program Files/Git" -> ["Program Files", "Git"]
func SplitPath(fullPath string, delimiter string) []string {
	var result []string
	lastIndex := 0
	for i := 0; i < len(fullPath); i++ {
		if fullPath[i:i+1] == delimiter {
			if lastIndex > 0 {
				s := fullPath[lastIndex:i]
				if len(s) > 0 && s != delimiter {
					result = append(result, s)
				}
			}
			lastIndex = i + 1
		}
	}

	s := fullPath[lastIndex:]
	if len(s) > 0 && s != delimiter {
		result = append(result, s)
	}

	return result
}

func SplitDirectoryName(directoryName string) []string {
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

// GetFirstDirectory returns the first directory path of "path"
// path delimiter must be "/" in this function.
//
// - /usr/local/bin      -> /
// - relative/path       -> (Current)
// - ./relative/path     -> (Current)
// - ../relative/path    -> (Parent)
// - ../../relative/path -> (Parent)
// - C:/Program Files    -> C:\
func GetFirstDirectory(path string) string {
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

	currentPath, err := filepath.Abs("")
	if err != nil {
		return ""
	}
	return currentPath
}
