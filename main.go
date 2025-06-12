package main

import "fmt"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var db = Jstow[User]("user.json")

	for key, user := range db.All() {
		if user.Name == "Ryns" {
			fmt.Println(db.All()[key])
		}
	}
}
