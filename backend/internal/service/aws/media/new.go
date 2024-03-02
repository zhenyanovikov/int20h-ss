package media

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"oss-backend/internal/config"
)

type Service struct {
	s3uploader *s3manager.Uploader
}

func New(cfg config.Config) (*Service, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWS.Region),
		Credentials: credentials.NewStaticCredentials(
			cfg.AWS.AccessKeyID,
			cfg.AWS.SecretAccessKey,
			"",
		),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create session, %v", err)
	}

	return &Service{
		s3uploader: s3manager.NewUploader(sess),
	}, nil
}
