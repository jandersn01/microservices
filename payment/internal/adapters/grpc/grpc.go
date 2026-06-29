package grpc

import (
	"context"
	"fmt"

	"github.com/huseyinbabal/microservices/payment/internal/application/core/domain"
	"github.com/jandersn01/microservices-proto/golang/payment"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	log.WithContext(ctx).Info("Creating payment...")

	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)

	result, err := a.api.Charge(ctx, newPayment)

	if err != nil {
		if err.Error() == "payment over 1000 is not allowed"{
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to charge: %v", err))
	}
		return &payment.CreatePaymentResponse{PaymentId: result.ID}, nil
}
