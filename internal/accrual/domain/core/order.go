package core

type Product struct {
	description string
	price       string
}

type OrderReceipt struct {
	orderNumber string
	goods       []Product
}

type Answer struct {
	Number  int    `json:"order"`
	Status  string `json:"status"`
	Accrual int    `json:"accrual,omitempty"`
}
