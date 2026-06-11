package main

import "fmt"

func main() {

	messageLen := 35
	maxMessageLen := 20

	if messageLen <= maxMessageLen {
		fmt.Println("Message Sent")
	} else {
		fmt.Println("Message not sent")
	}

	s1 := "Ritvik"
	s2 := "Mukherjee"

	fmt.Println(concat(s1, s2))
}
