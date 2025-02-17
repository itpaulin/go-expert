package main

import (
	"context"
	"fmt"
)

type contextKey string

func main() {
	key := contextKey("paulo")
	ctx := context.WithValue(context.Background(), key, "value")
	makeOrder(ctx)
}

func makeOrder(ctx context.Context) {
	key := contextKey("paulo")
	value := ctx.Value(key)
	fmt.Println(value)

}
