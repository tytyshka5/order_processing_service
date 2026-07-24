package domain

import (
	"time"
)

type OrderStatus string

const (
	StatusNew        OrderStatus = "new"
	StatusConfirmed  OrderStatus = "confirmed"
	StatusProcessing OrderStatus = "processing"
	StatusCompleted  OrderStatus = "completed"
	StatusCancelled  OrderStatus = "cancelled"
)

type Order struct {
	ID         int64
	CustomerID int64
	Items      []OrderItem
	Status     OrderStatus
	CreatedAt  time.Time
}

func (item OrderItem) Total() int64 {
	return int64(item.Quantity) * item.UnitPrice
}

func (order Order) Total() int64 {
	var orderSum int64
	for _, item := range order.Items {
		orderSum += item.Total()
	}
	return orderSum
}
