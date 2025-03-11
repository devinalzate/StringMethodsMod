package utils

type Product struct {
	Name     string
	Price    float64
	Overview string
}

func CreateProduct(name_new string, price_new float64, overview_new string) Product {
	new_product := Product{}
	new_product.Name = name_new
	new_product.Price = price_new
	new_product.Overview = overview_new
	return new_product
}

func (p *Product) ChangeOverview(overview_update string) {
	p.Overview = overview_update
}

func (p *Product) RecalculatePriceWithIva() {
	p.Price = p.Price + (p.Price * 0.20)
}

func (p *Product) GetName() (text string) {
	text = "Producto: " + p.Name
	return text
}
