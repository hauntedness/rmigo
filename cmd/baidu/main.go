package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hauntedness/rtigo/pkg"
)

func main() {
	if len(os.Args) < 1 {
		log.Println("rtc.exe -h for all commands")
	}

	switch os.Args[1] {
	case "login":
		Login(os.Args)
	case "cd":
		CurrentDir(os.Args)
	case "ll":
		ListAll(os.Args)
	case "cat":
		Concatenate(os.Args)
	case "edit":
		EditItem(os.Args)
	default:
		panic(errors.New("rtc.exe -h for all commands"))
	}
}

func Login(args []string) {
	fs := flag.NewFlagSet("login", flag.ExitOnError)
	var user, password string
	fs.StringVar(&user, "user", "", "username of rtc")
	fs.StringVar(&password, "password", "", "password of rtc")
	if len(args) > 2 {
		args = args[2:]
	} else {
		args = []string{"-h"}
	}
	err := fs.Parse(args)
	if err != nil {
		panic(err)
	}
	if user != "" && password != "" {
		pkg.InitConfig(user, password)
		client := pkg.NewRTCClient()
		err = client.Login(user, password)
		if err != nil {
			panic(err)
		}
		//TODO try login
		log.Println("login success")
	}
}

func CurrentDir(args []string) {
	fs := flag.NewFlagSet("cd", flag.ExitOnError)
	var path string
	fs.StringVar(&path, "path", "", `the path, can be "~/project/iteration" or "../iteration" or "iteration" or "./iteration"`)
	if len(args) > 2 {
		args = os.Args[2:]
	} else {
		args = []string{"-h"}
	}
	err := fs.Parse(args)
	if err != nil {
		panic(err)
	}
	s := strings.Split(path, "/")
	if len(s) == 0 {
		panic(errors.New("invalid path"))
	}
	switch s[0] {
	case "~":
		pkg.Conf.CurrentDir = pkg.Conf.CurrentDir.Clear()
		//TODO, add one by one
		for _, v := range s[1:] {
			pkg.NewRTCClient().ListSprints(v)
		}
	case "..":
		pkg.Conf.CurrentDir = pkg.Conf.CurrentDir.Pop()
		//TODO, add one by one
	default:
		//TODO, add sub and sub of sub and ...
	}
	fmt.Println(pkg.Conf.CurrentDir.LineAge())
}

func ListAll(args []string) {
	fs := flag.NewFlagSet("cd", flag.ExitOnError)
	var a, i, p, d bool
	fs.BoolVar(&a, "a", false, "list all items instead of current users")
	fs.BoolVar(&i, "i", false, "list all items of iteration i instead of current iteration")
	fs.BoolVar(&p, "p", false, "list all items of project p instead of current project")
	fs.BoolVar(&d, "d", false, "list all defects instead of all stories")
	if len(args) > 2 {
		args = os.Args[2:]
	} else {
		args = []string{"-h"}
	}
	err := fs.Parse(args)
	if err != nil {
		panic(err)
	}

}

func Concatenate(args []string) {

}

func EditItem(args []string) {

}
