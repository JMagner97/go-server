package repository

import (
	"context"
	"database/sql"
	"go-server/Model"
)

type professorRepo struct{ db *sql.DB }

// ProfessorExists implements ProfessorRepo.
func (s *professorRepo) ProfessorExists(ctx context.Context, professor *Model.Professor) (bool, error) {
	var existingId int
	err := s.db.QueryRow("SELECT professorid FROM professors WHERE email = $1", professor.Email).Scan(&existingId)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, handleSQLError(err)
	}
	return existingId > 0, nil
}

// Save implements ProfessorRepo.
func (s *professorRepo) Save(ctx context.Context, professor Model.Professor) (bool, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return false, handleSQLError(err)
	}
	defer tx.Rollback() // Assicurati che la transazione sia annullata in caso di errore

	// Prima, inserisci l'utente nella tabella Users e ottieni l'userID generato
	userSQL := "INSERT INTO users(username, password, roleID) VALUES ($1, $2, $3) RETURNING userID"
	var userID int
	err = tx.QueryRowContext(ctx, userSQL, professor.Username, professor.Password, professor.Role).Scan(&userID)
	if err != nil {
		return false, handleSQLError(err)
	}

	// Poi, inserisci lo studente nella tabella Students con userID ottenuto
	studentSQL := "INSERT INTO professors(professorid, name, surname, email, departmentID) VALUES ($1, $2, $3, $4, $5)"
	_, err = tx.ExecContext(ctx, studentSQL, userID, professor.Name, professor.Surname, professor.Email, professor.DepartmentId)
	if err != nil {
		return false, handleSQLError(err)
	}

	// Se tutto va bene, conferma la transazione
	err = tx.Commit()
	if err != nil {
		return false, handleSQLError(err)
	}

	return true, nil
}

func NewProfessorRepo(db *sql.DB) ProfessorRepo {
	return &professorRepo{db: db}
}
