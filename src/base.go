package main

import (
	set "./setup"
	"bufio"
	"log"
	"os"
	//"regist"; "github.com/UltimaQuazar/go-mastodon"
	//"fmt"
)

var (
	read_uname = func() (string, error) {
		uname, _, err := bufio.NewReader(os.Stdin).ReadLine()
		if err != nil {
			return "Err: ", err
		}
		return string(uname), nil
	}
	err_chk = func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
)

func main() {
	app := set.Creat()
	err := app.Run(os.Args)
	err_chk(err)
}
