package grpc

import (
	"context"
	"go-service-client/internal/entities"

	mathpb "github.com/Diinisalma/mini_project-service-math/api/v1"
	"google.golang.org/grpc"
)

type MathGrpcAdapter struct {
	client mathpb.MathServiceClient
}

func NewMathGrpcAdapter(conn *grpc.ClientConn) *MathGrpcAdapter {
	return &MathGrpcAdapter{
		client: mathpb.NewMathServiceClient(conn),
	}
}

var _ entities.MathGateway = &MathGrpcAdapter{}

func (m *MathGrpcAdapter) FetchAddition(ctx context.Context, payload entities.Numbers) (int32, error) {
	res, err := m.client.Addition(ctx, &mathpb.AdditionReq{A: payload.A, B: payload.B})
	if err != nil {
		return 0, err
	}
	return res.Result, nil
}

func (m *MathGrpcAdapter) FetchMultiply(ctx context.Context, payload entities.Numbers) (int32, error) {
	res, err := m.client.Multiply(ctx, &mathpb.MultiplyReq{A: payload.A, B: payload.B})
	if err != nil {
		return 0, err
	}
	return res.Result, nil
}

func (m *MathGrpcAdapter) FetchSubtraction(ctx context.Context, payload entities.Numbers) (int32, error) {
	res, err := m.client.Subtraction(ctx, &mathpb.SubtractionReq{A: payload.A, B: payload.B})
	if err != nil {
		return 0, err
	}
	return res.Result, nil
}
