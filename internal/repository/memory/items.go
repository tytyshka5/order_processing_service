package memory

import (
	"order_processing_service/internal/domain"
	"sort"
)

type ProductRepository struct {
	nextID   int64
	products map[int64]domain.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: make(map[int64]domain.Product),
		nextID:   1,
	}
}

func (r *ProductRepository) Create(item domain.Product) domain.Product {
	item.ID = r.nextID
	r.products[item.ID] = item
	r.nextID++
	return item
}

func (r *ProductRepository) GetByID(id int64) (domain.Product, error) {
	if item, found := r.products[id]; found {
		return item, nil
	}
	return domain.Product{}, domain.ErrProductNotFound
}

func (r *ProductRepository) List() []domain.Product {
	result := make([]domain.Product, 0, len(r.products))

	for _, product := range r.products {
		result = append(result, product)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}
