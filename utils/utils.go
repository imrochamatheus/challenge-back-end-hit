package utils

import "os"

func ReadQueryFile(path string) (string, error) {
	data, err := os.ReadFile(path)

	return string(data), err
}