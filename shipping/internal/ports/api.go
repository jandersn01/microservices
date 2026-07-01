package ports

import (
	"context"
	"github.com/jandersn01/microservices/shipping/internal/application/core/domain"
)

type APIPort interface {
	CalculateDeadline(ctx context.Context, orderId int64, items []domain.OrderItem) (domain.Shipping, error)
}