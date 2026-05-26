package handler

import (
	"github.com/openflagr/flagr/pkg/mapper/entity_restapi/e2r"
	"github.com/openflagr/flagr/pkg/mapper/entity_restapi/r2e"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/constraint"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/distribution"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/flag"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/segment"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/tag"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/variant"

	"github.com/go-openapi/runtime/middleware"
)

// CRUD is the CRUD interface
type CRUD interface {
	// Flags
	FindFlags(flag.FindFlagsParams) middleware.Responder
	CreateFlag(flag.CreateFlagParams) middleware.Responder
	GetFlag(flag.GetFlagParams) middleware.Responder
	PutFlag(flag.PutFlagParams) middleware.Responder
	DeleteFlag(flag.DeleteFlagParams) middleware.Responder
	RestoreFlag(flag.RestoreFlagParams) middleware.Responder
	SetFlagEnabledState(flag.SetFlagEnabledParams) middleware.Responder
	GetFlagSnapshots(params flag.GetFlagSnapshotsParams) middleware.Responder
	GetFlagEntityTypes(params flag.GetFlagEntityTypesParams) middleware.Responder

	//Tags
	CreateTag(tag.CreateTagParams) middleware.Responder
	DeleteTag(tag.DeleteTagParams) middleware.Responder
	FindTags(tag.FindTagsParams) middleware.Responder
	FindAllTags(params tag.FindAllTagsParams) middleware.Responder

	// Segments
	CreateSegment(segment.CreateSegmentParams) middleware.Responder
	FindSegments(segment.FindSegmentsParams) middleware.Responder
	PutSegment(segment.PutSegmentParams) middleware.Responder
	DeleteSegment(segment.DeleteSegmentParams) middleware.Responder
	PutSegmentsReorder(segment.PutSegmentsReorderParams) middleware.Responder

	// Constraints
	CreateConstraint(constraint.CreateConstraintParams) middleware.Responder
	FindConstraints(constraint.FindConstraintsParams) middleware.Responder
	PutConstraint(params constraint.PutConstraintParams) middleware.Responder
	DeleteConstraint(params constraint.DeleteConstraintParams) middleware.Responder

	// Distributions
	FindDistributions(distribution.FindDistributionsParams) middleware.Responder
	PutDistributions(distribution.PutDistributionsParams) middleware.Responder

	// Variants
	CreateVariant(variant.CreateVariantParams) middleware.Responder
	FindVariants(variant.FindVariantsParams) middleware.Responder
	PutVariant(variant.PutVariantParams) middleware.Responder
	DeleteVariant(variant.DeleteVariantParams) middleware.Responder
}

// NewCRUD creates a new CRUD instance
func NewCRUD() CRUD { _ = "STUB: not implemented"; return *new(CRUD) }

type crud struct{}

var (
	e2rMapFlag          = e2r.MapFlag
	e2rMapFlags         = e2r.MapFlags
	e2rMapFlagSnapshots = e2r.MapFlagSnapshots

	r2eMapAttachment    = r2e.MapAttachment
	r2eMapDistributions = r2e.MapDistributions
)

func (c *crud) FindFlags(params flag.FindFlagsParams) middleware.Responder {
	_ = "STUB: not implemented"
	// Add Unscoped so GORM doesn't automatically override `deleted_at`
	return *new(middleware.Responder)
}

// Always preload tags for searchability

func (c *crud) GetFlag(params flag.GetFlagParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

// Flag with given ID doesn't exist, so we 404

// Something else happened, return a 500

func (c *crud) GetFlagSnapshots(params flag.GetFlagSnapshotsParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) GetFlagEntityTypes(params flag.GetFlagEntityTypesParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) PutFlag(params flag.PutFlagParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) SetFlagEnabledState(params flag.SetFlagEnabledParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) RestoreFlag(params flag.RestoreFlagParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) DeleteFlag(params flag.DeleteFlagParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) DeleteTag(params tag.DeleteTagParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) FindTags(params tag.FindTagsParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) FindAllTags(params tag.FindAllTagsParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) CreateTag(params tag.CreateTagParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

// Find the existing tag to associate if it exists

func (c *crud) CreateSegment(params segment.CreateSegmentParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) FindSegments(params segment.FindSegmentsParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) PutSegment(params segment.PutSegmentParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) PutSegmentsReorder(params segment.PutSegmentsReorderParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) DeleteSegment(params segment.DeleteSegmentParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) CreateConstraint(params constraint.CreateConstraintParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) FindConstraints(params constraint.FindConstraintsParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) PutConstraint(params constraint.PutConstraintParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) DeleteConstraint(params constraint.DeleteConstraintParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

// PutDistributions puts the whole distributions and overwrite the old ones
func (c *crud) PutDistributions(params distribution.PutDistributionsParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) FindDistributions(params distribution.FindDistributionsParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) CreateVariant(params variant.CreateVariantParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) FindVariants(params variant.FindVariantsParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) PutVariant(params variant.PutVariantParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

func (c *crud) DeleteVariant(params variant.DeleteVariantParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}
