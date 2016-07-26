package main

import (
  "fmt"
  "os"
  "os/exec"
  "log"
)

func main() {
	arg := "Default"
	if(len(os.Args) > 1){
		arg = os.Args[1]
	}
	fmt.Println(arg)
	cmdGit := "git"
	cmdArgs := []string{"add", "."}
	tryExec(cmdGit, cmdArgs)

	cmdArgs = []string{"commit", "-m", arg}
	tryExec(cmdGit, cmdArgs)

	cmdArgs = []string{"push"}
	tryExec(cmdGit, cmdArgs)

	sha := string(out)
	fmt.Println(sha)
}

func tryExec(cmd []string, cmdArgs []string){
		out, err := exec.Command(cmdGit, cmdArgs...).Output()
	if(err != nil){
		fmt.Println("failed")
		fmt.Println(out)
		log.Fatal(err)
	}
}