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
	out, err := exec.Command(cmdGit, "add .").Output()
	if(err != nil){
		log.Fatal(err)
	}
	sha := string(out)
	fmt.Println(sha)
}