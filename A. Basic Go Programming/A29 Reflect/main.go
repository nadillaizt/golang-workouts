package main

import "fmt"
import "reflect"

type student struct {
    Name  string
    Grade int
}

// func main() {
//     var number = 23
//     var reflectValue = reflect.ValueOf(number)

//     fmt.Println("tipe  variabel :", reflectValue.Type())

//     if reflectValue.Kind() == reflect.Int {
//         fmt.Println("nilai variabel :", reflectValue.Int())
//     }
// }

// func (s *student) getPropertyInfo() {
//     var reflectValue = reflect.ValueOf(s)

//     if reflectValue.Kind() == reflect.Ptr {
//         reflectValue = reflectValue.Elem()
//     }

//     var reflectType = reflectValue.Type()

//     for i := 0; i < reflectValue.NumField(); i++ {
//         fmt.Println("nama      :", reflectType.Field(i).Name)
//         fmt.Println("tipe data :", reflectType.Field(i).Type)
//         fmt.Println("nilai     :", reflectValue.Field(i).Interface())
//         fmt.Println("")
//     }
// }

// func main() {
//     var s1 = &student{Name: "wick", Grade: 2}
//     s1.getPropertyInfo()
// }

func (s *student) SetName(name string) {
    s.Name = name
}

func main() {
    var s1 = &student{Name: "john wick", Grade: 2}
    fmt.Println("nama :", s1.Name)

    var reflectValue = reflect.ValueOf(s1)
    var method = reflectValue.MethodByName("SetName")
    method.Call([]reflect.Value{
        reflect.ValueOf("wick"),
    })

    fmt.Println("nama :", s1.Name)
}