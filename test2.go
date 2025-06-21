package main

import "fmt"

type User2 struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func test() {
	var db, _ = Jstow[User2]("db.json")

	users, _ := db.Search("Name", "Jason")
	for _, user := range users {
		user.Age = 17
		db.Update("Name", "Jason", user)
	}
	fmt.Println(db.All())

}
