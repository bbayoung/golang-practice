package main

import "fmt"

type Item struct {
	name string
	price float64
	quantity int
}

type DiscountItem struct {
	Item
	discountRate float64
}

func (t *Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func (t * DiscountItem) Cost() float64 {
	return t.price * float64(t.quantity) * (float64(100) - t.discountRate) / float64(100)
}

func main() {
	product := Item{"Sample", 25000, 3}
	saleProduct := DiscountItem{Item{"Sample2", 50000, 3}, 10.0}

	fmt.Println(product.Cost())
	fmt.Println(saleProduct.Cost())
}
