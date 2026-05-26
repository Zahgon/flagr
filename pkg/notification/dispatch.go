package notification

var (
	// Semaphore to limit concurrent notification sends. Default 100.
	notificationSemaphore = make(chan struct{}, 100)
)

func recordNotificationMetrics(provider string, operation Operation, success bool) {
	_ = "STUB: not implemented"
	return
}

// SendNotification dispatches a notification to all configured notifiers asynchronously.
// Notifications are sent in a background goroutine and failures do not affect the caller.
func SendNotification(n Notification) {
	_ = "STUB: not implemented"
	// Capture notifiers BEFORE spawning goroutine to avoid test pollution
	// when Notifiers is modified between test runs
	return
}

// Set timestamp if not already set by caller

// Acquire semaphore slot

// Send to all notifiers concurrently, aggregate errors

func CalculateDiff(pre, post string) string { _ = "STUB: not implemented"; return "" }

func prettyPrintJSON(s string) string { _ = "STUB: not implemented"; return "" }
