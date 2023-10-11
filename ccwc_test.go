package ccwc

import (
	"os"
	"testing"
)

func TestCounts(t *testing.T) {
	filename := "test.txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Errorf("ccwc: %s: No such file or directory", filename)
	}
	c, l, w, m := Counts(content)
	if c != "342190" {
		t.Errorf("Wrong c | Expected 342190, got %q", c)
	}
	if l != "7145" {
		t.Errorf("Wrong c | Expected 7145, got %q", l)
	}
	if w != "58164" {
		t.Errorf("Wrong c | Expected 58164, got %q", w)
	}
	if m != "339292" {
		t.Errorf("Wrong c | Expected 339292, got %q", m)
	}
}
