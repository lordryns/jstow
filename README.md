# Jstow 

Jstow brings the functionality of SQL to JSON, in essence, Jstow allows you to store data in a Row/Column format in a JSON file as opposed to the traditional Key, Value pair method you may be familiar with.

To get started, Run 

```bash 
go get github.com/lordryns/jstow
``` 

After everything is installed, you can go straight to using it in your project.

Before getting started with Jstow, it is important to understand its design principles.

- Every file is a table
- Values must be consistent hence the usage of structs

Every file is a table, and we define to a struct to control the row/column behaviour of the table eg 
```go 
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

After creating the struct, we use it in creating our database
```go 
var db = Jstow[User]("users.json")
```


This will locate a users.json file or create one if it doesen't already exist while keeping **name** and **age** as the columns.

Below is a full example:

```go 
package main 
import (
  "fmt"
  "github.com/lordryns/jstow"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main(){
  var db = Jstow[User]("users.json")
  var values = db.All() // use to get the entire table
	fmt.Println(values)

}
``` 


### All important methods 
- All 
Use this to get the entire table in the form of a map
```go 
allData := db.All()
``` 

- Insert
This inserts a new row at the bottom of the table
```go 
user = User{...}
err := db.Insert(user)
``` 

- Search 
This method takes an unconventional approach and instead treats the columns like strings, see the example below
```go 
users, err := db.Search("Name", "John") // this finds and returns any row that matches 'John' under the column 'Name'
``` 
Note: Search returns a list so in this case it returns []User and so you'd need to loop through it to access its contents.


- Update 
The update method is slightly more complex than what we've been doing so far, it requires the Search method to work properly.
You'd need to search for the rows you want to change, loop through them, update the changes and then update them
```go
	users, _ := db.Search("Name", "John") // search returns a list of users
	for _, user := range users {
		user.Age = 17 // update every user's age
		db.Update("Name", "John", user) // update every user
	}
  ```

- Delete
This deletes everything that fits a specific requirement (similar to search), in the example below, any row with the name John will be deleted.
```go
err := db.Delete("Name", "John")
```


