package sqs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awssqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/fabioods/go-orders/internal/config"
	"github.com/fabioods/go-orders/internal/errorcode"
	"github.com/fabioods/go-orders/pkg/errorformatted"
	"github.com/fabioods/go-orders/pkg/trace"
)

type SQSRepository struct {
	Client *awssqs.Client
	Queue  string
}

func NewSQSRepository(cfg *config.Config) *SQSRepository {
	opsts := []func(*awsCfg.LoadOptions) error{
		awsCfg.WithRegion(cfg.S3Config.S3Region),
	}

	if cfg.SQSConfig.SQSAccessKey != "" && cfg.SQSConfig.SQSSecretKey != "" {
		opsts = append(opsts, awsCfg.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.SQSConfig.SQSAccessKey,
			cfg.SQSConfig.SQSSecretKey,
			"",
		)))
	}

	sqsCfg, err := awsCfg.LoadDefaultConfig(context.TODO(), opsts...)
	if err != nil {
		panic("Error to load SQS config")
	}

	client := awssqs.NewFromConfig(sqsCfg)
	return &SQSRepository{
		Client: client,
		Queue:  cfg.SQSConfig.SQSQueue,
	}
}

func (s *SQSRepository) SendMessage(ctx context.Context, body interface{}) error {
	msgBody, err := json.Marshal(body)
	if err != nil {
		return errorformatted.UnexpectedError(trace.GetTrace(), errorcode.ErrorSendSQSMessageMarshal, "Error to marshal message to SQS %s", err.Error())
	}

	result, err := s.Client.SendMessage(ctx, &awssqs.SendMessageInput{
		MessageBody: aws.String(string(msgBody)),
		QueueUrl:    aws.String(s.Queue),
	})
	if err != nil {
		return errorformatted.UnexpectedError(trace.GetTrace(), errorcode.ErrorSendSQSMessage, "Error to send message to SQS %s", err.Error())
	}

	fmt.Printf("Message sent! ID: %s\n", *result.MessageId)

	return nil
}
