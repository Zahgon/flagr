package handler

import (
	"encoding/json"

	"github.com/brandur/simplebox"
	"github.com/openflagr/flagr/swagger_gen/models"
)

type dataRecordEncryptor interface {
	Encrypt([]byte) (string, error)
}

type simpleboxEncryptor struct{ key [simplebox.KeySize]byte }

func (se *simpleboxEncryptor) Encrypt(b []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func newSimpleboxEncryptor(k string) dataRecordEncryptor {
	_ = "STUB: not implemented"
	return *new(dataRecordEncryptor)
}

const (
	frameOutputModePayloadRawJSON = "payload_raw_json"
)

// DataRecordFrameOptions represents the options we can set to create a DataRecordFrame
type DataRecordFrameOptions struct {
	Encrypted       bool
	Encryptor       dataRecordEncryptor
	FrameOutputMode string
}

type rawPayload struct {
	Payload json.RawMessage `json:"payload"`
}

type stringPayload struct {
	Payload   string `json:"payload"`
	Encrypted bool   `json:"encrypted"`
}

// DataRecordFrame represents the structure we can json.Marshal into data recorders
type DataRecordFrame struct {
	evalResult models.EvalResult
	options    DataRecordFrameOptions
}

// MarshalJSON defines the behavior of MarshalJSON for DataRecordFrame
func (drf *DataRecordFrame) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPartitionKey gets the partition key from entityID
func (drf *DataRecordFrame) GetPartitionKey() string { _ = "STUB: not implemented"; return "" }

// Output sets the payload using its input and returns the json marshal bytes
func (drf *DataRecordFrame) Output() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
