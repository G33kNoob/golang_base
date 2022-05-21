package main

import "fmt"

func main() {
	name1 := "Ahsan"
	name2 := "Zahwa"
	umur := 4
	if name1 == "Ahsan" && name2 == "Zahwa" {
		fmt.Println("dan true lah")
	}
	if name1 == "Ahsan" || name2 == "Zahwa" {
		fmt.Println("atau true lah")
	}
	if name2 == "Ahsan" || umur == 3 {
		fmt.Println("atau true lah")
	}
}
