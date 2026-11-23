package main

import (
	"log"
	"os/exec"
)

func main() {
	err := exec.Command("wscript", "login.vbs").Run()
	if err != nil {
		log.Fatal(err)
	}
}
