package imSQL

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
		PasswordLastChanged  string `json:"password_last_changed" db:"password_last_changed"`
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
		password_last_changed,
		password_lifetime,
		account_locked
	FROM 
		mysql.user
	WHERE
		User NOT IN ('root','mysql.sys','mysql.session')
	AND
		Host = 'localhost'
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
	PASSWORD EXPIRE NEVER
	ACCOUNT '%s'
	`

	//delete a user
	StmtDeleteOneUser = `
	DROP USER IF EXISTS '%s'@'%s'
	`
)
