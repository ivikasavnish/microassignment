package products

import (
	"encoding/json"
	"fmt"
)

type ProdType int

const (
	Grocery ProdType = iota
	Electronics
	Kitchen
)

type Seller struct {
}

type Product struct {
	Name   string
	Type   ProdType
	Price  float64
	Seller Seller
}

func NewProduct(name string, ptype ProdType, price float64, seller Seller) Product {
	return Product{
		Name:   name,
		Type:   ptype,
		Price:  price,
		Seller: seller,
	}
}

func (p Product) Discount(disc float64) float64 {
	return p.Price * (1 - disc/100)
}
func (p Product) PrintJson(disc float64) string {
	out, err := json.MarshalIndent(p, "\t", "n")
	if err != nil {
		fmt.Errorf(" Error is %v", err)
	}
	return string(out)
}
