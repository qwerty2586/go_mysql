package mysqlmanage

import (
	"flag"
	"fmt"
	"testing"
)

func TestQueryAllVars(t *testing.T) {
	flag.Parse()
	conn, err := NewConn(mysql_addr, mysql_port, mysql_user, mysql_pass)
	if err != nil {
		t.Error(conn, err)
	}

	conn.SetCharset("utf8")
	conn.SetCollation("utf8_general_ci")
	conn.MakeDBI()

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	vars, err := ShowVariables(db)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(vars)

}

func TestUpdateOneVars(t *testing.T) {
	flag.Parse()
	conn, err := NewConn(mysql_addr, mysql_port, mysql_user, mysql_pass)
	if err != nil {
		t.Error(conn, err)
	}

	conn.SetCharset("utf8")
	conn.SetCollation("utf8_general_ci")
	conn.MakeDBI()

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = SetDynamicVariables(db, "wait_timeout", "9998")
	if err != nil {
		t.Error(err)
	}

}

func TestUpdateOneNotExistsVars(t *testing.T) {
	flag.Parse()
	conn, err := NewConn(mysql_addr, mysql_port, mysql_user, mysql_pass)
	if err != nil {
		t.Error(conn, err)
	}

	conn.SetCharset("utf8")
	conn.SetCollation("utf8_general_ci")
	conn.MakeDBI()

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = SetDynamicVariables(db, "wait_timeout2", "9998")
	if err != nil {
		fmt.Println(err.Error())
	}

}
