package entities

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price int
}

func (product *Product) String() string {
	return fmt.Sprintf("- ID: %v, Name: %v, Price: %v", product.ID, product.Name, product.Price)
}
