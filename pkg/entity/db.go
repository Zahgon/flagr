package entity

import (
	"sync"

	// sqlite driver with pure go
	// mysql driver
	// postgres driver

	"gorm.io/gorm"
)

var (
	singletonDB   *gorm.DB
	singletonOnce sync.Once
)

// AutoMigrateTables stores the entity tables that we can auto migrate in gorm
var AutoMigrateTables = []any{
	Flag{},
	Constraint{},
	Distribution{},
	FlagSnapshot{},
	Segment{},
	User{},
	Variant{},
	Tag{},
	FlagEntityType{},
}

func connectDB() (db *gorm.DB, err error) { _ = "STUB: not implemented"; return nil, nil }

// GetDB gets the db singleton
func GetDB() *gorm.DB { _ = "STUB: not implemented"; return nil }

// NewSQLiteDB creates a new sqlite db
// useful for backup exports and unit tests
func NewSQLiteDB(filePath string) *gorm.DB { _ = "STUB: not implemented"; return nil }

// NewTestDB creates a new test db
func NewTestDB() *gorm.DB { _ = "STUB: not implemented"; return nil }

// PopulateTestDB seeds the test db
func PopulateTestDB(flag Flag) *gorm.DB { _ = "STUB: not implemented"; return nil }
