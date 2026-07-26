package usecases

import (
	"context"
	"fmt"
	"go-service-client/internal/entities"
)

type AllResponse struct {
	Addition    int32
	Multiply    int32
	Subtraction int32
}

type AllRespChannel struct {
	Name   string
	Result int32
}

type MathUsecase struct {
	gateway entities.MathGateway
}

func NewMathUsecase(gateway entities.MathGateway) *MathUsecase {
	return &MathUsecase{
		gateway: gateway,
	}
}

func (u *MathUsecase) FetchAllOperation(ctx context.Context, payload entities.Numbers) {
	respCh := make(chan AllRespChannel)
	go func() {
		res, _ := u.gateway.FetchAddition(ctx, payload)
		respCh <- AllRespChannel{Name: "addition", Result: res}
	}()

	go func() {
		res, _ := u.gateway.FetchMultiply(ctx, payload)
		respCh <- AllRespChannel{Name: "multiply", Result: res}
	}()

	go func() {
		res, _ := u.gateway.FetchSubtraction(ctx, payload)
		respCh <- AllRespChannel{Name: "subtraction", Result: res}
	}()

	var res AllResponse
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
