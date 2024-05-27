package helper

import "database/sql"

func CommirOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		PanicIfError(errRollback)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}
