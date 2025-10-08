package ports

import (
	"github.com/qs-lzh/microservices/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(*domain.Order) error
}
