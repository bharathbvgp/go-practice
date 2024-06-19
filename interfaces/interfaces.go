package main

import (
	"fmt"
	"math"
)

// 	I need an interface of type geometry

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width , height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width;
}
// (r rect) is the receiver mean to which type the method is associated
func (r rect) perimeter() float64 {
	return  (2*r.width + 2*r.height);
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius;
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius;
}

func measure(g geometry) {
	fmt.Println(g);
	fmt.Println(g.area());
	fmt.Println(g.perimeter())
}

func modify(r rect) {
	r.width = 100;
	fmt.Println("Modified data .." , r);
}

func main() {
	r := rect { width: 10 , height: 10 }
	measure(r); // passed by value
	fmt.Println("before modifying")
	fmt.Println(r);
	fmt.Println("After calling function..");
	modify(r);
	fmt.Println(r);

	// creating a circle instance
	c := circle {
		radius: 10,
	}
	measure(c);
}