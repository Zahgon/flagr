package handler

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/openflagr/flagr/pkg/config"
	"github.com/openflagr/flagr/pkg/util"
	"github.com/openflagr/flagr/swagger_gen/models"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

var (
	saramaNewAsyncProducer = sarama.NewAsyncProducer
)

func mustParseKafkaVersion(version string) sarama.KafkaVersion {
	_ = "STUB: not implemented"
	return *new(sarama.KafkaVersion)
}

// NewKafkaRecorder creates a new Kafka recorder
var NewKafkaRecorder = func() DataRecorder {
	cfg := sarama.NewConfig()

	tlscfg := createTLSConfiguration(
		config.Config.RecorderKafkaCertFile,
		config.Config.RecorderKafkaKeyFile,
		config.Config.RecorderKafkaCAFile,
		config.Config.RecorderKafkaVerifySSL,
		config.Config.RecorderKafkaSimpleSSL,
	)
	if tlscfg != nil {
		cfg.Net.TLS.Enable = true
		cfg.Net.TLS.Config = tlscfg
	}

	if config.Config.RecorderKafkaSASLUsername != "" && config.Config.RecorderKafkaSASLPassword != "" {
		cfg.Net.SASL.Enable = true
		cfg.Net.SASL.User = config.Config.RecorderKafkaSASLUsername
		cfg.Net.SASL.Password = config.Config.RecorderKafkaSASLPassword
	}

	cfg.Net.MaxOpenRequests = config.Config.RecorderKafkaMaxOpenReqs

	cfg.Producer.Compression = sarama.CompressionCodec(config.Config.RecorderKafkaCompressionCodec)
	cfg.Producer.RequiredAcks = sarama.RequiredAcks(config.Config.RecorderKafkaRequiredAcks)
	cfg.Producer.Idempotent = config.Config.RecorderKafkaIdempotent
	cfg.Producer.Retry.Max = config.Config.RecorderKafkaRetryMax
	cfg.Producer.Flush.Frequency = config.Config.RecorderKafkaFlushFrequency
	cfg.Version = mustParseKafkaVersion(config.Config.RecorderKafkaVersion)

	brokerList := strings.Split(config.Config.RecorderKafkaBrokers, ",")
	producer, err := saramaNewAsyncProducer(brokerList, cfg)
	if err != nil {
		logrus.WithField("kafka_error", err).Fatal("Failed to start Sarama producer:")
	}

	// We will just log to STDOUT if we're not able to produce messages.
	if producer != nil {
		go func() {
			for err := range producer.Errors() {
				logrus.WithField("kafka_error", err).Error("failed to write access log entry")
			}
		}()
	}

	var encryptor dataRecordEncryptor
	if config.Config.RecorderKafkaEncrypted && config.Config.RecorderKafkaEncryptionKey != "" {
		encryptor = newSimpleboxEncryptor(config.Config.RecorderKafkaEncryptionKey)
	}

	return &kafkaRecorder{
		topic:               config.Config.RecorderKafkaTopic,
		partitionKeyEnabled: config.Config.RecorderKafkaPartitionKeyEnabled,
		producer:            producer,
		options: DataRecordFrameOptions{
			Encrypted:       config.Config.RecorderKafkaEncrypted,
			Encryptor:       encryptor,
			FrameOutputMode: config.Config.RecorderFrameOutputMode,
		},
	}
}

func createTLSConfiguration(certFile string, keyFile string, caFile string, verifySSL bool, simpleSSL bool) (t *tls.Config) {
	_ = "STUB: not implemented"
	return nil
}

// will be nil by default if nothing is provided

type kafkaRecorder struct {
	producer            sarama.AsyncProducer
	topic               string
	options             DataRecordFrameOptions
	partitionKeyEnabled bool
}

func (k *kafkaRecorder) NewDataRecordFrame(r models.EvalResult) DataRecordFrame {
	_ = "STUB: not implemented"
	return *new(DataRecordFrame)
}

func (k *kafkaRecorder) AsyncRecord(r models.EvalResult) { _ = "STUB: not implemented"; return }

var logKafkaAsyncRecordToDatadog = func(r models.EvalResult) {
	if config.Global.StatsdClient == nil {
		return
	}
	config.Global.StatsdClient.Incr(
		"data_recorder.kafka",
		[]string{
			fmt.Sprintf("FlagID:%d", util.SafeUint(r.FlagID)),
		},
		float64(1),
	)
}
