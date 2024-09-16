package service

import (
	"context"
	"sync"
)

// FetchResult is a generic type for the result of a fetch operation
type fetchResult[T any] struct {
	Data T
	Err  error
}

// ConcurrentFetch is a utility function that wraps the concurrency logic for a single fetch operation
func concurrentFetch[T any](
	ctx context.Context,
	wg *sync.WaitGroup,

	errCh chan<- error,
	fetchFunc func(context.Context) (T, error),
) fetchResult[T] {
	var result fetchResult[T]
	wg.Add(1)
	go func() {
		defer wg.Done()
		data, err := fetchFunc(ctx)
		if err != nil {
			errCh <- err

			return
		}

		result.Data = data
	}()

	return result
}
