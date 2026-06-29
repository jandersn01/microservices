package api

import (
	"github.com/jandersn01/microservices/order/internal/application/core/domain"
	"github.com/jandersn01/microservices/order/internal/ports"
)

type Application struct{
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication (db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(order *domain.Order) (*domain.Order, error) {
	err := a.db.Save(order)
	if err != nil {
		return nil, err
	}


	paymentErr := a.payment.Charge(order)


	if paymentErr != nil {
		order.Status = "Canceled"
	} else {
		order.Status = "Paid"
	}
	
	saveErr := a.db.Save(order)
	if saveErr != nil {
		return nil, saveErr
	}

	if paymentErr != nil{
		return nil, paymentErr
	}

	return order, nil


}