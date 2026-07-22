package entities

import (
	"context"
)

type HttpMethod[T any] interface {
	PostJSON(ctx context.Context, path string, payload NumbersReq) (*NumbersResp, error)
}