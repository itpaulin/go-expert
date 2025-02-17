package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

	defer cancel()

	makeOrder(ctx)
}

func makeOrder(ctx context.Context) {

	select {
	case <-ctx.Done():
		println("Timeout, order canceled!")
		return
	case <-time.After(5 * time.Second):
		println("Order completed!")
		return
	}

}
