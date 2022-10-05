package main

import (
	"fmt"
)

func main() {
	msg := RenderImportantMessageFromBears()
	fmt.Println(msg)
}

func RenderImportantMessageFromBears() string {
	return "jello world"
}
