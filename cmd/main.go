package main

import (
	"context"
	"go-service-client/internal/adapter"
	"go-service-client/internal/commons"
	"go-service-client/internal/config"
	"go-service-client/internal/entities"
	"go-service-client/internal/usecases"
)

func main() {
	cfg := config.Load()

	client := commons.NewRestClient(cfg.Client.BaseURL, cfg.Client.Timeout)
	ctx := context.Background()

	mathAdapter := adapter.NewMathAdapter(adapter.Paths{
		Addition: adapter.PathConfig{
			Url:     "/api/v1/math/addition",
			Timeout: cfg.Timeout.Addition,
		},
		Multiply: adapter.PathConfig{
			Url:     "/api/v1/math/multiply",
			Timeout: cfg.Timeout.Multiply,
		},
		Subtraction: adapter.PathConfig{
			Url:     "/api/v1/math/subtraction",
			Timeout: cfg.Timeout.Subtraction,
		},
	}, client)

	u := usecases.NewMathUsecase(mathAdapter)
	u.FetchAllOperation(ctx, entities.Numbers{A: 3, B: 2})
}
