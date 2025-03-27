package s3

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/fabioods/go-orders/internal/config"
	"github.com/fabioods/go-orders/internal/errorcode"
	"github.com/fabioods/go-orders/pkg/errorformatted"
	"github.com/fabioods/go-orders/pkg/trace"
)

type UploadRepository struct {
	cfg    *config.Config
	client *s3.Client
}

func NewUploadRepository(cfg *config.Config) *UploadRepository {
	opsts := []func(*awsCfg.LoadOptions) error{
		awsCfg.WithRegion(cfg.S3Config.S3Region),
	}

	if cfg.S3Config.S3AccessKey != "" && cfg.S3Config.S3SecretKey != "" {
		opsts = append(opsts, awsCfg.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.S3Config.S3AccessKey,
			cfg.S3Config.S3SecretKey,
			"",
		)))
	}

	awsCfg, err := awsCfg.LoadDefaultConfig(context.TODO(), opsts...)

	if err != nil {
		panic(fmt.Sprintf("Error to load S3 config: %v", err))
	}

	client := s3.NewFromConfig(awsCfg)

	return &UploadRepository{
		cfg:    cfg,
		client: client,
	}
}

func (u *UploadRepository) Upload(ctx context.Context, file multipart.File, fileName string) (string, error) {
	_, err := u.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(u.cfg.S3Config.S3Bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		errFtm := errorformatted.UnexpectedError(trace.GetTrace(), errorcode.ErrorSendS3File, "Error to send file to S3 %s", err.Error())
		return "", errFtm
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.cfg.S3Config.S3Bucket, u.cfg.S3Config.S3Region, fileName)

	return url, nil
}

func (u *UploadRepository) Delete(ctx context.Context, fileName string) error {
	_, err := u.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(u.cfg.S3Config.S3Bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return errorformatted.UnexpectedError(trace.GetTrace(), errorcode.ErrorDeleteS3File, "Error to delete file from S3 %s", err.Error())
	}

	return nil
}
