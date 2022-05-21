package main

import "fmt"

func main1000() {
	var (
		name1 = "Ahsan"
		name2 = "Abdillah"
		name3 = "Zahwa"
	)
	fmt.Println(name1 == name2)
	fmt.Println(name1 < name2)
	fmt.Println(name1 > name2)
	fmt.Println(name1 > name3)

	number1 := 200
	number2 := 400
	fmt.Println(number2 <= number1)
}
