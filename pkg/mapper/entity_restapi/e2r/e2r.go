package e2r

import (
	"github.com/openflagr/flagr/pkg/entity"

	"github.com/openflagr/flagr/swagger_gen/models"
)

// MapFlag maps flag
func MapFlag(e *entity.Flag) (*models.Flag, error) { _ = "STUB: not implemented"; return nil, nil }

// MapFlags maps flags
func MapFlags(e []entity.Flag) ([]*models.Flag, error) { _ = "STUB: not implemented"; return nil, nil }

// MapFlagSnapshot maps flag snapshot
func MapFlagSnapshot(e *entity.FlagSnapshot) (*models.FlagSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MapFlagSnapshots maps flag snapshots
func MapFlagSnapshots(e []entity.FlagSnapshot) ([]*models.FlagSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MapSegment maps segment
func MapSegment(e *entity.Segment) *models.Segment { _ = "STUB: not implemented"; return nil }

// MapSegments maps segments
func MapSegments(e []entity.Segment) []*models.Segment { _ = "STUB: not implemented"; return nil }

// MapTagEntity maps tag entity
func MapTag(e *entity.Tag) *models.Tag { _ = "STUB: not implemented"; return nil }

// MapTags maps tags
func MapTags(e []entity.Tag) []*models.Tag { _ = "STUB: not implemented"; return nil }

// MapConstraint maps constraint
func MapConstraint(e *entity.Constraint) *models.Constraint { _ = "STUB: not implemented"; return nil }

// MapConstraints maps constraints
func MapConstraints(e []entity.Constraint) []*models.Constraint {
	_ = "STUB: not implemented"
	return nil
}

// MapDistribution maps to a distribution
func MapDistribution(e *entity.Distribution) *models.Distribution {
	_ = "STUB: not implemented"
	return nil
}

// MapDistributions maps distribution
func MapDistributions(e []entity.Distribution) []*models.Distribution {
	_ = "STUB: not implemented"
	return nil
}

// MapVariant maps variant
func MapVariant(e *entity.Variant) *models.Variant { _ = "STUB: not implemented"; return nil }

// MapVariants maps variant
func MapVariants(e []entity.Variant) []*models.Variant { _ = "STUB: not implemented"; return nil }
