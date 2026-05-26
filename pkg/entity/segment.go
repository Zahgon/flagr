package entity

import (
	"github.com/zhouzhuojie/conditions"
	"gorm.io/gorm"
)

// SegmentDefaultRank is the default rank when we create the segment
const SegmentDefaultRank = uint(999)

// Segment is the unit of segmentation
type Segment struct {
	gorm.Model
	FlagID         uint   `gorm:"index:idx_segment_flagid"`
	Description    string `gorm:"type:text"`
	Rank           uint
	RolloutPercent uint
	Constraints    ConstraintArray
	Distributions  []Distribution

	// Purely for evaluation
	SegmentEvaluation SegmentEvaluation `gorm:"-" json:"-"`
}

// PreloadConstraintsDistribution preloads constraints and distributions
// for segment
func PreloadConstraintsDistribution(db *gorm.DB) *gorm.DB { _ = "STUB: not implemented"; return nil }

// Preload preloads the segment
func (s *Segment) Preload(db *gorm.DB) error { _ = "STUB: not implemented"; return nil }

// SegmentEvaluation is a struct that holds the necessary info for evaluation
type SegmentEvaluation struct {
	ConditionsExpr    conditions.Expr
	DistributionArray DistributionArray
	FlagIDStr         string // pre-formatted flagID string used as salt in rollout
}

// PrepareEvaluation prepares the segment for evaluation by parsing constraints
// and denormalize distributions
func (s *Segment) PrepareEvaluation() error { _ = "STUB: not implemented"; return nil }
