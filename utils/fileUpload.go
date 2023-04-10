package utils

import (
	"os"
)

func DeleteFile(filename string) (string, error) {
	err := os.Remove(filename)
	if err != nil {
		return "", err
	}

	return "Success", nil
}
