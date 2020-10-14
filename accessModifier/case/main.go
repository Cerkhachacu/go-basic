package cases

import "fmt"

// Student public
var Student = struct {
	Number string
	Value  int
}{}

func init() {
	Student.Number = "20"
	Student.Value = 30

	fmt.Println("--> package case !imported")
}
