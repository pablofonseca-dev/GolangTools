package utils

import (
	"fmt"
	"strconv"
)

func ConvertStrToFloat(copy string) float64 {
	number, error := strconv.ParseFloat(copy, 64)
	if error != nil {
		fmt.Printf("Error in ConvertStrToFloat\n%v", error)
	}
	return number
}
