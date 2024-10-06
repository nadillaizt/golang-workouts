package main

import "fmt"
import "strings"

// // strings contains
// func main() {
//     var isExists = strings.Contains("john wick", "wick")
//     fmt.Println(isExists)
// }

// // strings hasprefix
// func main() {
// 	var isPrefix1 = strings.HasPrefix("john wick", "jo")
// 	fmt.Println(isPrefix1) // true

// 	var isPrefix2 = strings.HasPrefix("john wick", "wi")
// 	fmt.Println(isPrefix2) // false
// }

// // strings hasSuffix
// func main() {
// 	var isSuffix1 = strings.HasSuffix("john wick", "ic")
// 	fmt.Println(isSuffix1) // false

// 	var isSuffix2 = strings.HasSuffix("john wick", "ck")
// 	fmt.Println(isSuffix2) // true
// }

// // strings count
// func main() {
// 	var howMany = strings.Count("ethan hunt", "t")
// 	fmt.Println(howMany) // 2
// }

// // strings index
// func main() {
// 	var index1 = strings.Index("ethan hunt", "ha")
// 	fmt.Println(index1) // 2
// }

// // strings replace
// func main() {
// 	var text = "banana"
// 	var find = "a"
// 	var replaceWith = "o"

// 	var newText1 = strings.Replace(text, find, replaceWith, 1)
// 	fmt.Println(newText1) // "bonana"

// 	var newText2 = strings.Replace(text, find, replaceWith, 2)
// 	fmt.Println(newText2) // "bonona"

// 	var newText3 = strings.Replace(text, find, replaceWith, -1)
// 	fmt.Println(newText3) // "bonono"
// }

// // strings repeat
// func main() {
// 	var str = strings.Repeat("na", 4)
// 	fmt.Println(str) // "nananana"
// }

// // strings split
// func main() {
// 	var string1 = strings.Split("the dark knight", " ")
// 	fmt.Println(string1) // output: ["the", "dark", "knight"]

// 	var string2 = strings.Split("batman", "")
// 	fmt.Println(string2) // output: ["b", "a", "t", "m", "a", "n"]
// }

// // strings join
// func main() {
// 	var data = []string{"banana", "papaya", "tomato"}
// 	var str = strings.Join(data, "-")
// 	fmt.Println(str) // "banana-papaya-tomato"
// }

// strings tolower
func main() {
	var str = strings.ToLower("aLAy")
	fmt.Println(str) // "alay"
}

// // strings toupper
// func main() {
// 	var str = strings.ToUpper("eat!")
// 	fmt.Println(str) // "EAT!"
// }
