package domain

import (
	"time"
)
type OrderStatus string

const (
	StatusNew OrderStatus = "new"
	StatusConfirmed OrderStatus = "confirmed"
	StatusProcessing OrderStatus = "processing"
	StatusCompleted OrderStatus = "completed"
	StatusCancelled OrderStatus = "status"
)

type Order struct {
	ID int64
	CustomerID int64
	Items []OrderItem
	Status OrderStatus
	CreatedAt time.Time 
}

func (Product OrderItem) ProductPrice() int64 {
	return int64(Product.Quntity)*Product.UnitPrice
}

func (Order Order) OrderPrice() int64 {
	var orderSum int64
	for _,item:= range Order.Items {
		orderSum+=item.ProductPrice()
	}
	return orderSum
}