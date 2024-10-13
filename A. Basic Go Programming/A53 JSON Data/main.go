package main

import "encoding/json"
import "fmt"

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	// var jsonString = `{"Name": "john wick", "Age": 27}`
	// var jsonData = []byte(jsonString)

	// var data User

	// var err = json.Unmarshal(jsonData, &data)
	// if err != nil {
	//     fmt.Println(err.Error())
	//     return
	// }

	// fmt.Println("user :", data.FullName)
	// fmt.Println("age  :", data.Age)

	// ----------------------------------------------------------------
	var jsonString = `[
    {"Name": "john wick", "Age": 27},
    {"Name": "ethan hunt", "Age": 32}
]`

	var data []User

	var err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1:", data[0].FullName)
	fmt.Println("user 2:", data[1].FullName)
	// ----------------------------------------------------------------
	// var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	// var jsonData, err = json.Marshal(object)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// var jsonString = string(jsonData)
	// fmt.Println(jsonString)
}
