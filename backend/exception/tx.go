package exception

import (
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) error {
	err := recover()
	if err != nil {
		return tx.Rollback()
	} else {
		return tx.Commit()
	}
}
