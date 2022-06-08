package helpers

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
