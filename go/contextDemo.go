package main

import (
	"context"
	"fmt"
	"time"
)

func goroutine01(ctx context.Context) {
	for {
		select {
		case _, _ = <-ctx.Done():
			fmt.Println("goroutine01 is done")
			return
		}
	}
}

func goroutine02(ctx context.Context) {
	newCtx, _ := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	go goroutine0201(newCtx)
	//time.Sleep(time.Second)
	//cancel()
	for {
		select {
		case _, _ = <-ctx.Done():
			fmt.Println("goroutine02 is done")
			return
		}
	}
}

func goroutine0201(ctx context.Context) {
	if timeOut, ok := ctx.Deadline(); ok {
		fmt.Println("goroutine0201", timeOut)
	}
	go goroutine020101(ctx)
	for {
		select {
		case _, _ = <-ctx.Done():
			fmt.Println("goroutine0201 is done")
			return
		}

	}
}
func goroutine020101(ctx context.Context) {
	if timeOut, ok := ctx.Deadline(); ok {
		fmt.Println("goroutine020101", timeOut)
	}
	for {
		select {
		case _, _ = <-ctx.Done():
			fmt.Println("goroutine020101 is done")
			return
		}
	}
}

func goroutine03(ctx context.Context) {
	for {
		select {
		case _, _ = <-ctx.Done():
			fmt.Println("goroutine03 is done")
			return
		}
	}
}

func main() {
	ctx := context.Background()
	go goroutine01(ctx)
	go goroutine02(ctx)
	go goroutine03(ctx)
	for {

	}
}
