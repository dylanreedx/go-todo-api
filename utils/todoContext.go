package utils

import (
	"context"
	"time"
)

// Create a context that will be cancelled after 10 seconds.
func TodoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
