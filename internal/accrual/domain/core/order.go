package core

type Product struct {
	description string
	price       string
}

type OrderReceipt struct {
	orderNumber string
	goods       []Product
}
