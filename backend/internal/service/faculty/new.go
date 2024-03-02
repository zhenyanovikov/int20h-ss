package faculty

import "oss-backend/internal/persistence"

type Service struct {
	facultyRepo persistence.Faculty
}

func New(facultyRepo persistence.Faculty) *Service {
	return &Service{
		facultyRepo: facultyRepo,
	}
}
