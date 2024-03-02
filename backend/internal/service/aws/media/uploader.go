package media

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"oss-backend/internal/config"
)

func (s *Service) Upload(ctx context.Context, fileReader io.Reader, extension string) (url string, err error) {
	filename := "media/" + uuid.New().String() + extension
	result, err := s.s3uploader.UploadWithContext(ctx,
		&s3manager.UploadInput{
			Bucket: aws.String(config.Get().AWS.S3.Bucket),
			Key:    aws.String(filename),
			Body:   fileReader,
		})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	return result.Location, nil
}
