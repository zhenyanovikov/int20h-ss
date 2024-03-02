package mail

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go/service/ses"
	"oss-backend/internal/models"
)

func (s Service) SendEmail(ctx context.Context, recipient *models.User, subject, text string) error {
	email, err := s.sesClient.SendEmailWithContext(ctx, &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{&recipient.Email},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: &text,
				},
			},
			Subject: &ses.Content{
				Data: &subject,
			},
		},
		Source: s.sourceEmail,
	})
	if err != nil {
		return fmt.Errorf("failed to send email, %v", err)
	}

	slog.Info(fmt.Sprintf("Email sent to %s with message ID %s", recipient.Email, *email.MessageId))

	return nil
}
