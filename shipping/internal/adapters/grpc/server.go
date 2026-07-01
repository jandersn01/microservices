package grpc 

import (
	"fmt"
	"log"
	"net"

	"github.com/jandersn01/microservices/shipping/config"
	"github.com/jandersn01/microservices/shipping/internal/ports"
	"github.com/jandersn01/microservices-proto/golang/shipping"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
	port int
	server *grpc.Server
	shipping.UnimplementedShippingServiceServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("falha ao escutar na porta %d, erro: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	a.server = grpcServer

	shipping.RegisterShippingServer(grpcServer, a)

	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	log.Printf("Iniciando serviço de shipping na porta %d ...", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("falha ao servir grpc: %v", err)
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}