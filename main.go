package testego

import "fmt"

type Fruit struct {
	Name  string
	Price float64
}
type Market interface {
	Selling() bool
	Marketing() bool
}

func (fruit Fruit) Selling() {
	fmt.Println("Selling: ", fruit.Name)
}

func (fruit Fruit) Marketing() {
	fmt.Println("Marketing: ", fruit.Price)
}

func CreateFruit(name string, price float64) Fruit {
	return Fruit{
		Name:  name,
		Price: price,
	}
}

func Start(fruit Fruit) {
	fruit.Selling()
	fruit.Marketing()
}
