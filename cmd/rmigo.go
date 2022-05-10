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
	pkg.LoadConfig()
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
	fs.Usage = func() {
		fmt.Println(`passing the path, can be "~/project/iteration" or "../iteration" or "iteration" or "./iteration"`)
	}
	var path string
	if len(args) > 2 {
		args = args[2:]
	} else {
		args = []string{"-h"}
	}
	err := fs.Parse(args)
	if err != nil {
		panic(err)
	}
	path = fs.Arg(0)
	s := strings.Split(path, "/")
	if len(s) == 0 || len(s) > 3 {
		panic(errors.New("invalid path"))
	}
	var cd = pkg.Conf.CurrentDir
	switch s[0] {
	case "~":
		cd = cd.Clear()
		cd = move(s, cd)
	case "..":
		cd = cd.Pop()
		cd = move(s, cd)
	default:
		cd = move(s, cd)
	}
	pkg.Conf.CurrentDir = cd
	fmt.Println(pkg.Conf.CurrentDir.LineAge())
	pkg.CreateConfig()
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

func move(s []string, cd *pkg.Node) *pkg.Node {
	client := pkg.NewRTCClient()
	if cd == nil || cd.Type == "void" {
		if len(s) > 1 {
			p, err := client.GetProject(s[1])
			if err != nil {
				panic(err)
			}
			cd = cd.Push(p.Node)
		}
		if len(s) > 2 {
			sprint, err := client.GetSprint(s[2])
			if err != nil {
				panic(err)
			}
			cd = cd.Push(sprint.Node)
		}
	} else if cd.Type == "project" {
		if len(s) > 1 {
			sprint, err := client.GetSprint(s[1])
			if err != nil {
				panic(err)
			}
			cd = cd.Push(sprint.Node)
		}
	}
	return cd
}
