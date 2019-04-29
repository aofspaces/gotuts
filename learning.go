package main

import "fmt"

var (
	name, location string = "Zero", "bangkok"
	date           int    = 1
)

func main() {
	fmt.Println(name, date, location)
	phoneBrand := "iPhone"

	showDetail := func() {
		fmt.Println(name, date, location, phoneBrand)
	}

	showDetail()
}
