package main
// nil_poniter
import "fmt"

type User struct {
    Name string
}

func GetUser(id int) *User {
    if id < 0 {
        return nil
    }
    return &User{Name: "John Doe"}
}

func nil_poniter() {
    user := GetUser(-1) 
    fmt.Println(user.Name) 
}

