// defer = mengakhiri execute before selesai
// exit = end secara paksa

package main

import "fmt"
import "os"

// func main() {
//     defer fmt.Println("halo")
//     fmt.Println("selamat datang")
// }

// func main() {
//     orderSomeFood("pizza")
//     orderSomeFood("burger")
// }

// func orderSomeFood(menu string) {
//     defer fmt.Println("Terimakasih, silakan tunggu")
//     if menu == "pizza" {
//         fmt.Print("Pilihan tepat!", " ")
//         fmt.Print("Pizza ditempat kami paling enak!", "\n")
//         return
//     }

//     fmt.Println("Pesanan anda:", menu)
// }

// func main() {
//     number := 3

//     if number == 3 {
//         fmt.Println("halo 1")
//         defer fmt.Println("halo 3")
//     }

//     fmt.Println("halo 2")
// }

// func main() {
//     number := 3

//     if number == 3 {
//         fmt.Println("halo 1")
//         func() {
//             defer fmt.Println("halo 3")
//         }()
//     }

//     fmt.Println("halo 2")
// }

// os exit = Semua statement setelah exit tidak akan dieksekusi, termasuk juga defer.
func main() {
    defer fmt.Println("halo")
    os.Exit(1)
    fmt.Println("selamat datang")
}