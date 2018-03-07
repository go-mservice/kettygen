package main

import (
	"os"
	"os/exec"
	"fmt"
)

func goget(url string) {
	cmd := exec.Command("go", "get", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		println("Error:", err.Error())
	}
}

func main() {
	goget("github.com/golang/protobuf/proto")
	goget("github.com/yyzybb537/protoc-gen-ketty")

	args := []string{"--ketty_out=plugins=:."}

	gopath := os.Getenv("GOPATH")
	if gopath != "" {
		args = append([]string{"-I" + gopath + "/src/github.com/yyzybb537/ketty/pbinclude"}, args...)
		args = append([]string{"-I."}, args...)
	}

	args = append(args, os.Args[1:]...)
	cmd := exec.Command("protoc", args...)

	fmt.Printf("$ protoc")
	for _, arg := range args {
		fmt.Printf(" %s", arg)
    }
	fmt.Printf("\n")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		println("Error:", err.Error())
	}
}
