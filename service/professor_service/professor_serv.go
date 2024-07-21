package professorservice

import (
	"context"
	"errors"
	"go-server/Model"
	repository "go-server/Repository"
	"go-server/data/request"
)

type profServ struct {
	ProfessorRepo repository.ProfessorRepo
}

// Create implements ProfessorService.
func (s *profServ) Create(ctx context.Context, request request.ProfessorRequest) (bool, error) {
	student := Model.Professor{
		Name: request.Name,
		//Id:           request.Id,
		Username:     request.Username,
		Password:     request.Password,
		Role:         2,
		Surname:      request.Surname,
		Email:        request.Email,
		DepartmentId: request.DepartmentId,
	}
	exists, err := s.ProfessorRepo.ProfessorExists(ctx, &student)
	if err != nil {
		return false, err
	}
	if !exists {
		innerResult, err := s.ProfessorRepo.Save(ctx, student)
		return innerResult, err
	} else {
		return false, errors.New("professor already exists")
	}
}

func NewProfessorService(profRepo repository.ProfessorRepo) ProfessorService {
	return &profServ{ProfessorRepo: profRepo}
}
