package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()                            // Context Created
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second) // Define Timeout for context
	defer cancel()                                         // Cancel the context in the end of operation
	bookHotel(ctx)                                         // Call the function that return the context status
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Print("Hotel booking cancelled. Timeout reached")
		return
	case <-time.After(2 * time.Second):
		fmt.Print("Hotel booked!")
	}
}
