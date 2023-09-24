package main

import (
	"context"
	"fmt"
	"time"
)

func pupolateContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "id", "1234")
}

func showContext(ctx context.Context) {
	val := ctx.Value("id")
	fmt.Println(val)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}

}
func main() {
	//fmt.Println("--- Go context ---")
	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()
	//ctx = pupolateContext(ctx)
	//go showContext(ctx)
	//select {
	//case <-ctx.Done():
	//	fmt.Println("Oh no, I have exceeded the deadline")
	//}
	//time.Sleep(2 * time.Second)

	// Create a context with a 2-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure the context is canceled when done.

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Work done")
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			// The context was canceled due to a timeout.
			fmt.Println("Operation canceled due to timeout")
		} else if ctx.Err() == context.Canceled {
			// The context was canceled for some other reason.
			fmt.Println("Operation canceled for other reason")
		} else {
			// The context finished without error.
			fmt.Println("Operation completed successfully")
		}
	}

}
