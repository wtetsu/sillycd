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
	"strings"

	"github.com/wtetsu/sillycd/pkg/util"
)

// "foo-bar-baz"
//   "f":   10
//   "fo":  20
//   "fb":  20
//   "fbb": 30
//   "foo-bar-baz": 999999999
func computeDirectoryScore(directoryName string, specifiedName string) int {
	if len(specifiedName) > len(directoryName) {
		return 0
	}
	if directoryName == specifiedName {
		return 999999999
	}
	var score int

	names := util.SplitDirectoryName(directoryName)
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

	if score > 0 {
		return score
	}

	score = roughCount(directoryName, specifiedName)

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

func roughCount(name string, target string) int {
	lowerName := strings.ToLower(name)
	lowerTarget := strings.ToLower(target)

	score := 0
	nameIndex := 0
	var targetIndex int
	for targetIndex = 0; targetIndex < len(lowerTarget); targetIndex++ {
		ch := lowerTarget[targetIndex : targetIndex+1]

		nextIndex := index(lowerName, ch, nameIndex)
		if nextIndex < 0 {
			break
		}

		score++
		nameIndex = nextIndex

		if nameIndex >= len(lowerName) {
			break
		}
	}
	if targetIndex < len(target) {
		score = 0
	}
	return score
}

func index(str string, chars string, startFrom int) int {
	strLen := len(str)
	charsLen := len(chars)

	foundIndex := -1
	for i := startFrom; i < strLen; i++ {
		if i > strLen-charsLen {
			break
		}
		if str[i:i+charsLen] == chars {
			foundIndex = i
			break
		}
	}
	return foundIndex
}
