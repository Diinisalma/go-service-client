package commons

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-service-client/internal/entities"
	"io"
	"net/http"
	"time"
)

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

func (r *RestClient) PostJSON(ctx context.Context, path string, payload entities.NumbersReq) (*entities.NumbersResp, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("gagal encode JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", r.baseURL+path, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("gagal buat request: %w", err)
	}
	req.Header.Set("Content-type", "application/json")

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("gagal melakukan HTTP request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gagal baca response: %w", err)
	}
	defer resp.Body.Close()

	var result entities.NumbersResp
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("gagal decode response: %w", err)
	}

	return &result, nil
}
