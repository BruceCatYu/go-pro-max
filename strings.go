package gopromax

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// ParseLocalTime Parse string to local time
func ParseLocalTime(s string) (t time.Time) {
	t, _ = time.ParseInLocation(time.RFC3339, s, time.Local)
	return
}

// EqualFold is strings.EqualFold, ASCII only. It reports whether s and t
// are equal, ASCII-case-insensitively.
func EqualFold(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if Lower(s[i]) != Lower(t[i]) {
			return false
		}
	}
	return true
}

// Lower returns the ASCII lowercase version of b.
func Lower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
func StringCut(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

// ToMd5
func ToMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// GBKToUTF8 trans GBK string to UTF8 string
func GBKToUTF8(b []byte) []byte {
	if d, err := io.ReadAll(transform.NewReader(bytes.NewReader(b), simplifiedchinese.GBK.NewDecoder())); err != nil {
		return []byte{}
	} else {
		return d
	}
}

// StripSpaces remove all space char in string
func StripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return r
	}, str)
}
