package main

import "fmt"
import "database/sql"
import _ "github.com/lib/pq"

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	connStr := "postgres://postgres:admin123@localhost:5432/db_belajar_golang?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("select id, name, grade from tb_student where age = $1", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}

func main() {
	sqlQuery()
}
