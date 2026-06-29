package domain 

import "time"

type OrderItem struct{
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

type Order struct {
	ID         int64       `json:"id"`
	CustomerID int64       `json:"customer_id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  int64       `json:"created_at"`
}

func NewOrder(CustomerID int64, OrderItems []OrderItem) *Order {
	return &Order{
		CreatedAt:  time.Now().Unix(),
		Status:     "pending",
		CustomerID: CustomerID,
		OrderItems: OrderItems,
	}
}

func (o *Order) TotalPrice() float32 {
	var totalPrice float32
	for _, orderItem := range o.OrderItems {
		totalPrice += orderItem.UnitPrice * float32(orderItem.Quantity)
	}
	return totalPrice
}

func (o *Order) TotalItems() int32{
	var totalQuantity int32
	for _, orderItem := range o.OrderItems {
		totalQuantity += orderItem.Quantity
	}
	return totalQuantity
}