package serv

import "context"

// Service provides and interface for adding integers
type Service interface {
	Fibonacci(ctx context.Context, n int) ([]uint64, error)
}
