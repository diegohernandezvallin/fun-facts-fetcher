package util

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}

	return value
}

func StrToReadCloser(str string) io.ReadCloser {
	r := ioutil.NopCloser(strings.NewReader(str))

	return r
}
