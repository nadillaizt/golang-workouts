// package main

// import "fmt"
// import "net/http"

// func index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "apa kabar!")
// }

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "halo!")
// 	})

// 	http.HandleFunc("/index", index)

// 	fmt.Println("starting web server at http://localhost:8080/")
// 	http.ListenAndServe(":8080", nil)
// }

package main

import "fmt"
import "html/template"
import "net/http"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var data = map[string]string{
            "Name":    "john wick",
            "Message": "have a nice day",
        }

        var t, err = template.ParseFiles("template.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }

        t.Execute(w, data)
    })

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}