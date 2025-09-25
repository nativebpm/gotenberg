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

var InvoiceData = Invoice{
	InvoiceNumber: "123",
	CreatedAt:     "January 1, 2023",
	DueAt:         "February 1, 2023",
	From: Party{
		Name:         "Sparksuite, Inc.",
		Address:      "12345 Sunny Road",
		CityStateZip: "Sunnyville, TX 12345",
	},
	To: Party{
		Name:    "Acme Corp.",
		Contact: "John Doe",
		Email:   "john@example.com",
	},
	Payment: Payment{
		Method:      "Check",
		CheckNumber: "1000",
	},
	Items: []Item{
		{Description: "Website design", Price: "$300.00", Last: false},
		{Description: "Hosting (3 months)", Price: "$75.00", Last: false},
		{Description: "Domain name (1 year)", Price: "$10.00", Last: true},
	},
	Total: "$385.00",
}
