package imSQL

import "testing"

func TestUser(t *testing.T) {
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

	allusers, err := FindAllUserInfo(db, 1, 0)
	if err != nil {
		t.Error(allusers, err)
	}

	newuser, err := NewUser("dev2", "dev2", "localhost")
	if err != nil {
		t.Error(newuser, err)
	}

	err = newuser.AddOneUser(db)
	if err != nil {
		t.Error(err)
	}
}
