package helper

import "database/sql"

func CommitOrRollBack(tx *sql.Tx) {
	err := recover()
	if err != nil {
		er := tx.Rollback()
		IfPanic(er)
		panic(err)
	} else {
		er := tx.Commit()
		if er != nil {
			panic(er)
		}
	}
}
