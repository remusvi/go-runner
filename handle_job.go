package main

import (
	"fmt"
	"log"
	"os/exec"
)

func handle_job(job *Job, platform Platform) bool {
	//fmt.Println("handling jobs")
	for _, v := range job.instructions {
		fmt.Printf("\033[32m %s \033[31m\n", v)
		cmd := exec.Command("/bin/zsh", "-c", v) //why no work
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Command execution failed: %v\nOutput: %s", err, string(output))
		}
		fmt.Println(string(output))
	}
	return false
}

func execute_windows() string {

}
