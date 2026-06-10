package main

import "fmt"

func main() {

	fmt.Println("Hi")
	fmt.Println("ritvik")

	congrats := "Happy birthday ritvik"
	fmt.Println(congrats)

	number := 34
	fmt.Printf("The type is %T\n", number)

	val1, val2 := 35.2, "Hello again ritvik"
	println(val1, val2)

	// truncation

	accountNumber := 88.6
	accountNumberInt := int(accountNumber)
	println(accountNumberInt)
}
