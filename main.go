package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	flag.StringVar(&user, "u", "", "帐号")
	flag.StringVar(&pwd, "p", "", "密码")
	flag.Parse()
	fmt.Printf("user: %s, password: %s \n", user, pwd)

}
