package main

import (
	"testing"

	"github.com/hauntedness/rtigo/pkg"
)

func TestCurrentDir(t *testing.T) {
	pkg.InitConfig("user1", "123456")
	pkg.LoadConfig()
	var args = []string{"./rmigo.exe", "cd", "~/p1/sprint11"}
	CurrentDir(args)
}

func TestCurrentDir2(t *testing.T) {
	pkg.LoadConfig()
	var args = []string{"./rmigo.exe", "cd", ".."}
	CurrentDir(args)
}
