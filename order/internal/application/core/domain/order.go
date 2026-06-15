package domain 

import "time"

type OrderItem struct{
	ProductCode string `json:"product_code"`
	UnitPrice float32 `json:"unit_price"`
	Quantity int32 `json:"quantity"`
}

type Order struct {
	ID int64 `json:"id"`
	CustumerID int64 `json:"customer_id"`
	Status string `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt int64 `json:"created_at"`

}

func NewOrder(CustumerID int64, OrderItems []OrderItem) *Order {
	return Order{
		CreatedAt: time.Now().Unix(),
		Status: "pending",
		CustumerID: CustumerID,
		OrderItems: OrderItems,
	}
}