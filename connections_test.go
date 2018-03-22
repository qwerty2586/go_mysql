package imSQL

import (
	"os"
	"strconv"
	"testing"
)

var mysql_addr = os.Getenv("MYSQL_ADDR")
var mysql_port_tmp, _ = strconv.Atoi(os.Getenv("MYSQL_PORT"))
var mysql_port = uint64(mysql_port_tmp)
var mysql_user = os.Getenv("MYSQL_USER")
var mysql_pass = os.Getenv("MYSQL_PASS")

func TestNewConn(t *testing.T) {

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
}
