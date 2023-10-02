package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db := sql.DB{}
	fmt.Println("Hello, world!")
	err := get(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("No errors")
}

func get(db sql.DB) error {
	_, err := db.Exec("SELECT * FROM ", "users")
	if err != nil {
		return err
	}
	fmt.Println("Created")
	return nil
}
