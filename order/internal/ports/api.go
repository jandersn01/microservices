package ports

import "https://github.com/jandersn01/microservices/order/internal/application/core/domain"

type APIPort interface{
	PlaceOrder(order *domain.Order) (*domain.Order, error)
}