package handler

import (
	"sync"

	"github.com/openflagr/flagr/swagger_gen/models"
)

var (
	singletonDataRecorder     DataRecorder
	singletonDataRecorderOnce sync.Once
)

// DataRecorder can record and produce the evaluation result
type DataRecorder interface {
	AsyncRecord(models.EvalResult)
	NewDataRecordFrame(models.EvalResult) DataRecordFrame
}

// GetDataRecorder gets the data recorder
func GetDataRecorder() DataRecorder { _ = "STUB: not implemented"; return *new(DataRecorder) }
