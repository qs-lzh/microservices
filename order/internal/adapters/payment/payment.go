package payment

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/qs-lzh/microservices-proto/golang/payment"
	"github.com/qs-lzh/microservices/order/internal/application/core/domain"
)
