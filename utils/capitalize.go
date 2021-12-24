package utils

import (
	"fmt"
	"strings"
)

func Capitalize(copy string) (string, error) {
	if copy == "" {
		return "", fmt.Errorf("copy of \"\" could not be processed because is empty")
	}
	return strings.Title(copy), nil
}
