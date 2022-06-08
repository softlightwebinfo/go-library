package functions

import (
	"fmt"
	"strconv"
)

func StringToInt64(id string) (int64, error) {
	return strconv.ParseInt(id, 10, 64)
}

func Int64ToString(id int64) string {
	return strconv.FormatInt(id, 10)
}

func Float32ToString(id float32) string {
	return fmt.Sprintf("%v", id)
}

func BoolToString(show bool) string {
	return If(show, "Si", "No").(string)
}

func Price(price float32) string {
	return fmt.Sprintf("%0.2f€", price)
}

func Percent(percent float32) string {
	return fmt.Sprintf("%0.2f%%", percent)
}

type parserData interface {
	int | float64 | string
}

func Unique[T parserData](intSlice []T) []T {
	keys := make(map[T]bool)
	list := []T{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
