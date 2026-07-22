package main

import (
	"context"
	"fmt"
	"go-service-client/internal/commons"
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
	client := commons.NewRestClient("http://localhost:1323")
	payload := entities.NumbersReq{A: 3, B: 2}
	ctx := context.Background()
	respCh := make(chan allRespChannel)

	go func() {
		res, _ := repositories.Addition(ctx, client, payload)
		respCh <- allRespChannel{Name: "addition", Result: res}
	}()

	go func() {
		res, _ := repositories.Multiply(ctx, client, payload)
		respCh <- allRespChannel{Name: "multiply", Result: res}
	}()

	go func() {
		res, _ := repositories.Subtraction(ctx, client, payload)
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
