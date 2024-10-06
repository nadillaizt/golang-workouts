// package main

// import "fmt"

// func main() {
//     var secret interface{}

//     secret = "ethan hunt"
//     fmt.Println(secret)

//     secret = []string{"apple", "manggo", "banana"}
//     fmt.Println(secret)

//     secret = 12.4
//     fmt.Println(secret)
// }

// var data map[string]interface{}

// var data map[string]interface{}

// // data = map[string]interface{}{
// //     "name":      "ethan hunt",
// //     "grade":     2,
// //     "breakfast": []string{"apple", "manggo", "banana"},
// // }

// // var data map[string]any

// // data = map[string]any{
// //     "name":      "ethan hunt",
// //     "grade":     2,
// //     "breakfast": []string{"apple", "manggo", "banana"},
// // }

package main

import "fmt"
import "strings"

func main() {
    var secret interface{}

    secret = 2
    var number = secret.(int) * 10
    fmt.Println(secret, "multiplied by 10 is :", number)

    secret = []string{"apple", "manggo", "banana"}
    var gruits = strings.Join(secret.([]string), ", ")
    fmt.Println(gruits, "is my favorite fruits")
}