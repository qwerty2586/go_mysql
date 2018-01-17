package imSQL

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/juju/errors"
)

type (
	Users struct {
		Host                 string `json:"host" db:"host"`
		User                 string `json:"user" db:"user"`
		MaxQuestions         uint64 `json:"max_questions" db:"max_questions"`
		MaxUpdates           uint64 `json:"max_updates" db:"max_updates"`
		MaxConnections       uint64 `json:"max_connections" db:"max_connections"`
		MaxUserConnections   uint64 `json:"max_user_connections" db:"max_user_connections"`
		Plugin               string `json:"plugin" db"plugin"`
		AuthenticationString string `json:"authentication_string" db:"authentication_string"`
		PasswordExpired      string `json:"password_expired" db:"password_expired"`
		PasswordLifetime     uint64 `json:"password_lifetime" db:"password_lifetime"`
		AccountLocked        string `json:"account_locked" db:"account_locked"`
	}
)

const (
	//get all users infor.
	StmtQueryAllUsersInfo = `
	SELECT 
		Host,
		User,
		max_questions,
		max_updates,
		max_connections,
		max_user_connections,
		plugin,
		authentication_string,
		password_expired,
		password_lifetime,
		account_locked
	FROM 
		mysql.user
	WHERE
		User NOT IN ('root','mysql.sys','mysql.session')
	AND
		Host = 'localhost'
	LIMIT %d OFFSET %d
	`

	//create a user
	StmtAddOneUser = `
	CREATE USER IF NOT EXISTS 
	'%s'@'%s'
	IDENTIFIED BY '%s'
	WITH 
	MAX_QUERIES_PER_HOUR %d
	MAX_UPDATES_PER_HOUR %d
	MAX_CONNECTIONS_PER_HOUR %d
	MAX_USER_CONNECTIONS %d
	PASSWORD EXPIRE '%s'    
	ACCOUNT '%s'
	`

	//alter a user
	StmtUpdateOneUser = `
	ALTER USER IF EXISTS
	'%s'@'%s'
	IDENTIFIED BY '%s'
	WITH
	MAX_QUERIES_PER_HOUR %d
	MAX_UPDATES_PER_HOUR %d
	MAX_CONNECTIONS_PER_HOUR %d
	MAX_USER_CONNECTIONS %d
	PASSWORD EXPIRE '%s'    
	ACCOUNT '%s'
	`

	//delete a user
	StmtDeleteOneUser = `
	DROP USER IF EXISTS '%s'@'%s'
	`
)

/*
NewUser return a new user handler.
This function have three args,other args is options.
*/
func NewUser(username string, password string, addr string) (*Users, error) {

	newuser := new(Users)

	newuser.Host = addr
	newuser.User = username
	newuser.AuthenticationString = password

	newuser.MaxQuestions = 0
	newuser.MaxUpdates = 0
	newuser.MaxConnections = 0
	newuser.MaxUserConnections = 0
	newuser.Plugin = "mysql_native_password"
	newuser.PasswordExpired = "N"
	newuser.PasswordLifetime = 0
	newuser.AccountLocked = "N"

	return newuser, nil
}

/*
SetMaxQuestions will set user max qps.
*/
func (user *Users) SetMaxQuestions(max_questions uint64) {
	user.MaxQuestions = max_questions
}

/*
SetMaxUpdates will set user max updates.
*/
func (user *Users) SetMaxUpdates(max_updates uint64) {
	user.MaxUpdates = max_updates
}

/*
SetMaxConnections will set max connections.
*/
func (user *Users) SetMaxConnections(max_connections uint64) {
	user.MaxConnections = max_connections
}

/*
set user password life time.
*/
func (user *Users) SetPasswordLifeTime(password_lifetime uint64) {
	user.PasswordLifetime = password_lifetime
}

/*
enable/disable user password expired.
*/
func (user *Users) SetPasswordExipred(password_expired string) {
	user.PasswordExpired = password_expired
}

/*
lock/unlock user account.
*/
func (user *Users) SetAccountLocked(account_locked string) {
	user.AccountLocked = account_locked
}

/*
add one user.
*/
func (user *Users) AddOneUser(db *sql.DB) error {

	var password_option string
	var lock_option string

	// set password expire option
	if user.PasswordExpired == "N" {
		switch {
		case user.PasswordLifetime == 0:
			password_option = fmt.Sprint("NEVER")
		case user.PasswordLifetime >= 360:
			password_option = fmt.Sprint("DEFAULT")
		default:
			password_option = fmt.Sprintf("INTERVAL %d DAY", user.PasswordLifetime)
		}
	} else {
		password_option = fmt.Sprint(" ")
	}

	//set lock option.
	if user.AccountLocked == "N" {
		lock_option = fmt.Sprint("UNLOCK")
	} else {
		lock_option = fmt.Sprint("LOCK")

	}

	//Query Stmt.
	Query := fmt.Sprintf(StmtAddOneUser, user.User, user.Host, user.AuthenticationString, user.MaxQuestions, user.MaxUpdates, user.MaxConnections, user.MaxUserConnections, password_option, lock_option)

	_, err := db.Exec(Query)
	if err != nil {
		switch {
		//user is exists.
		case err.(*mysql.MySQLError).Number == 1045:
			return errors.NewAlreadyExists(err, user.User)
		default:
			return errors.Trace(err)
		}
	}
	return nil

}

/*
alter user ...
*/
func (user *Users) UpdateOneUser(db *sql.DB) error {

	var password_option string
	var lock_option string

	// set password expire option
	if user.PasswordExpired == "N" {
		switch {
		case user.PasswordLifetime == 0:
			password_option = fmt.Sprint("NEVER")
		case user.PasswordLifetime >= 360:
			password_option = fmt.Sprint("DEFAULT")
		default:
			password_option = fmt.Sprintf("INTERVAL %d DAY", user.PasswordLifetime)
		}
	} else {
		password_option = fmt.Sprint(" ")
	}

	//set lock option.
	if user.AccountLocked == "N" {
		lock_option = fmt.Sprint("UNLOCK")
	} else {
		lock_option = fmt.Sprint("LOCK")

	}

	// Query Stmt.
	Query := fmt.Sprintf(StmtUpdateOneUser, user.User, user.Host, user.AuthenticationString, user.MaxQuestions, user.MaxUpdates, user.MaxConnections, user.MaxUserConnections, password_option, lock_option)

	result, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.NotFoundf(user.User)
	}

	return nil
}

/*
drop user.
*/
func (user *Users) DeleteOneUser(db *sql.DB) error {

	Query := fmt.Sprintf(StmtDeleteOneUser, user.User, user.Host)

	result, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.NotFoundf(user.User)
	}

	return nil
}

/*
get all users information.
*/
func FindAllUserInfo(db *sql.DB, limit uint64, skip uint64) ([]Users, error) {

	//save users info.
	var allusers []Users

	Query := fmt.Sprintf(StmtQueryAllUsersInfo, limit, skip)

	rows, err := db.Query(Query)
	if err != nil {
		return []Users{}, errors.Trace(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tmpuser Users

		err = rows.Scan(
			&tmpuser.Host,
			&tmpuser.User,
			&tmpuser.MaxQuestions,
			&tmpuser.MaxUpdates,
			&tmpuser.MaxConnections,
			&tmpuser.MaxUserConnections,
			&tmpuser.Plugin,
			&tmpuser.AuthenticationString,
			&tmpuser.PasswordExpired,
			&tmpuser.PasswordLifetime,
			&tmpuser.AccountLocked,
		)
		if err != nil {
			continue
		}

		allusers = append(allusers, tmpuser)
	}
	return allusers, nil
}
