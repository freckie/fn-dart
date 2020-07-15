package utils

import (
	"bytes"
	"fn-dart/models"
	"strings"

	"github.com/suapapa/go_hangul/encoding/cp949"
)

var /* const */ stringsToRemove = []string{"  ", "\t", "\n", "혻"}

func TrimAll(value string) string {
	result := strings.TrimSpace(value)
	for _, str := range stringsToRemove {
		result = strings.Replace(result, str, "", -1)
	}
	result = strings.Replace(result, "\u00a0", " ", -1)
	return result
}

func ReadCP949(data string) (string, error) {
	br := bytes.NewReader([]byte(data))
	r, err := cp949.NewReader(br)
	if err != nil {
		return "", err
	}

	b := make([]byte, 10*1024)

	c, err := r.Read(b)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(b[:c])), nil
}

func MakePrevData(items []models.APIResultListItem) []string {
	result := make([]string, len(items))

	for idx, item := range items {
		result[idx] = item.RceptNo
	}

	return result
}

func IsContain(target string, arr []string) bool {
	for idx, _ := range arr {
		if arr[idx] == target {
			return true
		}
	}
	return false
}
