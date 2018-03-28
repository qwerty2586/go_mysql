package imSQL

import (
	"database/sql"
	"fmt"

	"github.com/juju/errors"
)

type (
	DB struct {
		Name string `json:"name" db:"name"`
	}
)

const (
	/*Create one database*/
	StmtCreateOneDatabase = `
	CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci
	`

	/*Drop a database.*/
	StmtDropOneDatabase = `
	DROP DATABASE %s	
	`

	StmtQueryUsersDB = `
	SELECT Db FROM mysql.db WHERE User = '%s'
	`
)

/*
NewDB return a new db handler.
*/
func NewDB(name string) (*DB, error) {
	newdb := new(DB)

	newdb.Name = name

	return newdb, nil
}

/*
Create one database.
*/
func (dbi *DB) CreateOneDB(db *sql.DB) error {

	Query := fmt.Sprintf(StmtCreateOneDatabase, dbi.Name)
	_, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}

/*
Drop one databases.
*/
func (dbi *DB) DropOneDB(db *sql.DB) error {

	Query := fmt.Sprintf(StmtDropOneDatabase, dbi.Name)
	_, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}
