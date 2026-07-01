package domain

type OrderItem struct {
	ProductCode string
	Quantity    int32
}

type Shipping struct {
	OrderID   int64
	DeliveryDeadlineDays int32
}