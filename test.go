package main

import (
	"log"
	//"os"
	"os/exec"
	"runtime"
	"fmt"
)

func main() {
	cmd := exec.Command("ls")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}