package repositories

import (
	"context"
	"errors"
	"fmt"
	"go-service-client/internal/commons"
	"go-service-client/internal/entities"
	"time"
)

func callTimeout(ctx context.Context, c *commons.RestClient, path string, payload entities.NumbersReq, timeout time.Duration) (int32, error) {
	start := time.Now()

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	resp, err := c.PostJSON(ctx, path, payload)
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

func Addition(ctx context.Context, c *commons.RestClient, payload entities.NumbersReq) (int32, error) {
	res, err := callTimeout(ctx, c, "/math/addition", payload, 30*time.Second)
	return res, err
}

func Subtraction(ctx context.Context, c *commons.RestClient, payload entities.NumbersReq) (int32, error) {
	res, err := callTimeout(ctx, c, "/math/subtraction", payload, 30*time.Second)
	return res, err
}

func Multiply(ctx context.Context, c *commons.RestClient, payload entities.NumbersReq) (int32, error) {
	res, err := callTimeout(ctx, c, "/math/multiply", payload, 10*time.Second)
	return res, err
}
