package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l", "/Users/liuqingzheng/")
	err := cmd.Run()
	if err != nil {
		fmt.Println("执行命令出错", err)
	}
}
