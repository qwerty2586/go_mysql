package imSQL

import (
	"fmt"
	"testing"
)

func TestVars(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 3306, "root", "111111")
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

	err = SetDynamicVariables(db, "wait_timeout", "9998")
	if err != nil {
		t.Error(err)
	}

	err = SetDynamicVariables(db, "wait_timeout2", "9998")
	if err != nil {
		fmt.Println(err.Error())
	}

}
