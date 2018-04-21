package main

import (
	"errors"
	"fmt"
	"math"
)

type Sphere struct {
	Radius float64
}

// *Spher is a reference to the sphere struc.  This is equivalent
// to function prototypes in javascript
// s is equivalent to this in js
func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}

func (s *Sphere) Volume() float64 {
	radiusCubed := s.Radius * s.Radius * s.Radius
	return (float64(4) / float64(3)) * math.Pi * radiusCubed
}

func createSphere() {
	s := Sphere{
		Radius: 5,
	}
	fmt.Println(s.SurfaceArea())
	fmt.Println(s.Volume())
}

type Triangle struct {
	base   float64
	height float64
}

// Triagle is a value
func (t Triangle) changeValueBase(f float64) {
	t.base = f
	return
}

// Triangle is a reference
func (t *Triangle) changeReferenceBase(f float64) {
	t.base = f
	return
}

func valueVsReferenceMethods() {
	// prints 3 because the Triangle referenced is a value
	// If you need to operate on a struct, but do not want to
	// modify the original initialization of a struct, use a value. (this one)
	t := Triangle{base: 3, height: 1}
	t.changeValueBase(4)
	fmt.Println(t.base)

	// prints 4 because the Triange referenced is a reference
	// If you need to modify (or mutate) the original initialization
	// of a struct, use a pointer (this example)
	u := Triangle{base: 3, height: 1}
	u.changeReferenceBase(4)
	fmt.Println(u.base)
}

type Robot interface {
	PowerOn() error
}

type T850 struct {
	Name string
}

func (a *T850) PowerOn() error {
	return nil
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}

func Boot(r Robot) error {
	return r.PowerOn()
}

func turnOnRobots() {
	t := T850{
		Name: "The Terminator",
	}

	r := R2D2{
		Broken: true,
	}

	err := Boot(&r)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

	err = Boot(&t)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}
}

func main() {
	createSphere()
	valueVsReferenceMethods()
	turnOnRobots()
}
