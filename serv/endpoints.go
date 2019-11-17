package serv

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	FibonacciEndpoint endpoint.Endpoint
}

// MakeFibonacciEndpoint returns the response from our service "Fibonacci"
func MakeFibonacciEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(fibonacciRequest)
		b, err := srv.Fibonacci(ctx, req.N)
		if err != nil {
			return fibonacciResponse{b, err.Error()}, nil
		}
		return fibonacciResponse{b, ""}, nil
	}
}

// Fibonacci endpoint mapping
func (e Endpoints) Fibonacci(ctx context.Context, n int) ([]uint64, error) {
	req := fibonacciRequest{N: n}
	resp, err := e.FibonacciEndpoint(ctx, req)
	if err != nil {
		return []uint64{}, err
	}
	FibonacciResp := resp.(fibonacciResponse)
	if FibonacciResp.Err != "" {
		return []uint64{}, errors.New(FibonacciResp.Err)
	}
	return FibonacciResp.Serie, nil
}
