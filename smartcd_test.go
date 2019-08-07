package main

import (
	"fmt"
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func Test01(t *testing.T) {
	result := shorten("abc-def-ghi")

	fmt.Println(result)

	if !contains(result, "abc-def-ghi") {
		t.Fatal("failed test")
	}
	if !contains(result, "abcdefghi") {
		t.Fatal("failed test")
	}
	if !contains(result, "a-d-g") {
		t.Fatal("failed test")
	}
	if !contains(result, "adg") {
		t.Fatal("failed test")
	}
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
