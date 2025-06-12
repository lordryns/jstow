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
```
```

After creating the struct, we use it in creating our database
```go 
var db = Jstow[User]("users.json")
```
``` 

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

```
```
```
```
