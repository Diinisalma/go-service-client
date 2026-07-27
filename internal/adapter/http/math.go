package adapter

import (
	"context"
	"errors"
	"fmt"
	"go-service-client/internal/commons"
	"go-service-client/internal/entities"
	"time"
)

type httpClient interface {
	PostJSON(ctx context.Context, path string, payload commons.NumbersReq) (*commons.NumbersResp, error)
}

type PathConfig struct {
	Url     string
	Timeout time.Duration
}

type Paths struct {
	Addition    PathConfig
	Multiply    PathConfig
	Subtraction PathConfig
}

type MathAdapter struct {
	Paths  Paths
	client httpClient
}

var _ entities.MathGateway = &MathAdapter{}

func NewMathAdapter(paths Paths, client httpClient) *MathAdapter {
	return &MathAdapter{
		Paths:  paths,
		client: client,
	}
}

func callTimeout(ctx context.Context, c httpClient, path string, payload entities.Numbers, timeout time.Duration) (int32, error) {
	start := time.Now()

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	resp, err := c.PostJSON(ctx, path, commons.NumbersReq{
		A: payload.A,
		B: payload.B,
	})
	elapsed := time.Since(start)
	if err != nil {
		fmt.Printf("[%v] %s → ERROR setelah %v: %v \n", timeout, path, elapsed, err)
		return 0, err
	}

	if resp.Error != "" {
		fmt.Printf("[%v] %s → server error %v: %v \n", timeout, path, elapsed, err)
		return 0, errors.New(resp.Error)
	}

	fmt.Printf("[%v] %s → %v (result: %d) \n", timeout, path, elapsed, resp.Result)
	return resp.Result, nil
}

func (m *MathAdapter) FetchAddition(ctx context.Context, payload entities.Numbers) (int32, error) {
	res, err := callTimeout(ctx, m.client, m.Paths.Addition.Url, payload, m.Paths.Addition.Timeout)
	return res, err
}

func (m *MathAdapter) FetchSubtraction(ctx context.Context, payload entities.Numbers) (int32, error) {
	res, err := callTimeout(ctx, m.client, m.Paths.Subtraction.Url, payload, m.Paths.Subtraction.Timeout)
	return res, err
}

func (m *MathAdapter) FetchMultiply(ctx context.Context, payload entities.Numbers) (int32, error) {
	res, err := callTimeout(ctx, m.client, m.Paths.Multiply.Url, payload, m.Paths.Multiply.Timeout)
	return res, err
}
