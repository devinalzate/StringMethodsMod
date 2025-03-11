package utils

type Product struct {
	name     string
	price    float64
	overview string
}

func CreateProduct(name_new string, price_new float64, overview_new string) Product {
	new_product := Product{}
	new_product.name = name_new
	new_product.price = price_new
	new_product.overview = overview_new
	return new_product
}

func (p *Product) ChangeOverview(overview_update string) {
	p.overview = overview_update
}

func (p *Product) RecalculatePriceWithIva() {
	p.price = p.price * 0.20
}
