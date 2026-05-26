package entity

import (
	"github.com/openflagr/flagr/swagger_gen/models"
	"github.com/zhouzhuojie/conditions"
	"gorm.io/gorm"
)

// Constraint is the unit of constraints
type Constraint struct {
	gorm.Model

	SegmentID uint `gorm:"index:idx_constraint_segmentid"`
	Property  string
	Operator  string
	Value     string `gorm:"type:text"`
}

// ConstraintArray is an array of Constraint
type ConstraintArray []Constraint

// OperatorToExprMap maps from the swagger model operator to condition operator
var OperatorToExprMap = map[string]string{
	models.ConstraintOperatorEQ:          "==",
	models.ConstraintOperatorNEQ:         "!=",
	models.ConstraintOperatorLT:          "<",
	models.ConstraintOperatorLTE:         "<=",
	models.ConstraintOperatorGT:          ">",
	models.ConstraintOperatorGTE:         ">=",
	models.ConstraintOperatorEREG:        "=~",
	models.ConstraintOperatorNEREG:       "!~",
	models.ConstraintOperatorIN:          "IN",
	models.ConstraintOperatorNOTIN:       "NOT IN",
	models.ConstraintOperatorCONTAINS:    "CONTAINS",
	models.ConstraintOperatorNOTCONTAINS: "NOT CONTAINS",
}

// ToExpr transfer the constraint to conditions.Expr for evaluation
func (c *Constraint) ToExpr() (conditions.Expr, error) {
	_ = "STUB: not implemented"
	return *new(conditions.Expr), nil
}

func (c *Constraint) toExprStr() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Trim the value to be resilient against untrimmed values from API callers.

// For EREG/NEREG with quoted string values, use regex literal form /pattern/
// to avoid Go text/scanner escape issues with sequences like \d, \., \s, etc.
// The scanner interprets escape sequences inside quoted strings and rejects
// unrecognized ones (like \., \d), but regex literals are read character-by-character
// without escape processing, so patterns with backslashes work correctly.

// Only use regex literal form when the pattern doesn't contain "/",
// because the conditions parser doesn't support escaping "/" inside //.

// isQuotedString reports whether s is a double-quoted string like "foo".
func isQuotedString(s string) bool { _ = "STUB: not implemented"; return false }

// Validate validates Constraint
func (c *Constraint) Validate() error { _ = "STUB: not implemented"; return nil }

// ToExpr maps ConstraintArray to expr by joining 'AND'
func (cs ConstraintArray) ToExpr() (conditions.Expr, error) {
	_ = "STUB: not implemented"
	return *new(conditions.Expr), nil
}
