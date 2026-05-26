package entity

import (
	"gorm.io/gorm"
)

// Flag is the unit of flags
type Flag struct {
	gorm.Model

	Key         string `gorm:"type:varchar(64);uniqueIndex:idx_flag_key"`
	Description string `gorm:"type:text"`
	CreatedBy   string
	UpdatedBy   string
	Enabled     bool
	Segments    []Segment
	Variants    []Variant
	Tags        []Tag `gorm:"many2many:flags_tags;"`
	SnapshotID  uint
	Notes       string `gorm:"type:text"`

	DataRecordsEnabled bool
	EntityType         string

	FlagEvaluation FlagEvaluation `gorm:"-" json:"-"`
}

// FlagEvaluation is a struct that holds the necessary info for evaluation
type FlagEvaluation struct {
	VariantsMap map[uint]*Variant
}

// Preloads just the tags
func PreloadFlagTags(db *gorm.DB) *gorm.DB { _ = "STUB: not implemented"; return nil }

// PreloadSegmentsVariantsTags preloads segments, variants and tags for flag
func PreloadSegmentsVariantsTags(db *gorm.DB) *gorm.DB { _ = "STUB: not implemented"; return nil }

// Preload preloads the segments, variants and tags into flags
func (f *Flag) Preload(db *gorm.DB) error { _ = "STUB: not implemented"; return nil }

// PreloadTags preloads the tags into flags
func (f *Flag) PreloadTags(db *gorm.DB) error { _ = "STUB: not implemented"; return nil }

// PrepareEvaluation prepares the information for evaluation
func (f *Flag) PrepareEvaluation() error { _ = "STUB: not implemented"; return nil }

// CreateFlagKey creates the key based on the given key
func CreateFlagKey(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// CreateFlagEntityType creates the FlagEntityType if not exists
func CreateFlagEntityType(db *gorm.DB, key string) error { _ = "STUB: not implemented"; return nil }
