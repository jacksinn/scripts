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
	out, err := exec.Command(cmdGit, cmdArgs...).Output()
	if(err != nil){
		fmt.Println("failed")
		fmt.Println(out)
		log.Fatal(err)
	}

	cmdArgs = []string{"commit", "-m", arg}
	out, err = exec.Command(cmdGit, cmdArgs...).Output()
	if(err != nil){
		fmt.Println("failed")
		fmt.Println(out)
		log.Fatal(err)
	}

	cmdArgs = []string{"push"}
	out, err = exec.Command(cmdGit, cmdArgs...).Output()
	if(err != nil){
		fmt.Println("failed")
		fmt.Println(out)
		log.Fatal(err)
	}
	sha := string(out)
	fmt.Println(sha)
}