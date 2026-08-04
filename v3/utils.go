package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Overview string
	Stock    int
}

type ProductsList struct {
	List []Product
}

func (pl *ProductsList) CreateProduct(name_new string, price_new float64, overview_new string) Product {
	new_product := Product{}
	new_product.Name = name_new
	new_product.Price = price_new
	new_product.Overview = overview_new
	new_product.Stock = 0
	pl.List = append(pl.List, new_product)

	return new_product
}

//Metodos a aplciar teniendolos en lista
func (pl *ProductsList) GetProductOfList(name string) *Product {
	for _, d := range pl.List {
		if d.Name == name {
			return &d
		}
	}
	return nil
}

func main() {
}

//metodos de solo producto

func (p *Product) GetProduct() {
	fmt.Printf("Producto: %v\nPrecio: %v\nDescripcion: %v\nStock: %v\n", p.Name, p.Price, p.Overview, p.Stock)
}

func (p *Product) ChangeOverview(overview_update string) {
	p.Overview = overview_update
}

//Precio final
func (p *Product) PriceWithIva() float64 {
	p.Price = p.Price + (p.Price * 0.20)
	return p.Price
}

//Metodos de stock
func (p *Product) GetStock() int {
	return p.Stock
}

func (p *Product) StockDecrese() {
	p.Stock = p.Stock - 1
}

func (p *Product) UpdateStock(new_stock int) {
	p.Stock = new_stock
}
