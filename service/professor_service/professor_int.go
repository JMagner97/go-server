package professorservice

import (
	"context"
	"go-server/data/request"
)

type ProfessorService interface {
	Create(ctx context.Context, request request.ProfessorRequest) (bool, error)
}
