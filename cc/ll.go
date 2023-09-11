package main

import (
	"fmt"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("ls")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
}
