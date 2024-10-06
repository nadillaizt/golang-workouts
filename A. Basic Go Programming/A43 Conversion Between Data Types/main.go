package main

import "fmt"
import "strconv"

// // strconv itoa
// func main() {
//     var str = "124"
//     var num, err = strconv.Atoi(str)

//     if err == nil {
//         fmt.Println(num) // 124
//     }
// }

// // strconv parseint
// func main() {
// 	var str = "124"
// 	var num, err = strconv.ParseInt(str, 10, 64)

// 	if err == nil {
// 		fmt.Println(num) // 124
// 	}
// }

// // strconv formatint
// func main() {
// var num = int64(24)
// var str = strconv.FormatInt(num, 8)

// fmt.Println(str) // 30
// }

// // strconv parsefloat
// func main() {
//     var str = "24.12"
// var num, err = strconv.ParseFloat(str, 32)

// if err == nil {
//     fmt.Println(num) // 24.1200008392334
// }
// }

// strconv formatfloat
func main() {
	var num = float64(24.12)
	var str = strconv.FormatFloat(num, 'f', 6, 64)

	fmt.Println(str) // 24.120000
}
