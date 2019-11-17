package serv

import (
	"context"
	"encoding/json"
	"net/http"
)

// In the first part of the file we are mapping requests and responses to their JSON payload.
type fibonacciRequest struct {
	N int `json:"n"`
}

type fibonacciResponse struct {
	Serie []uint64 `json:"serie"`
	Err   string   `json:"err,omitempty"`
}

// In the second part we will write "decoders" for our incoming requests
func decodeFibonacciRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req fibonacciRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Last but not least, we have the encoder for the response output
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}