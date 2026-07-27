package main

import (
	"context"
	"fmt"
	grpcAdapter "go-service-client/internal/adapter/grpc"
	"go-service-client/internal/config"
	"go-service-client/internal/entities"
	"go-service-client/internal/usecases"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.Load()

	// client := commons.NewRestClient(cfg.HttpClient.BaseURL, cfg.HttpClient.Timeout)

	conn, err := grpc.NewClient(cfg.GrpcClient.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed to connect to gRPC server:", err)
		return
	}
	defer conn.Close()

	ctx := context.Background()

	// mathAdapter := adapter.NewMathAdapter(adapter.Paths{
	// 	Addition: adapter.PathConfig{
	// 		Url:     "/api/v1/math/addition",
	// 		Timeout: cfg.Timeout.Addition,
	// 	},
	// 	Multiply: adapter.PathConfig{
	// 		Url:     "/api/v1/math/multiply",
	// 		Timeout: cfg.Timeout.Multiply,
	// 	},
	// 	Subtraction: adapter.PathConfig{
	// 		Url:     "/api/v1/math/subtraction",
	// 		Timeout: cfg.Timeout.Subtraction,
	// 	},
	// }, client)

	mathAdapter := grpcAdapter.NewMathGrpcAdapter(conn)

	u := usecases.NewMathUsecase(mathAdapter)
	u.FetchAllOperation(ctx, entities.Numbers{A: 3, B: 2})
}
