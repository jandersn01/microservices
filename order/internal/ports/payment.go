package ports

import "github.com/huseyinbabal/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(payment *domain.Order) error
}