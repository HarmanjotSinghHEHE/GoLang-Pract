package main

import "fmt"

type //person struct {
	name string
	Age  int
}

//func updatename(x *person, newname string) {
	x.name = newname
}
//func updateage(x *person, newage int) {
	x.Age = newage
}

func main() {
	person := person{name: "Harman", Age: 21}
	//person.name = "Shaina"
	//person.Age = 21

	fmt.Println(" OG\n", person.name, "\n", person.Age)
	//fmt.Println("UPDATED")
	updatename(&person, "Shaina")
	defer updateage(&person, 23) //testing defer
	fmt.Println(" UP\n", person.name, "\n", person.Age)

}
