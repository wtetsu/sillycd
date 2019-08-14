package sillycd

import (
	"testing"

	"github.com/wtetsu/sillycd/pkg/sillycd"
)

func Test01(t *testing.T) {
	result := sillycd.Shorten("abc-def-ghi")

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
	if !contains(result, "ab-de-gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdegh") {
		t.Fatal("failed test")
	}
}

func Test02(t *testing.T) {
	result := sillycd.Shorten("ab-de-gh")

	if !contains(result, "ab-de-gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdegh") {
		t.Fatal("failed test")
	}
	if !contains(result, "a-d-g") {
		t.Fatal("failed test")
	}
	if !contains(result, "adg") {
		t.Fatal("failed test")
	}
	if !contains(result, "ab-de-gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdegh") {
		t.Fatal("failed test")
	}
}

func Test03(t *testing.T) {
	result := sillycd.Shorten("abc-d-ghi")

	if !contains(result, "ab-d-gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdgh") {
		t.Fatal("failed test")
	}
	if !contains(result, "a-d-g") {
		t.Fatal("failed test")
	}
	if !contains(result, "adg") {
		t.Fatal("failed test")
	}
	if !contains(result, "ab-d-gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdgh") {
		t.Fatal("failed test")
	}
}

func Test04(t *testing.T) {
	result := sillycd.Shorten("abc_def_ghi")

	if !contains(result, "abc_def_ghi") {
		t.Fatal("failed test")
	}
	if !contains(result, "abcdefghi") {
		t.Fatal("failed test")
	}
	if !contains(result, "a_d_g") {
		t.Fatal("failed test")
	}
	if !contains(result, "adg") {
		t.Fatal("failed test")
	}
	if !contains(result, "ab_de_gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdegh") {
		t.Fatal("failed test")
	}
}

func Test05(t *testing.T) {
	result := sillycd.Shorten("ab_de_gh")

	if !contains(result, "ab_de_gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdegh") {
		t.Fatal("failed test")
	}
	if !contains(result, "a_d_g") {
		t.Fatal("failed test")
	}
	if !contains(result, "adg") {
		t.Fatal("failed test")
	}
	if !contains(result, "ab_de_gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdegh") {
		t.Fatal("failed test")
	}
}

func Test06(t *testing.T) {
	result := sillycd.Shorten("abc_d_ghi")

	if !contains(result, "ab_d_gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdgh") {
		t.Fatal("failed test")
	}
	if !contains(result, "a_d_g") {
		t.Fatal("failed test")
	}
	if !contains(result, "adg") {
		t.Fatal("failed test")
	}
	if !contains(result, "ab_d_gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdgh") {
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
