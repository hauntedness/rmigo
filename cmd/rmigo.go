package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/hauntedness/rtigo/pkg"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) == 2 && (strings.ToLower(os.Args[1]) == "-h" || strings.ToLower(os.Args[1]) == "--help") {
		fmt.Printf("commands:\n\trmigo login\n\trmigo cd\n\trmigo ll\n\trmigo search\n\trmigo cat\n\trmigo edit")
		os.Exit(0)
	}
	pkg.LoadConfig()
	switch os.Args[1] {
	case "login":
		Login(os.Args)
	case "cd":
		CurrentDir(os.Args)
	case "ll":
		ListAll(os.Args)
	case "search":
		Search(os.Args)
	case "cat":
		Cat(os.Args)
	case "edit":
		EditItem(os.Args)
	default:
		panic(errors.New("rtc.exe -h for all commands"))
	}
}

func Login(args []string) {
	fs := flag.NewFlagSet("login", flag.ExitOnError)
	var user, password string
	var err error
	fs.StringVar(&user, "user", "", "username of rtc")
	fs.StringVar(&password, "password", "", "password of rtc")
	parse(fs, args)
	if user != "" && password != "" {
		pkg.InitConfig(user, password)
		client := pkg.NewRTCClient()
		err = client.Login(user, password)
		if err != nil {
			panic(err)
		}
		//TODO try login
		fmt.Println("login success")
	}
}

func CurrentDir(args []string) {
	fs := flag.NewFlagSet("cd", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println("Usage of cd:")
		fs.PrintDefaults()
		fmt.Println("rmigo cd ~/proj/xx")
		fmt.Println("rmigo cd ../xx")
		fmt.Println("rmigo cd xx")
		fmt.Println("rmigo cd ./xx")
	}
	var path string
	parse(fs, args)
	path = fs.Arg(0)
	s := strings.Split(path, "/")
	if len(s) == 0 || len(s) > 3 {
		panic(errors.New("invalid path:" + path))
	}
	var cd = pkg.Conf.CurrentDir
	switch s[0] {
	case "~":
		cd = cd.Clear()
	case "..":
		cd = cd.Pop()
	}
	cd = move(s, cd)
	pkg.Conf.CurrentDir = cd
	fmt.Println(pkg.Conf.CurrentDir.LineAge())
	pkg.CreateConfig()
}

func ListAll(args []string) {
	fs := flag.NewFlagSet("ll", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println("Usage of ll:")
		fs.PrintDefaults()
		fmt.Println("")
		fmt.Println("rmigo ll      \tlist all stories of current iteration")
		fmt.Println("rmigo ll ../xx\tlist all stories of sibling iteration xx")
		fmt.Println("rmigo ll -d   \tlist all defects of current iteration")
	}
	var d bool
	var p string
	fs.BoolVar(&d, "d", false, "list defect, default story")
	fs.StringVar(&p, "p", "", "list all items of project p, default current project")
	parse(fs, args)
	if fs.NArg() == 1 {
		CurrentDir([]string{"rmigo", "cd", fs.Arg(0)})
	}
	if pkg.Conf.CurrentDir.Type != "sprint" {
		panic("invalid iteration")
	}
	var iteration = pkg.Conf.CurrentDir.ID
	client := pkg.NewRTCClient()
	if d {
		defects := client.ListDefectOfSprint(iteration)
		for i := range defects {
			fmt.Printf("%+v\n", defects[i])
		}
	} else {
		s := client.ListStoryOfSprint(iteration)
		for i := range s {
			fmt.Printf("%+v\n", s[i])
		}
	}
}

func Cat(args []string) {

}

func EditItem(args []string) {

}
func Search(args []string) {
	fs := flag.NewFlagSet("find", flag.ExitOnError)
	var a, i, p bool
	fs.BoolVar(&a, "a", false, "list all items including story and defect, default only story")
	fs.BoolVar(&i, "i", false, "list all items of iteration i, default current iteration")
	fs.BoolVar(&p, "p", false, "list all items of project p, default current project")
	parse(fs, args)
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

func parse(fs *flag.FlagSet, args []string) {
	if len(args) > 2 {
		args = args[2:]
	} else {
		args = []string{"-h"}
	}
	err := fs.Parse(args)
	if err != nil {
		panic(err)
	}
}
