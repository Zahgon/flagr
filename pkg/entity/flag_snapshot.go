package entity

import (
	"fmt"

	"github.com/openflagr/flagr/pkg/config"
	"github.com/openflagr/flagr/pkg/notification"
	"github.com/openflagr/flagr/pkg/util"
	"gorm.io/gorm"
)

// FlagSnapshot is the snapshot of a flag
// Any change of the flag will create a new snapshot
type FlagSnapshot struct {
	gorm.Model
	FlagID    uint `gorm:"index:idx_flagsnapshot_flagid"`
	UpdatedBy string
	Flag      []byte `gorm:"type:text"`
}

// SaveFlagSnapshot saves the Flag Snapshot and sends a notification.
func SaveFlagSnapshot(db *gorm.DB, flagID uint, updatedBy string, operation notification.Operation, componentType notification.ComponentType, componentID uint, componentKey string) {
	_ = "STUB: not implemented"
	return
}

// Use Unscoped to include soft-deleted flags. This is necessary for:
// 1. Delete operations: we need to snapshot the flag after it's been soft-deleted
// 2. Restore operations: we need to update the flag that was previously soft-deleted
// This is safe because flagID comes from validated request params and the operation
// is explicitly tracked (create/update/delete/restore).

// Use Unscoped to update soft-deleted flags (e.g., after delete operation).
// Without Unscoped(), GORM would add "deleted_at IS NULL" condition and fail.

// Find the most recent snapshot before the current one (use Unscoped to include any soft-deleted).
// ErrRecordNotFound is expected for the first snapshot of a flag.

var logFlagSnapshotUpdate = func(flagID uint, updatedBy string) {
	if config.Global.StatsdClient == nil {
		return
	}

	config.Global.StatsdClient.Incr(
		"flag.snapshot.updated",
		[]string{
			fmt.Sprintf("FlagID:%d", flagID),
			fmt.Sprintf("UpdatedBy:%s", util.SafeStringWithDefault(updatedBy, "null")),
		},
		float64(1),
	)
}
