package commons

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type NumbersReq struct {
	A int32 `json:"a"`
	B int32 `json:"b"`
}

type NumbersResp struct {
	Error  string `json:"error,omitempty"`
	Result int32  `json:"result,omitempty"`
}

type RestClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewRestClient(baseUrl string, timeout time.Duration) *RestClient {
	return &RestClient{
		httpClient: &http.Client{
			Timeout: timeout, // Ini untuk timeout http client mulai dari buka koneksi TCP sampai pembacaan response body
		},
		baseURL: baseUrl,
	}
}

func (r *RestClient) PostJSON(ctx context.Context, path string, payload NumbersReq) (*NumbersResp, error) {
	var stream bytes.Buffer
	err := json.NewEncoder(&stream).Encode(payload)
	if err != nil {
		return nil, fmt.Errorf("gagal encode JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", r.baseURL+path, &stream)
	if err != nil {
		return nil, fmt.Errorf("gagal buat request: %w", err)
	}
	req.Header.Set("Content-type", "application/json")

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("gagal melakukan HTTP request: %w", err)
	}
	defer resp.Body.Close()

	var result NumbersResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("gagal decode response: %w", err)
	}

	return &result, nil
}
