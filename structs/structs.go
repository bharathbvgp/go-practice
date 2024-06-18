package main

import "fmt"

// creating a person stucture

type person struct {
	name string
	age int
}

func newperson(name string) *person {
	p := person {
		name : name,
	}
	p.age = 23;
	return &p;
}


func main() {
	fmt.Println(person { name : "anil" , age : 23 });
	fmt.Println(person { name : "gurumurthy" , age : 32});
	fmt.Println(*newperson("pallavi"));
	s := person {
		name : "Anil",
		age : 22,
	}
	fmt.Println(s.age);
	fmt.Println(s.name);

	// getting a reference of person s

	sr := &s;
	fmt.Println(sr.name);
	sr.name = "Aravind";
	fmt.Println(s.name);
	fmt.Println(sr.name);

	// anonymous struct

	dog := struct {
		name string
		isGood bool
	}{
		name : "kia",
		isGood: false,
	}

	fmt.Println(dog);

}