package main

import (
	"errors"
	"fmt"

	"order_processing_service/internal/domain"
	"order_processing_service/internal/repository/memory"
)

func main() {

	repo := memory.NewProductRepository()

	product1 := repo.Create(domain.Product{
		Name:         "Телефон",
		PriceKopecks: 199999,
		Stock:        13,
	})
	fmt.Printf("created product: id=%d name=%s\n", product1.ID, product1.Name)
	product2 := repo.Create(domain.Product{
		Name:         "Ноутбук",
		PriceKopecks: 1999,
		Stock:        10,
	})
	fmt.Printf("created product: id=%d name=%s\n", product2.ID, product2.Name)
	var id int64 = 1
	product3, err := repo.GetByID(id)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err.Error())
	} else {
		fmt.Printf("found product: id=%d name=%s\n", product3.ID, product3.Name)
	}
	id = 999
	_, err = repo.GetByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			fmt.Printf("product %d not found\n", id)
		}
	}
	items := repo.List()
	fmt.Println("Products:")
	for _, v := range items {
		fmt.Printf("%s, \n", v.Name)
	}
	changedID := items[0].ID
	items[0].Name = "Asus"

	storedProduct, err := repo.GetByID(changedID)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			fmt.Printf("product %d not found\n", id)
		}
	} else {
		fmt.Printf("found product: id=%d name=%s\n", storedProduct.ID, storedProduct.Name)
	}
}
