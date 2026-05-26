package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/openflagr/flagr/pkg/entity"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/flag"
	"gorm.io/gorm"
)

func (c *crud) CreateFlag(params flag.CreateFlagParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

// LoadSimpleBooleanFlagTemplate loads the simple boolean flag template into
// a new flag. It creates a single segment, variant ('on'), and distribution.
func LoadSimpleBooleanFlagTemplate(flag *entity.Flag, tx *gorm.DB) error {
	_ = "STUB: not implemented"
	// Create our default segment
	return nil
}

// .. and our default Variant

// .. and our default Distribution
