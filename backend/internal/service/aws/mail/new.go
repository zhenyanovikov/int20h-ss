package mail

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"oss-backend/internal/config"
)

type Service struct {
	sesClient   *ses.SES
	sourceEmail *string
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

	sesClient := ses.New(sess)

	return &Service{
		sesClient:   sesClient,
		sourceEmail: aws.String(cfg.AWS.SES.SourceEmail),
	}, nil
}
