package main

import "fmt"

var number *int
var name *string

// func main() {
// 	var numberA int = 4
// 	var numberB *int = &numberA
	
// 	fmt.Println("numberA (value)   :", numberA)  // 4
// 	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220
	
// 	fmt.Println("numberB (value)   :", *numberB) // 4
// 	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220
// }

// // efek perubahan nilai pointer
// func main() {
// 	var numberA int = 4
// 	var numberB *int = &numberA
	
// 	fmt.Println("numberA (value)   :", numberA)
// 	fmt.Println("numberA (address) :", &numberA)
// 	fmt.Println("numberB (value)   :", *numberB)
// 	fmt.Println("numberB (address) :", numberB)
	
// 	fmt.Println("")
	
// 	numberA = 5
	
// 	fmt.Println("numberA (value)   :", numberA)
// 	fmt.Println("numberA (address) :", &numberA)
// 	fmt.Println("numberB (value)   :", *numberB)
// 	fmt.Println("numberB (address) :", numberB)
// }

// parameter pointer
func main() {
    var number = 4
    fmt.Println("before :", number) // 4

    change(&number, 10)
    fmt.Println("after  :", number) // 10
}

func change(original *int, value int) {
    *original = value
}
