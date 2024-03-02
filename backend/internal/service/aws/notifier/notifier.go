package notifier

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/google/uuid"
	"oss-backend/internal/models"
	"oss-backend/internal/persistence"
)

func (s Service) SendNotification(ctx context.Context, dto *models.SendNotificationDTO) error {
	var receiverStudents []models.Student

	if len(dto.StudentIDs) > 0 {
		students, err := s.repo.User().ListStudents(ctx, persistence.WithIDIn(dto.StudentIDs))
		if err != nil {
			return fmt.Errorf("failed to get students: %v", err)
		}

		receiverStudents = append(receiverStudents, students...)
	}

	if dto.GroupID != uuid.Nil {
		students, err := s.repo.User().ListStudentsByGroupID(ctx, dto.GroupID)
		if err != nil {
			return fmt.Errorf("failed to get students: %v", err)
		}

		receiverStudents = append(receiverStudents, students...)
	}

	if dto.FacultyID != uuid.Nil {
		students, err := s.repo.User().ListStudentsByFacultyID(ctx, dto.FacultyID)
		if err != nil {
			return fmt.Errorf("failed to get groups: %v", err)
		}

		receiverStudents = append(receiverStudents, students...)
	}

	if len(receiverStudents) == 0 {
		return fmt.Errorf("no students found")
	}

	for _, student := range receiverStudents {
		if err := s.sendEmail(ctx, student.User, dto.Subject, dto.Body); err != nil {
			slog.Error("failed to send email, %v", err)
		}
	}

	return nil
}

func (s Service) sendEmail(ctx context.Context, recipient *models.User, subject, text string) error {
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
