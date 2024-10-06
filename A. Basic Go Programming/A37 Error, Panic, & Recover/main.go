// package main

// import (
//     "fmt"
//     "strconv"
// )

// func main() {
//     var input string
//     fmt.Print("Type some number: ")
//     fmt.Scanln(&input)

//     var number int
//     var err error
//     number, err = strconv.Atoi(input)

//     if err == nil {
//         fmt.Println(number, "is number")
//     } else {
//         fmt.Println(input, "is not number")
//         fmt.Println(err.Error())
//     }
// }

package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

// // custom error
// func main() {
//     var name string
//     fmt.Print("Type your name: ")
//     fmt.Scanln(&name)

//     if valid, err := validate(name); valid {
//         fmt.Println("halo", name)
//     } else {
//         fmt.Println(err.Error())
//     }
// }

// // panic error
// func main() {
//     var name string
//     fmt.Print("Type your name: ")
//     fmt.Scanln(&name)

//     if valid, err := validate(name); valid {
//         fmt.Println("halo", name)
//     } else {
//         panic(err.Error())
//         fmt.Println("end")
//     }
// }

// // recover
// func catch() {
//     if r := recover(); r != nil {
//         fmt.Println("Error occured", r)
//     } else {
//         fmt.Println("Application running perfectly")
//     }
// }

// func main() {
//     defer catch()

//     var name string
//     fmt.Print("Type your name: ")
//     fmt.Scanln(&name)

//     if valid, err := validate(name); valid {
//         fmt.Println("halo", name)
//     } else {
//         panic(err.Error())
//         fmt.Println("end")
//     }
// }

// recover pada IIFE
func main() {
    data := []string{"superman", "aquaman", "wonder woman"}

    for _, each := range data {

        func() {

            // recover untuk IIFE dalam perulangan
            defer func() {
                if r := recover(); r != nil {
                    fmt.Println("Panic occured on looping", each, "| message:", r)
                } else {
                    fmt.Println("Application running perfectly")
                }
            }()

            panic("some error happen")
        }()

    }
}