type student struct {
    name        string
    height      float64
    age         int32
    isGraduated bool
    hobbies     []string
}

var data = student{
    name:        "wick",
    height:      182.5,
    age:         26,
    isGraduated: false,
    hobbies:     []string{"eating", "sleeping"},
}

// NUMERIC TO BINER
fmt.Printf("%b\n", data.age)
// 11010


fmt.Printf("%c\n", 1400)
// ո

fmt.Printf("%c\n", 1235)
// ӓ

// NUMERIC TO BASE 10 NUMERIC 
fmt.Printf("%d\n", data.age)
// 26

// NUMERIC DECIMAL TO SCIENTIFIC NOTATION
fmt.Printf("%e\n", data.height)
// 1.825000e+02

fmt.Printf("%E\n", data.height)
// 1.825000E+02

// NUMERIC DECIMAL TO height decimal bs ditentukan
fmt.Printf("%f\n", data.height)
// 182.500000

fmt.Printf("%.9f\n", data.height)
// 182.500000000

fmt.Printf("%.2f\n", data.height)
// 182.50

fmt.Printf("%.f\n", data.height)
// 182

// sama kek sblmnya tp lebih besar digitnya
fmt.Printf("%e\n", 0.123123123123)
// 1.231231e-01

fmt.Printf("%f\n", 0.123123123123)
// 0.123123

fmt.Printf("%g\n", 0.123123123123)
// 0.123123123123

