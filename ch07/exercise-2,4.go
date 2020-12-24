package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
)

// Object interface
type Object interface {
	Volume() float64
}

// Cylinder
type Cylinder struct {
	Radius float64
	Height float64
}

func (o Cylinder) Volume() float64 {
	return math.Pi * o.Radius * o.Radius * o.Height
}

// Cube
type Cube struct {
	Width  float64
	Depth  float64
	Height float64
}

func (o Cube) Volume() float64 {
	return o.Width * o.Depth * o.Height
}

// Sphere
type Sphere struct {
	Radius float64
}

func (o Sphere) Volume() float64 {
	return 4 * math.Pi * o.Radius * o.Radius * o.Radius / 3
}

func inspect(instance interface{}) {
	fmt.Println()
	fmt.Println("From value to type")
	mr := reflect.ValueOf(instance)
	mrt := mr.Type()
	fmt.Printf("%v %v \n", mr, mrt)
	for i := 0; i < mr.NumField(); i++ {
		fmt.Printf("  %v: %v\n", mrt.Field(i).Name, mr.Field(i).Interface())
	}
	for j := 0; j < mr.NumMethod(); j++ {
		fmt.Printf("  %v -> %v\n", mrt.Method(j).Name, mr.Method(j).Type())
	}

	fmt.Println("================================================")
}

func main() {
	cylinder := Cylinder{3, 12}
	fmt.Println(cylinder.Volume())
	inspect(cylinder)

	cube := Cube{5, 5, 5}
	fmt.Println(cube.Volume())
	inspect(cube)

	sphere := Sphere{7}
	fmt.Println(sphere.Volume())
	inspect(sphere)

	var f *os.File
	inspect(f)
}
