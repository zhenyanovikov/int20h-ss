package notifier

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"oss-backend/internal/config"
	"oss-backend/internal/persistence"
)

type Service struct {
	repo        persistence.Repo
	sesClient   *ses.SES
	sourceEmail *string
}

func New(cfg config.Config, repo persistence.Repo) (*Service, error) {
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
		repo:        repo,
		sesClient:   sesClient,
		sourceEmail: aws.String(cfg.AWS.SES.SourceEmail),
	}, nil
}
