package controllers

import (
	"database/sql"

	"github.com/revel/revel"
)

// Transact is a nice interface to use for database transactions
// TODO: Possibly update this for Query's
func Transact(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			revel.WARN.Printf("Rolling back transaction: %s", err)
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	err = txFunc(tx)
	return err
}
