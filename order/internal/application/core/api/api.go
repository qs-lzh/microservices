package api

import (
	"github.com/qs-lzh/microservices/order/internal/application/core/domain"
	"github.com/qs-lzh/microservices/order/internal/ports"
)

type Application struct {
	db      ports.DBPort
	payment payment.Adapter
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		return domain.Order{}, paymentErr
	}

	return order, nil
}
