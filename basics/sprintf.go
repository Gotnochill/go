package main

import "fmt"

func sprintfUse() {

	const name = "Walter Mukherjee"
	const openRate = 30.55

	msg := fmt.Sprintf(
		"Hi %s , your open rate is %.5f percent",
		name,
		openRate,
	)

	fmt.Println(msg)
}
