package printer

import "fmt"

// Println wraps fmt.Println
func Println(str string) string {
	fmt.Println(str)
	return str
}
