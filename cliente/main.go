package main

import (
	"context"
	"log"
	"time"

	"github.com/jandersn01/microservices-proto/golang/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. Conecta no microsserviço Order (porta 3000)
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Falha ao conectar no Order: %v", err)
	}
	defer conn.Close()

	client := order.NewOrderClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Enviando requisição de compra para o microsserviço Order...")

	// 2. Monta o payload (Dados da compra)
	req := &order.CreateOrderRequest{
		CustomerId: 999,
		OrderItems: []*order.OrderItem{
			{
				ProductCode: "CELULAR-XYZ",
				UnitPrice:   1500.50,
				Quantity:    1,
			},
		},
	}

	// 3. Dispara a requisição
	res, err := client.Create(ctx, req)
	if err != nil {
		log.Fatalf("Erro ao processar a compra: %v", err)
	}

	// 4. Imprime o resultado final!
	log.Printf("SUCESSO! A comunicação entre os microsserviços funcionou.")
	log.Printf("ID do Pedido retornado: %d", res.OrderId)
}