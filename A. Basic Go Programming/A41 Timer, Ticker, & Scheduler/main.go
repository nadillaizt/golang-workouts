package main

import "fmt"
import "time"
import "os"

// // time sleep bersifat blocking
// // jadi statement di bwh nya ga akan jalan sampe selesai timer nya
// func main () {
//     fmt.Println("start")
//     time.Sleep(time.Second * 4)
//     fmt.Println("after 4 seconds")
// }

// // bisa jadi scheduler
// func main () {
// 	for true {
// 		fmt.Println("Hello !!")
// 		time.Sleep(1 * time.Second)
// 	}
// }

// // time new timer
// // setelah jeda waktu yang ditentukan sebuah data akan
// // dikirimkan lewat channel C. Penggunaan fungsi ini harus
// // diikuti dengan statement untuk penerimaan data dari channel C
// func main() {
// 	var timer = time.NewTimer(4 * time.Second)
// 	fmt.Println("start")
// 	<-timer.C
// 	fmt.Println("finish")
// }

// // timer afterfunc
// func main() {
// 	var ch = make(chan bool)

// 	time.AfterFunc(4*time.Second, func() {
// 		fmt.Println("expired")
// 		ch <- true
// 	})

// 	fmt.Println("start")
// 	<-ch
// 	fmt.Println("finish")
// }

// // time after
// // fungsi timer.After() akan mengembalikan data channel,
// // sehingga perlu menggunakan tanda <-
// func main() {
// 	<-time.After(4 * time.Second)
// 	fmt.Println("expired")
// }


// // scheduler using ticker
// func main() {
//     done := make(chan bool)
//     ticker := time.NewTicker(time.Second)

//     go func() {
//         time.Sleep(10 * time.Second) // wait for 10 seconds
//         done <- true
//     }()

//     for {
//         select {
//         case <-done:
//             ticker.Stop()
//             return
//         case t := <-ticker.C:
//             fmt.Println("Hello !!", t)
//         }
//     }
// }

// kombinasi timer n goroutine
func timer(timeout int, ch chan<- bool) {
    time.AfterFunc(time.Duration(timeout)*time.Second, func() {
        ch <- true
    })
}

func watcher(timeout int, ch <-chan bool) {
    <-ch
    fmt.Println("\ntime out! no answer more than", timeout, "seconds")
    os.Exit(0)
}

func main() {
    var timeout = 5
    var ch = make(chan bool)

    go timer(timeout, ch)
    go watcher(timeout, ch)

    var input string
    fmt.Print("what is 725/25 ? ")
    fmt.Scan(&input)

    if input == "29" {
        fmt.Println("the answer is right!")
    } else {
        fmt.Println("the answer is wrong!")
    }
}