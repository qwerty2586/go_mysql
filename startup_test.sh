#########################################################################
# File Name: startup_test.sh
# Author: Lei Tian
# mail: taylor840326@gmail.com
# Created Time: 2018-02-06 17:20
#########################################################################
#!/bin/bash

export MYSQL_ADDR="172.18.10.136"
export MYSQL_PORT=3306
export MYSQL_USER="root"
export MYSQL_PASS="111111"

# Test connections.
go test -timeout 30m -v -test.run TestNewConn

# Test users.
go test -timeout 30m -v -test.run TestQueryAllUsers
go test -timeout 30m -v -test.run TestCreateOneUser
go test -timeout 30m -v -test.run TestQueryAllUsers
go test -timeout 30m -v -test.run TestUpdateOneUser
go test -timeout 30m -v -test.run TestQueryAllUsers
go test -timeout 30m -v -test.run TestDeleteOneUser
go test -timeout 30m -v -test.run TestQueryAllUsers

# Test variables.
go test -timeout 30m -v -test.run TestVars