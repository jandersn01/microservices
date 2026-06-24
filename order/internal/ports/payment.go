package ports

import "github.com/jandersn01/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(payment *domain.Order) error
}