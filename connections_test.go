package imSQL

import (
	"flag"
	"testing"
)

var mysql_addr = flag.String("addr", "127.0.0.1", "mysql listen address.default 127.0.0.1")
var mysql_port = flag.Uint64("port", 3306, "mysql listen port,default 3306")
var mysql_user = flag.String("user", "root", "mysql administrator name.default root")
var mysql_pass = flag.String("pass", "111111", "mysql administrator password.default 111111")

func TestNewConn(t *testing.T) {

	flag.Parse()
	conn, err := NewConn(*mysql_addr, *mysql_port, *mysql_user, *mysql_pass)
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
