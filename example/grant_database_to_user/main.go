package main

import (
	"github.com/qwerty2586/go_mysql/mysqlmanage"
	"log"
)

func t(err error) {
	log.Fatalln("terminating with error: ", err)
}

func main() {
	mysqlmanage.DEGUG = true

	conn, err := mysqlmanage.NewConn("localhost", 3306, "root", "root")
	if err != nil {
		t(err)
	}

	conn.SetCharset("utf8")
	conn.SetCollation("utf8_general_ci")
	conn.MakeDBI()

	db, err := conn.OpenConn()
	if err != nil {
		t(err)
	}

	newuser, _ := mysqlmanage.NewUser("dev", "dev", "localhost")
	err = newuser.AddOneUser(db)
	if err != nil {
		t(err)
	}

	newdb, _ := mysqlmanage.NewDB("dev")
	err = newdb.CreateOneDB(db)
	if err != nil {
		t(err)
	}

	// granting newuser privileges on newdb
	newuser.PullGrants(db)
	newuser.AddGrantsForDB(newdb, []string{"ALL PRIVILEGES"})
	newuser.PushGrants(db, true)

}
