package notification

import (
	"context"
	"net/http"
	"time"
)

func minDuration(a, b time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// doRequestWithRetry performs an HTTP request with exponential backoff and jitter.
// On success (status < 500), it returns (resp, nil).
// On failure after retries, it returns the last response (if any) and an error.
func doRequestWithRetry(ctx context.Context, client *http.Client, req *http.Request, maxRetries int, baseDelay, maxDelay time.Duration) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Don't close body here; caller will handle it if resp is returned

// Close any previous failed response body before returning success

// Success or client error (4xx) is considered final; no retry on 4xx

// 5xx - retryable
// Close previous lastResp.Body before overwriting to prevent resource leak

// Final attempt failed with 5xx - caller is responsible for closing body

// nextDelayWithJitter returns the next retry delay using exponential backoff
// with up to 50% jitter to prevent thundering herd.
func nextDelayWithJitter(prevDelay, maxDelay time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
