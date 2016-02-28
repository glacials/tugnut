package main

import (
	"testing"

	"golang.org/x/net/context"
)

func TestServe(t *testing.T) {
	ctx := context.Background()

	_ = buildMux(ctx)
}
