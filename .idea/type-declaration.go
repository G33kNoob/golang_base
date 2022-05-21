package main

import "fmt"

func main2() {
	type Married bool
	var MarriedStat Married = true
	fmt.Println(MarriedStat)
	type Ktp int8
	nomorKtp := 12121212
	fmt.Println(nomorKtp)
}
