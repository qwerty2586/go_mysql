# Golang MySQL Management Library
-----

### 1. introduce
-----

A MySQL Go Library.


### 2. Requirements
-----

1. Go 1.7+
1. MySQL 5.7


### 3.Installation
-----

Simple install the package to your $GOPATH with the go tool from shell:

    # go get -u github.com/qwerty2586/go_mysql/mysqlmanage

Make sure git command is installed on your OS.

#### 4. Usage
-----

example:

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

	allusers, err := FindAllUserInfo(db, 10, 0)
	if err != nil {
		t.Error(allusers, err)
	}

	fmt.Println(allusers)

	newuser, err := NewUser("dev2", "dev2", "localhost")
	if err != nil {
		t.Error(newuser, err)
	}

	err = newuser.AddOneUser(db)
	if err != nil {
		t.Error(err)
	}

	newuser.SetAccountLocked("YES")
	newuser.SetMaxConnections(10000)
	newuser.SetMaxUserConections(1000)
	newuser.SetMaxQuestions(10)
	newuser.SetMaxUpdates(2)

	err = newuser.UpdateOneUser(db)
	if err != nil {
		t.Error(err)
	}

	err = newuser.DeleteOneUser(db)
	if err != nil {
		t.Error(err)
	}

### Donate
-----

if you like this project and want to buy me a cola,you can through:

| PayPal                                                                                                               | 微信                                                                 |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------- |
| [![](https://www.paypalobjects.com/webstatic/paypalme/images/pp_logo_small.png)](https://www.paypal.me/taylor840326) | ![](https://github.com/taylor840326/blog/raw/master/imgs/weixin.png) |


