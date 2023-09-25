package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	testYears(12)
}

func testYears(age int) {
	if age > 10 {
		fmt.Printf("Hey, you are %v age years old", age)
	}
}
