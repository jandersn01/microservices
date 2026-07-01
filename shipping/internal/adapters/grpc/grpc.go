package grpc 

import (
	"context"
	"fmt"

	"github.com/jandersn01/microservices/shipping/internal/application/core/domain"
	"github.com/jandersn01/microservices-proto/golang/shipping"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) CalculateDeadline(ctx context.Context, request *shipping.CalculateDeadlineRequest) (*shipping.CalculateDeadlineResponse, error) {

	log.WithContext(ctx).Info("Calculating shipping deadline")

	var domainItems []domain.OrderItem
	for _, item := range request.Items {
		domainItems = append(domainItems, domain.OrderItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	result, err := a.application.CalculateDeadLine(ctx, request.OrderId, domainItems)

	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("falha ao calcular o prazo de frete: %v", err))
	}

	return &shipping.CreateShippingResponse{
		OrderId:              result.OrderID,
		DeliveryDeadlineDays: result.DeliveryDeadlineDays,
	}, nil

}