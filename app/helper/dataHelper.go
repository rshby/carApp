package helper

import "strconv"

func ToFloat(price string) float64 {
	num, _ := strconv.ParseFloat(price, 64)
	return num
}
