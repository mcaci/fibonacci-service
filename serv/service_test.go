package serv

import (
	"context"
	"testing"
)

func setup() (srv Service, ctx context.Context) {
	return NewFibonacciService(), context.Background()
}

func TestFibonacci(t *testing.T) {
	srv, ctx := setup()

	serie, err := srv.Fibonacci(ctx, 3)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// testing status
	if ok := len(serie) == 3; !ok {
		t.Errorf("expected service to be ok")
	}
}
