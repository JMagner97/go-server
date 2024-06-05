package helper

import "database/sql"

func CommirOrRollback(tx *sql.Tx) error {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		return errRollback
	} else {
		errCommit := tx.Commit()
		return errCommit
	}
}
