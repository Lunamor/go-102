// Declare a struct type to maintain information about a person.  Declare a
// function that creates new values of your type.  Call this function from main
// and display the value.
//
// Template available at: http://play.golang.org/p/ta6oFzjgwn
package main

// Add your imports here.
import "fmt"

// Declare a struct type `person` to maintain information about a person.
type person struct {
    name string
    age int
}

// Declare a function that creates new values of your `person` type.
func makePerson(name string, age int) person {
    return person{name, age}
}

func main() {
    p1 := makePerson("Mike", 10)
    fmt.Printf("%s is %d\n", p1.name, p1.age)
    fmt.Println(p1)
}
