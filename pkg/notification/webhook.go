package notification

import (
	"context"
	"net/http"
)

type webhookNotifier struct {
	httpClient *http.Client
}

// maxErrorBodyBytes limits how much of a webhook error response body is read
// to prevent memory exhaustion from a misconfigured or malicious endpoint.
const maxErrorBodyBytes = 4096

func NewWebhookNotifier() Notifier { _ = "STUB: not implemented"; return *new(Notifier) }

func (w *webhookNotifier) Send(ctx context.Context, n Notification) error {
	_ = "STUB: not implemented"
	return nil
}

// Execute request with retry

func (w *webhookNotifier) Name() string { _ = "STUB: not implemented"; return "" }
