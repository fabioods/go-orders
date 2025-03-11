package s3

import (
	"context"
	"mime/multipart"
)

type UploadRepository struct{}

func NewUploadRepository() *UploadRepository {
	return &UploadRepository{}
}

func (u *UploadRepository) Upload(ctx context.Context, file multipart.File, fileName string) (string, error) {
	return "", nil
}

func (u *UploadRepository) Delete(ctx context.Context, fileName string) error {
	return nil
}
