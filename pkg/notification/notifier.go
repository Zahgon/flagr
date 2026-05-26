package notification

import (
	"context"
	"sync"
	"time"
)

type Notifier interface {
	Send(ctx context.Context, n Notification) error
	Name() string
}

type Operation string

const (
	OperationCreate  Operation = "create"
	OperationUpdate  Operation = "update"
	OperationDelete  Operation = "delete"
	OperationRestore Operation = "restore"
)

// ComponentType identifies which part of a flag was modified.
type ComponentType string

const (
	ComponentFlag         ComponentType = "flag"
	ComponentSegment      ComponentType = "segment"
	ComponentVariant      ComponentType = "variant"
	ComponentConstraint   ComponentType = "constraint"
	ComponentDistribution ComponentType = "distribution"
	ComponentTag          ComponentType = "tag"
)

type Notification struct {
	Operation     Operation     `json:"operation"`
	FlagID        uint          `json:"flag_id"`
	FlagKey       string        `json:"flag_key"`
	ComponentType ComponentType `json:"component_type,omitempty"`
	ComponentID   uint          `json:"component_id,omitempty"`
	ComponentKey  string        `json:"component_key,omitempty"`
	PreValue      string        `json:"pre_value,omitempty"`
	PostValue     string        `json:"post_value,omitempty"`
	Diff          string        `json:"diff,omitempty"`
	User          string        `json:"user,omitempty"`
	Timestamp     time.Time     `json:"timestamp"`
}

var (
	// Notifiers is the list of configured notifiers. Set directly for testing.
	Notifiers []Notifier
	once      sync.Once
)

// GetNotifiers returns the list of configured notifiers.
// It initializes the notifiers on first call using sync.Once.
// For testing, set Notifiers directly before calling GetNotifiers.
func GetNotifiers() []Notifier {
	_ = "STUB: not implemented"
	// If already set (e.g., by tests), return immediately
	return nil
}

type nullNotifier struct{}

func (n *nullNotifier) Send(ctx context.Context, notification Notification) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *nullNotifier) Name() string { _ = "STUB: not implemented"; return "" }

type MockNotifier struct {
	sent      []Notification
	mu        sync.Mutex
	sendError error
}

func NewMockNotifier() *MockNotifier { _ = "STUB: not implemented"; return nil }

func (m *MockNotifier) Send(ctx context.Context, n Notification) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockNotifier) Name() string { _ = "STUB: not implemented"; return "" }

func (m *MockNotifier) SetSendError(err error) { _ = "STUB: not implemented"; return }

func (m *MockNotifier) GetSentNotifications() []Notification { _ = "STUB: not implemented"; return nil }

func (m *MockNotifier) ClearSent() { _ = "STUB: not implemented"; return }
