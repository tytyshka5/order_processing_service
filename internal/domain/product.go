package domain

type Product struct {
	ID 	int64
	Name string
	PriceKopecks int64
	Stock int
}

type OrderItem struct {
	ProductID int64
	Quntity int
	UnitPrice int64
}