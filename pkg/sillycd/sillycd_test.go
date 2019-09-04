package sillycd

import (
	"testing"
)

func Test01(t *testing.T) {
	if computeDirectoryScore("abc", "abc") != 999999999 {
		t.Fatal("failed test")
	}
	if computeDirectoryScore("ABC", "ABC") != 999999999 {
		t.Fatal("failed test")
	}
	if computeDirectoryScore("ABC", "abc") != 15 {
		t.Fatal("failed test")
	}
	if computeDirectoryScore("ABC", "abc") != 15 {
		t.Fatal("failed test")
	}
}
func Test02(t *testing.T) {
	if computeDirectoryScore("foo-bar-baz", "fbb") != 30 {
		t.Fatal("failed test")
	}

	if computeDirectoryScore("foo-bar-baz", "fbab") != 40 {
		t.Fatal("failed test")
	}

	if computeDirectoryScore("foo-bar-baz", "fb") != 20 {
		t.Fatal("failed test")
	}

	if computeDirectoryScore("foo-bar-baz", "f") != 10 {
		t.Fatal("failed test")
	}

	if computeDirectoryScore("FOO-BAR-BAZ", "fbb") != 15 {
		t.Fatal("failed test")
	}
}

func Test03(t *testing.T) {
	if computeDirectoryScore("foo_bar_baz", "fbb") != 30 {
		t.Fatal("failed test")
	}

	if computeDirectoryScore("foo_bar_baz", "fbab") != 40 {
		t.Fatal("failed test")
	}

	if computeDirectoryScore("foo_bar_baz", "fb") != 20 {
		t.Fatal("failed test")
	}

	if computeDirectoryScore("foo_bar_baz", "f") != 10 {
		t.Fatal("failed test")
	}

	if computeDirectoryScore("FOO_BAR_BAZ", "fbb") != 15 {
		t.Fatal("failed test")
	}
}

func Test04(t *testing.T) {
	if computeDirectoryScore("foo bar baz", "fbb") != 30 {
		t.Fatal("failed test")
	}

	if computeDirectoryScore("foo bar baz", "fbab") != 40 {
		t.Fatal("failed test")
	}

	if computeDirectoryScore("foo bar baz", "fb") != 20 {
		t.Fatal("failed test")
	}

	if computeDirectoryScore("foo bar baz", "f") != 10 {
		t.Fatal("failed test")
	}

	if computeDirectoryScore("FOO BAR BAZ", "fbb") != 15 {
		t.Fatal("failed test")
	}
}

func Test05(t *testing.T) {
	if computeDirectoryScore("Program Files", "pf") != 10 {
		t.Fatal("failed test")
	}
	if computeDirectoryScore("Program Files (x86)", "pf") != 10 {
		t.Fatal("failed test")
	}
}
