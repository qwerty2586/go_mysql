package imSQL

import "testing"

func TestNewConn(t *testing.T) {
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
}
