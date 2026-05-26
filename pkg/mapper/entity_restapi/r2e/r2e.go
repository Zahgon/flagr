package r2e

import (
	"github.com/openflagr/flagr/pkg/entity"
	"github.com/openflagr/flagr/swagger_gen/models"
)

// MapDistributions maps distribution
func MapDistributions(r []*models.Distribution, segmentID uint) []entity.Distribution {
	_ = "STUB: not implemented"
	return nil
}

// MapDistribution maps distribution
func MapDistribution(r *models.Distribution, segmentID uint) entity.Distribution {
	_ = "STUB: not implemented"
	return *new(entity.Distribution)
}

// MapAttachment maps attachment
func MapAttachment(a any) (entity.Attachment, error) {
	_ = "STUB: not implemented"
	return *new(entity.Attachment), nil
}
