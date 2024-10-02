package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	randomizer := rand.New(rand.NewSource(10))
// 	fmt.Println("random ke-1:", randomizer.Int()) // 5221277731205826435
// 	fmt.Println("random ke-2:", randomizer.Int()) // 3852159813000522384
// 	fmt.Println("random ke-3:", randomizer.Int()) // 8532807521486154107
// }

// func main() {
// 	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
// 	fmt.Println("random ke-1:", randomizer.Int())
// 	fmt.Println("random ke-2:", randomizer.Int())
// 	fmt.Println("random ke-3:", randomizer.Int())
// }

// func main() {
// 	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
// 	fmt.Println("random int:", randomizer.Int())
// 	fmt.Println("random float32:", randomizer.Float32())
// 	fmt.Println("random uint:", randomizer.Uint32())
// }

var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	fmt.Println("random string 5 karakter:", randomString(5))
}
