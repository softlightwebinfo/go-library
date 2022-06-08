package helpers

import "fmt"

func Contains(slice []string, value string) (bool, int, string) {
	fmt.Println(slice, value)
	for index, s := range slice {
		if s == value {
			return true, index, s
		}
	}

	return false, -1, ""
}
