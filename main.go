package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var db, err = Jstow[User]("db.json")
	if err != nil {
		fmt.Println("Error, code broke!")
		return
	}
	// db.Insert(User{Name: "Jace", Age: 20})

	fmt.Println(db.All())
}
