package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	flag.Parse()

	// 引数取得（コマンド名）
	args := flag.Args()
	if len(args) != 0 {
		commandName := args[0]
		if len(args) > 2 {
			fmt.Printf("引数は1つにして下さい。")
		}
	}
	
	
	// コマンド実行
	cmd, err := exec.Command(commandName).Output()
	fmt.Printf("command:\n%s :Error:\n%v\n", cmd, err)
}