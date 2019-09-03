package main

import (
	"fmt"
	"math"

	"github.com/max_workspace/golang_learn/interface/car"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

var (
	areaIntf Shaper
)

func main() {
	// 测试接口的类型断言
	testInterfaceAssertion()

	// 测试接口的类型判断
	testInterfaceTypeSwitch(areaIntf)
	testInterfaceTypeSwitch(nil)
	testInterfaceTypeSwitch(nil)

	// 测试结构体、集合以及高阶函数
	testInterfaceS()
}

func testInterfaceAssertion() {
	fmt.Println("test interface assertion")
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1

	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
	fmt.Println("\n")
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func testInterfaceTypeSwitch(it Shaper) {
	fmt.Println("test interface type-switch")
	switch t := it.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
	fmt.Println("\n")
}

func testInterfaceS() {
	fmt.Println("test interface s")

	ford := &car.Car{"Fiesta", "Ford", 2008}
	bmw := &car.Car{"XL 450", "BMW", 2011}
	merc := &car.Car{"D600", "Mercedes", 2009}
	bmw2 := &car.Car{"X 800", "BMW", 2008}

	allCars := car.Cars([]*car.Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *car.Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)

	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := car.MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
	fmt.Println("\n")
}
