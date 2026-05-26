package entity

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

// Variant is the struct that represent the experience/variant of the evaluation entity
type Variant struct {
	gorm.Model
	FlagID     uint `gorm:"index:idx_variant_flagid"`
	Key        string
	Attachment Attachment `gorm:"type:text"`
}

// Validate validates the Variant
func (v *Variant) Validate() error { _ = "STUB: not implemented"; return nil }

// Attachment supports dynamic configuration in variant
type Attachment map[string]any

// Scan implements scanner interface
func (a *Attachment) Scan(value any) error { _ = "STUB: not implemented"; return nil }

// Value implements valuer interface
func (a Attachment) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Get returns the value of the key
func (a Attachment) Get(key string) any {
	_ = "STUB: not implemented"

	// GetString returns the string value of the key
	return *new(any)
}

func (a Attachment) GetString(key string) string { _ = "STUB: not implemented"; return "" }

// GetInt returns the int value of the key
func (a Attachment) GetInt(key string) int { _ = "STUB: not implemented"; return 0 }

// GetBool returns the bool value of the key
func (a Attachment) GetBool(key string) bool { _ = "STUB: not implemented"; return false }

// GetFloat64 returns the float64 value of the key
func (a Attachment) GetFloat64(key string) float64 { _ = "STUB: not implemented"; return 0 }
