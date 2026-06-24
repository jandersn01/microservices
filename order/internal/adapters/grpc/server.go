package grpc

import (
	"context"
	"fmt"
	"net"
	"log"

	"github.com/jandersn01/microservices-proto/golang/order"
	"github.com/jandersn01/microservices/order/internal/application/core/domain"
	"github.com/jandersn01/microservices/order/internal/ports"
	"google.golang.org/grpc"
	"github.com/jandersn01/microservices/order/config"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	app  ports.APIPort
	port int 
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		app: api,
		port: port,
	}
}

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	for _, orderItem := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	newOrder := domain.NewOrder(int64(request.CustomerId), orderItems)
	result, err := a.app.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{
		OrderId: int32(result.ID),
	}, nil
}

func (a Adapter) Run()  {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}
	grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server, error: %v", err)
	}
}