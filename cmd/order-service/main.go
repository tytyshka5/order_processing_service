package main

import (
	"fmt"
	"order_processing_service/internal/domain"
	"time"
)

func main() {

	computer := domain.Product{
		ID:           1,
		Name:         "Omni",
		PriceKopecks: 19999,
		Stock:        13,
	}

	mouse := domain.Product{
		ID:           2,
		Name:         "Logitech",
		PriceKopecks: 1999,
		Stock:        10,
	}

	mouseItem := domain.OrderItem{
		ProductID: mouse.ID,
		Quantity:  3,
		UnitPrice: mouse.PriceKopecks,
	}

	computerItem := domain.OrderItem{
		ProductID: computer.ID,
		Quantity:  2,
		UnitPrice: computer.PriceKopecks,
	}

	Order := domain.Order{
		ID:         1,
		CustomerID: 1,
		Items:      []domain.OrderItem{computerItem, mouseItem},
		Status:     domain.StatusNew,
		CreatedAt:  time.Now(),
	}

	orderSum := Order.Total()

	fmt.Printf("order=%d status=%s items=%d total=%d\n", Order.ID, Order.Status, len(Order.Items), orderSum)
}
