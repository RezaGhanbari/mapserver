package utils

import "strconv"

func ConvertStringToFloat64(s string) (float64 , error) {
	number, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, err
	}
	return number, nil
}










