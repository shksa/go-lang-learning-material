package main

import (
	"fmt"
)

// Item is a cutom type based on struct type
type Item struct {
	id       string
	price    float64
	quantity int
}

// Cost is a method on *Item type
func (item *Item) Cost() float64 {
	return item.price * float64(item.quantity)
}

// SpecialItem is a custom type that embeds the Item type
type SpecialItem struct {
	Item      // Anonymous field (embedding)
	catalogID int
}

// LuxuryItem is a type that embed's the type Item
type LuxuryItem struct {
	Item   // embedded type
	markup float64
}

// Cost is a method for the type *LuxuryItem that overrides the method Cost of the embedded type Item
func (luxuryItem *LuxuryItem) Cost() float64 {
	return luxuryItem.Item.Cost() * luxuryItem.markup
}

func main() {
	specialItem := SpecialItem{
		Item{"1xd4r", 23.56, 3}, 1234569,
	}
	fmt.Printf("%#v \n", specialItem)
	fmt.Printf("%v \n", specialItem.Cost()) // Can call Item type's Cost() method on SpecialItem type
	fmt.Printf("specialItem.id: %v, specialItem.price: %v, specialItem.quantity: %v, specialItem.catalogID: %v", specialItem.id, specialItem.price, specialItem.quantity, specialItem.catalogID)

	luxuryItem := LuxuryItem{
		Item{"1xd4r", 23.56, 3}, 3.5,
	}
	fmt.Printf("%#v \n", luxuryItem)
	fmt.Printf("%v \n", luxuryItem.Cost())
}
