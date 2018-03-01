package main

import (
	"os"
	"os/exec"
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
	args = append(args, os.Args[1:]...)
	cmd := exec.Command("protoc", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		println("Error:", err.Error())
	}
}
