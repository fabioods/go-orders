package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awssqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/fabioods/go-orders/internal/infra/sqs"
	"github.com/fabioods/go-orders/internal/usecase"
)

type OrderConsumer struct {
	ProcessOrderUseCase ProcessOrderUseCase
}

type ProcessOrderUseCase interface {
	Execute(ctx context.Context, input usecase.ProcessOrderInput) error
}

func NewOrderConsumer(processOrderUseCase ProcessOrderUseCase) *OrderConsumer {
	return &OrderConsumer{
		ProcessOrderUseCase: processOrderUseCase,
	}
}

func (o *OrderConsumer) ConsumeMessages(sqsRepository *sqs.SQSRepository) {
	ctx := context.Background()
	for {
		result, err := sqsRepository.Client.ReceiveMessage(ctx, &awssqs.ReceiveMessageInput{
			QueueUrl:            aws.String(sqsRepository.Queue),
			MaxNumberOfMessages: 5,
			WaitTimeSeconds:     10,
		})

		if err != nil {
			log.Printf("Error to receive messages: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if len(result.Messages) == 0 {
			time.Sleep(2 * time.Second)
			continue
		}

		for _, msg := range result.Messages {
			input := usecase.ProcessOrderInput{}
			fmt.Println(*msg.Body)
			if err := json.Unmarshal([]byte(*msg.Body), &input); err != nil {
				log.Printf("Error to unmarshal message: %v", err)
				continue
			}
			if err := o.ProcessOrderUseCase.Execute(ctx, input); err == nil {
				_, err := sqsRepository.Client.DeleteMessage(ctx, &awssqs.DeleteMessageInput{
					QueueUrl:      aws.String(sqsRepository.Queue),
					ReceiptHandle: msg.ReceiptHandle,
				})
				if err != nil {
					log.Printf("Erro to delete the message: %v", err)
				}
			} else {
				log.Printf("Error to process the message: %v", err)
			}
		}
	}
}
