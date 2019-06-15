package main

import "fmt"

// product
type Chair interface {
	sitOn() string
	printSku() string
	printPrice() float32
}

type chair struct {
	kind  string
	legs  int16
	price float32
	sku   string
}

// factory
func modernChair() Chair {
	return &chair{
		kind:  "modern",
		legs:  4,
		price: 14000.00,
		sku:   "md-123456",
	}
}

func classicChair() Chair {
	return &chair{
		kind:  "classic",
		legs:  4,
		price: 18000.00,
		sku:   "cs-55555",
	}
}

func (c *chair) sitOn() string {
	return "Sitting on " + c.kind
}

func (c *chair) printSku() string {
	return c.sku
}

func (c *chair) printPrice() float32 {
	return c.price
}

// abstract factory
type FurnitureFactory interface {
	CreateChair() Chair
}

type modern struct{}
type classic struct{}

func newModernFurnitureFactory() FurnitureFactory {
	return &modern{}
}

func (m *modern) CreateChair() Chair {
	return modernChair()
}

func newClassicFurnitureFactory() FurnitureFactory {
	return &classic{}
}

func (c *classic) CreateChair() Chair {
	return classicChair()
}

func main() {
	// chair
	modern := newModernFurnitureFactory()
	modernChair := modern.CreateChair()

	fmt.Println(modernChair.printSku())
	fmt.Println(modernChair.printPrice())

	classic := newClassicFurnitureFactory()
	classicChair := classic.CreateChair()

	fmt.Println(classicChair.printSku())
	fmt.Println(classicChair.printPrice())
}
