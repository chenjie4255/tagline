package tagline

import (
	"testing"
)

func TestStringUpper(t *testing.T) {
	type TestBody struct {
		Name  string `tl:"upper"`
		Title string `tl:"lower"`
	}

	bd := TestBody{"abc", "MANAGER"}

	if err := Process(&bd); err != nil {
		t.Fatalf("failed to call Post, %s", err)
	}

	if bd.Name != "ABC" {
		t.Fatal("name should be equal 'ABC'")
	}

	if bd.Title != "manager" {
		t.Fatal("title should be equal 'manager'")
	}
}
