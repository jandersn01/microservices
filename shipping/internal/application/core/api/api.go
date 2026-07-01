package api 

import (
	"context"
	"github.com/jandersn01/microservices/shipping/internal/application/core/domain"
)

type Application struct {}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) CalculateDeadLine(ctx context.Context, orderId int64, items []domain.OrderItem) (domain.Shipping, error) {
	var totalUnits int32

	for _, item := range items {
		totalUnits += item.Quantity
	}

	deadlineDays := int32(1) + (totalUnits / 5)

	shippingResult := domain.Shipping{
		OrderID: orderId,
		DeliveryDeadlineDays: deadlineDays,
	}
	return shippingResult, nil
}