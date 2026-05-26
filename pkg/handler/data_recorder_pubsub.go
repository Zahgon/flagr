package handler

import (
	"context"

	"cloud.google.com/go/pubsub/v2"
	"github.com/openflagr/flagr/pkg/config"
	"github.com/openflagr/flagr/swagger_gen/models"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

type pubsubRecorder struct {
	producer  *pubsub.Client
	publisher *pubsub.Publisher
	options   DataRecordFrameOptions
}

var (
	pubsubClient = func() (*pubsub.Client, error) {
		return pubsub.NewClient(
			context.Background(),
			config.Config.RecorderPubsubProjectID,
			option.WithCredentialsFile(config.Config.RecorderPubsubKeyFile),
		)
	}
)

// NewPubsubRecorder creates a new Pubsub recorder
var NewPubsubRecorder = func() DataRecorder {
	client, err := pubsubClient()
	if err != nil {
		logrus.WithField("pubsub_error", err).Fatal("error getting pubsub client")
	}

	return &pubsubRecorder{
		producer:  client,
		publisher: client.Publisher(config.Config.RecorderPubsubTopicName),
		options: DataRecordFrameOptions{
			Encrypted:       false, // not implemented yet
			FrameOutputMode: config.Config.RecorderFrameOutputMode,
		},
	}
}

func (p *pubsubRecorder) NewDataRecordFrame(r models.EvalResult) DataRecordFrame {
	_ = "STUB: not implemented"
	return *new(DataRecordFrame)
}

func (p *pubsubRecorder) AsyncRecord(r models.EvalResult) { _ = "STUB: not implemented"; return }
