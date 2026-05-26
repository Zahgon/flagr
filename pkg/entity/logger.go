package entity

import (
	"context"
	"time"

	gorm_logger "gorm.io/gorm/logger"
)

type Logger struct {
	LogLevel                  gorm_logger.LogLevel
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
}

func (l *Logger) LogMode(level gorm_logger.LogLevel) gorm_logger.Interface {
	_ = "STUB: not implemented"
	return *new(gorm_logger.Interface)
}

func (l *Logger) Info(ctx context.Context, s string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (l *Logger) Warn(ctx context.Context, s string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (l *Logger) Error(ctx context.Context, s string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	_ = "STUB: not implemented"
	return
}
