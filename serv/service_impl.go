package serv

import (
	"context"

	"github.com/mcaci/fibonacci/fibonacci"
)

type fibonacciService struct{}

// NewFibonacciService makes a new Service
func NewFibonacciService() Service {
	return fibonacciService{}
}

func (fibonacciService) Fibonacci(ctx context.Context, n int) ([]uint64, error) {
	return fibonacci.Serie(n), nil
}
