package main

import (
	"context"
	"fmt"
	"go-service-client/internal/commons"
	"go-service-client/internal/config"
	"go-service-client/internal/entities"
	"go-service-client/internal/repositories"
)

type allResponse struct {
	Addition    int32
	Multiply    int32
	Subtraction int32
}

type allRespChannel struct {
	Name   string
	Result int32
}

func main() {
	cfg := config.Load()

	client := commons.NewRestClient(cfg.Client.BaseURL, cfg.Client.Timeout)
	payload := entities.NumbersReq{A: 3, B: 2}
	ctx := context.Background()
	respCh := make(chan allRespChannel)

	go func() {
		res, _ := repositories.Addition(ctx, client, payload, cfg.Timeout.Addition)
		respCh <- allRespChannel{Name: "addition", Result: res}
	}()

	go func() {
		res, _ := repositories.Multiply(ctx, client, payload, cfg.Timeout.Multiply)
		respCh <- allRespChannel{Name: "multiply", Result: res}
	}()

	go func() {
		res, _ := repositories.Subtraction(ctx, client, payload, cfg.Timeout.Subtraction)
		respCh <- allRespChannel{Name: "subtraction", Result: res}
	}()

	var res allResponse
	for i := 0; i < 3; i++ {
		r := <-respCh
		switch r.Name {
		case "addition":
			res.Addition = r.Result
		case "multiply":
			res.Multiply = r.Result
		case "subtraction":
			res.Subtraction = r.Result
		}
	}
	fmt.Println(res)
}
