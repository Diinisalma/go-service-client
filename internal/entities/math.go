package entities

import "context"

type Numbers struct {
	A int32
	B int32
}

type MathGateway interface {
	FetchMultiply(ctx context.Context, numbers Numbers) (int32, error)
	FetchSubtraction(ctx context.Context, numbers Numbers) (int32, error)
	FetchAddition(ctx context.Context, numbers Numbers) (int32, error)
}
