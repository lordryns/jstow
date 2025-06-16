package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var db = Jstow[User]("db.json")

	var command string
	for {
		fmt.Print("Enter command [i, s, a]: ")
		fmt.Scan(&command)

		if command == "i" {
			var id string
			var name string
			var age string

			fmt.Print("Enter id: ")
			fmt.Scan(&id)

			fmt.Print("Enter name: ")
			fmt.Scan(&name)

			fmt.Print("Enter age: ")
			fmt.Scan(&age)

			age2, err := strconv.Atoi(age)
			if err != nil {
				fmt.Println("Age must be a number!")
			}

			user := User{Id: id, Name: name, Age: age2}
			err2 := db.Insert(user)

			if err != nil {
				fmt.Println(err2)
			}
		} else if command == "s" {
			var fieldName string
			var targetValue string

			fmt.Print("Enter field name: ")
			fmt.Scan(&fieldName)

			fmt.Print("Enter target value: ")
			fmt.Scan(&targetValue)

			var res, err = db.Search(fieldName, targetValue)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(res)
		} else if command == "a" {
			fmt.Println(db.All())
		} else {
			fmt.Println("Invalid command!")
		}
	}
}
