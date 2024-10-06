package main

import "fmt"

// func main() {
// 	var chicken map[string]int
// 	chicken = map[string]int{}
	
// 	chicken["januari"] = 50
// 	chicken["februari"] = 40
	
// 	fmt.Println("januari", chicken["januari"]) // januari 50
// 	fmt.Println("mei",     chicken["mei"])     // mei 0
// }

func main() {
	var data map[string]int
data["one"] = 1
// akan muncul error!

data = map[string]int{}
data["one"] = 1
// tidak ada error
}

// cara horizontal
var chicken1 = map[string]int{"januari": 50, "februari": 40}

// cara vertical
var chicken2 = map[string]int{
    "januari":  50,
    "februari": 40,
}

var chicken3 = map[string]int{}
var chicken4 = make(map[string]int)
var chicken5 = *new(map[string]int)

var chicken = map[string]int{
    "januari":  50,
    "februari": 40,
    "maret":    34,
    "april":    67,
}

for key, val := range chicken {
    fmt.Println(key, "  \t:", val)
}