// Package ccwc is a simplified clone of wc written in Go.
package ccwc

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

// Counts computes number of bytes, lines, words, and characters in content
func Counts(content []byte) (string, string, string, string) {
	c := fmt.Sprint(len(content))
	l := fmt.Sprint(bytes.Count(content, []byte("\n")))
	w := fmt.Sprint(len(bytes.Fields(content)))
	m := fmt.Sprint(utf8.RuneCountInString((string(content))))

	return c, l, w, m
}
