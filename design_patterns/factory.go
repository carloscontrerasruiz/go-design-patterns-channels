package main

import "fmt"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computadora struct {
	name  string
	stock int
}

func (c *Computadora) setStock(stock int) {
	c.stock = stock
}
func (c *Computadora) getStock() int {
	return c.stock
}
func (c *Computadora) setName(name string) {
	c.name = name
}
func (c *Computadora) getName() string {
	return c.name
}

type Laptop struct {
	Computadora
}

func newLaptop() IProduct {
	return &Laptop{
		Computadora: Computadora{
			name:  "Lenovo",
			stock: 25,
		},
	}
}

type Desktop struct {
	Computadora
}

func newDesktop() IProduct {
	return &Desktop{
		Computadora: Computadora{
			name:  "Desktop Computer",
			stock: 5,
		},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {
	switch computerType {
	case "laptop":
		return newLaptop(), nil
	case "desktop":
		return newDesktop(), nil
	default:
		return nil, fmt.Errorf("invalid type")
	}
}

func printNameAndStock(p IProduct) {
	fmt.Printf("Product name %s, with stock %d\n", p.getName(), p.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")

	printNameAndStock(laptop)
	printNameAndStock(desktop)
}
