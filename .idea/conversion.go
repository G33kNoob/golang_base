package main

import "fmt"

func main20() {
	var (
		nilaiawal   int8  = 100
		nilaikedua  int16 = int16(nilaiawal)
		nilaiketiga int32 = int32(nilaikedua)
	)
	fmt.Println(nilaiawal)
	fmt.Println(nilaikedua)
	fmt.Println(nilaiketiga)

	var (
		name = "Ahsan"
		getByte = name[0]
		changeToString string = string(getByte)
	)
	fmt.Println(name)
	fmt.Println(getByte)
	fmt.Println(changeToString)
}
