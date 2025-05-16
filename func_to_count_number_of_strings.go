package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "hello world"
	Counting_String(string1)

}
func Counting_String(String string) {
	fmt.Println(strings.Count(String,"")-1)
}
