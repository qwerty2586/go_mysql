package mysqlmanage

import "fmt"

var DEGUG = false

func debug(s string) {
	if DEGUG {
		fmt.Println("mysqlmanage debug: ", s)
	}
}
