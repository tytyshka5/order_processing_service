# Order Processing Service

Простой сервис обработки заказов, написанный на Go.

Проект демонстрирует основы предметно-ориентированного проектирования (Domain Model), работу со структурами, методами, пакетами и организацией кода в Go.

## Возможности

- создание товаров (`Product`);
- создание заказов (`Order`);
- хранение списка товаров в заказе (`OrderItem`);
- расчет общей стоимости заказа;
- хранение статуса заказа;
- хранение даты создания заказа.

## Структура проекта

```
order-processing-service/
├── cmd/
│   └── order-service/
│       └── main.go            # Точка входа
├── internal/
│   └── domain/
│       ├── order.go           # Модель заказа
│       └── product.go         # Модель товара
├── go.mod
└── README.md
```

## Модели

### Product

Описывает товар.

```go
type Product struct {
    ID            int64
    Name          string
    PriceKopecks  int64
    Stock         int
}
```

### OrderItem

Описывает товар внутри заказа.

```go
type OrderItem struct {
    ProductID int64
    Quantity  int
    UnitPrice int64
}
```

### Order

Описывает заказ.

```go
type Order struct {
    ID         int64
    CustomerID int64
    Items      []OrderItem
    Status     OrderStatus
    CreatedAt  time.Time
}
```

## Расчет стоимости заказа

Стоимость заказа вычисляется методом:

```go
func (o Order) OrderPrice() int64
```

Он суммирует стоимость всех позиций:

```
Количество × Цена
```

Например:

```
2 × 19999 = 39998 копеек
```

## Запуск

Клонировать репозиторий:

```bash
git clone https://github.com/tytyshka5/order_processing_service.git
```

Перейти в проект:

```bash
cd order_processing_service
```

Запустить приложение:

```bash
go run ./cmd/order-service
```

## Пример использования

```go
computer := domain.Product{
    ID:            1,
    Name:          "Omni",
    PriceKopecks:  19999,
    Stock:         13,
}

item := domain.OrderItem{
    ProductID: computer.ID,
    Quantity:  2,
    UnitPrice: computer.PriceKopecks,
}

order := domain.Order{
    ID:         1,
    CustomerID: 1,
    Items: []domain.OrderItem{
        item,
    },
    Status:    domain.OrderStatusNew,
    CreatedAt: time.Now(),
}

fmt.Println(order.OrderPrice())
```

Вывод:

```
39998
```

## Технологии

- Go 1.24+
- Standard Library

## Планы развития

- [ ] Валидация заказа
- [ ] Добавление и удаление товаров
- [ ] Изменение статуса заказа
- [ ] Проверка остатка товара
- [ ] Unit-тесты
- [ ] REST API
- [ ] Хранение данных в PostgreSQL

## Автор

GitHub: https://github.com/tytyshka5