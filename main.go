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

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/wtetsu/sillycd/pkg/sillycd"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Wrong argument")
		os.Exit(1)
	}

	var specifiedName = os.Args[1]
	var target string

	if filepath.IsAbs(specifiedName) {
		target = specifiedName
	} else {
		target = "./" + specifiedName
	}

	foundDirectory := sillycd.PickOutDirectory(target)

	if foundDirectory == "" {
		os.Exit(1)
	}
	fmt.Println(foundDirectory)
}
