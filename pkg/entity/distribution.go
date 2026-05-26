package entity

import (
	"gorm.io/gorm"
)

const (
	// TotalBucketNum represents how many buckets we can use to determine the consistent hashing
	// distribution and rollout
	TotalBucketNum uint = 1000

	// PercentMultiplier implies that the multiplier between percentage (100) and TotalBucketNum
	PercentMultiplier uint = TotalBucketNum / uint(100)
)

// Distribution is the struct represents distribution under segment and links to variant
type Distribution struct {
	gorm.Model
	SegmentID  uint `gorm:"index:idx_distribution_segmentid"`
	VariantID  uint `gorm:"index:idx_distribution_variantid"`
	VariantKey string

	Percent uint   // Percent is an uint from 0 to 100, percent is always derived from Bitmap
	Bitmap  string `gorm:"type:text" json:"-"`
}

// DistributionArray is useful for faster evaluation
type DistributionArray struct {
	VariantIDs          []uint
	PercentsAccumulated []int // useful for binary search to find the rollout variant
}

// DistributionDebugLog is useful for making debug logs
type DistributionDebugLog struct {
	BucketNum         uint
	DistributionArray DistributionArray
	VariantID         uint
	RolloutPercent    uint
}

// Rollout rolls out the entity based on the rolloutPercent
func (d DistributionArray) Rollout(entityID string, salt string, rolloutPercent uint) (variantID *uint, msg string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func (d DistributionArray) bucketByNum(bucketNum uint) (variantID uint, index int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (d DistributionArray) rollout(bucketNum uint, rolloutPercent uint, index int) bool {
	_ = "STUB: not implemented"
	return false
}

func crc32Num(entityID string, salt string) uint {
	_ = "STUB: not implemented"
	// crc32 is good in terms of uniform distribution
	// http://michiel.buddingh.eu/distribution-of-hash-values
	return 0
}
