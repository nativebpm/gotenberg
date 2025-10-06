package model

type Party struct {
	Name         string
	Address      string
	CityStateZip string
	Contact      string
	Email        string
}

type Payment struct {
	Method      string
	CheckNumber string
}

type Item struct {
	Description string
	Price       string
	Last        bool
}

type Invoice struct {
	InvoiceNumber string
	CreatedAt     string
	DueAt         string
	From          Party
	To            Party
	Payment       Payment
	Items         []Item
	Total         string
}
