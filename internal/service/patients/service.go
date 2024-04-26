package patients

import "github.com/drizzleent/patients/internal/repository"

type srv struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *srv {
	return &srv{
		repo: repo,
	}
}
