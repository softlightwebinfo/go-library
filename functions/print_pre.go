package functions

import "fmt"

func PrintPre(data ...interface{}) {
	fmt.Printf("%+v", data)
}
