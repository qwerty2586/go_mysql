package mysqlmanage

import (
	"database/sql"
	"fmt"
	"strings"
)

type (
	Users struct {
		Host                 string   `json:"host" db:"host"`
		User                 string   `json:"username" db:"user"`
		Active               uint64   `json:"active" db:"active"`
		DefaultSchema        string   `json:"default_schema" db:"default_schema"`
		MaxQuestions         uint64   `json:"max_questions" db:"max_questions"`
		MaxUpdates           uint64   `json:"max_updates" db:"max_updates"`
		MaxConnections       uint64   `json:"max_connections" db:"max_connections"`
		MaxUserConnections   uint64   `json:"max_user_connections" db:"max_user_connections"`
		Plugin               string   `json:"plugin" db"plugin"`
		AuthenticationString string   `json:"password" db:"authentication_string"`
		PasswordExpired      string   `json:"password_expired" db:"password_expired"`
		PasswordLifetime     uint64   `json:"password_lifetime" db:"password_lifetime"`
		AccountLocked        string   `json:"account_locked" db:"account_locked"`
		Privileges           []string `json:"privileges" db:"privileges"`
	}
)

const (
	//get one user info.
	StmtQueryOneUserInfo = `
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
		IFNULL(password_lifetime,0),
		account_locked
	FROM 
		mysql.user
	WHERE
		User = '%s'
	AND
		Host = '%s'
	`
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
		IFNULL(password_lifetime,0),
		account_locked
	FROM 
		mysql.user
	WHERE
		User NOT IN ('root','mysql.sys','mysql.session')
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
	PASSWORD EXPIRE %s    
	ACCOUNT %s
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
	PASSWORD EXPIRE %s    
	ACCOUNT %s
	`

	//delete a user
	StmtDeleteOneUser = `
	DROP USER IF EXISTS '%s'@'%s'
	`
	StmtFlushPrivileges = `
	FLUSH PRIVILEGES
	`
	StmtGrantPrivileges = `
	GRANT %s ON %s.* TO '%s'@'%s'
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

	newuser.DefaultSchema = "information_schema"
	newuser.MaxQuestions = 0
	newuser.MaxUpdates = 0
	newuser.MaxConnections = 0
	newuser.MaxUserConnections = 0
	newuser.Plugin = "mysql_native_password"
	newuser.PasswordExpired = "N"
	newuser.PasswordLifetime = 0
	newuser.AccountLocked = "N"
	newuser.Privileges = []string{"ALL PRIVILEGES"}

	return newuser, nil
}

/*
set user password
*/
func (user *Users) SetPassword(password string) {
	switch {
	case password == "":
		user.AuthenticationString = password
	default:
		user.AuthenticationString = password

	}
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
set max user connections.
*/
func (user *Users) SetMaxUserConections(max_user_connections uint64) {
	user.MaxUserConnections = max_user_connections
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
func (user *Users) SetAccountLocked(active uint64) {
	if active == 0 {
		user.AccountLocked = "Y"
	} else {
		user.AccountLocked = "N"
	}
}

/*
add user's privileges
*/
func (user *Users) AddPrivileges(privileges ...string) {
	if len(privileges) != 0 {
		if user.Privileges[0] == "ALL PRIVILEGES" {
			user.Privileges = []string{}
			user.Privileges = append(user.Privileges, privileges...)
		} else {
			user.Privileges = append(user.Privileges, privileges...)
		}
	}
}

/*
set user't default schema.
*/
func (user *Users) SetDefaultSchema(default_schema string) {
	user.DefaultSchema = default_schema
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
		return err
	}

	PrivList := strings.Join(user.Privileges, ",")
	GrantQuery := fmt.Sprintf(StmtGrantPrivileges, PrivList, user.DefaultSchema, user.User, user.Host)

	_, err = db.Exec(GrantQuery)
	if err != nil {
		return err
	}

	_, err = db.Exec(StmtFlushPrivileges)
	if err != nil {
		return err
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

	_, err := db.Exec(Query)
	if err != nil {
		return err
	}

	return nil
}

/*
drop user.
*/
func (user *Users) DeleteOneUser(db *sql.DB) error {

	Query := fmt.Sprintf(StmtDeleteOneUser, user.User, user.Host)

	_, err := db.Exec(Query)
	if err != nil {
		return err
	}

	return nil
}

/*
get one user information.
*/
func (user *Users) FindOneUserInfo(db *sql.DB) (Users, error) {

	var tmpuser Users
	Query := fmt.Sprintf(StmtQueryOneUserInfo, user.User, user.Host)

	rows, err := db.Query(Query)
	if err != nil {
		return Users{}, err
	}
	//defer rows.Close()

	for rows.Next() {

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

	}
	return tmpuser, nil
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
		return []Users{}, err
	}
	//defer rows.Close()

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
