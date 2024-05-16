package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ApiResponse struct {
	Data string `json:"data"`
}

func CallExternalAPI(url string, method string, body *interface{}, token string) (*map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := &http.Client{}
	var buf bytes.Buffer

	// Jika ada parameter Body/body request untuk getnya
	err := json.NewEncoder(&buf).Encode(body)

	req, err := http.NewRequestWithContext(ctx, method, url, &buf)
	if token != "" {
		req.Header.Add("Authorization", token)
	}
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	defer client.CloseIdleConnections()

	// Convert struct to JSON
	respBody, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(respBody, &responseBody)

	return &responseBody, nil
}
