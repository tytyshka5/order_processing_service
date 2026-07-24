package main

import (
	"fmt"
	"github.com/tytyshka5/order-processing-service/internal/domain"
	"time"
)

func main() {

	Computer := domain.Product{
		ID:           1,
		Name:         "Omni",
		PriceKopecks: 19999,
		Stock:        13,
	}

	ComputerItem := domain.OrderItem{
		ProductID: Computer.ID,
		Quntity:   2,
		UnitPrice: Computer.PriceKopecks,
	}

	ComputerOrder := domain.Order{
		ID:         1,
		CustomerID: 1,
		Items:      []domain.OrderItem{ComputerItem},
		Status:     domain.StatusNew,
		CreatedAt:  time.Now(),
	}

	orderSum := ComputerOrder.OrderPrice()

	fmt.Printf("order=%d status=%s items=%d total=%d", ComputerOrder.ID, ComputerOrder.Status, len(ComputerOrder.Items), orderSum)
}
