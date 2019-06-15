package main

import "fmt"

// Abstraction
type CarFactory interface {
	BuildCar() string
}

type car struct {
	brand string
	model string
	price float32
}

// concreate factory for Benz
type Benz struct{}

func (b *Benz) BuildCar() string {
	return "Building new benz..."
}

// concreate factory for BMW
type BMW struct{}

func (b *BMW) BuildCar() string {
	return "Building new bmw..."
}

// factory method
func NewCar(brand string) CarFactory {
	switch brand {
	case "benz":
		return &Benz{}
	case "bmw":
		return &BMW{}
	default:
		return nil
	}
}

func main() {
	myCar := NewCar("benz")
	fmt.Println(myCar.BuildCar())
}
