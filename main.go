package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"gopkg.in/yaml.v3"
)

type Job struct {
	name         string
	instructions []string
}

func read_yaml(yaml_path *string) map[string]any {
	yamlfile, err := os.ReadFile(*yaml_path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var m map[string]any
	if err := yaml.Unmarshal(yamlfile, &m); err != nil {
		panic(err)
	}
	return m
}

func job_builder(job_map map[string]any) []Job {
	job_arr := []Job{}

	for _, v := range job_map["jobs"].([]any) {
		executable_job := Job{name: fmt.Sprint(v)}
		if job_map[executable_job.name] != nil {
			for _, v := range job_map[executable_job.name].([]any) {
				executable_job.instructions = append(executable_job.instructions, fmt.Sprint(v))
				//fmt.Println(executable_job)
			}
		}
		job_arr = append(job_arr, executable_job)
	}
	return job_arr
}

// func handle_job(job *Job) bool {
// 	//fmt.Println("handling jobs")
// 	for _, v := range job.instructions {
// 		fmt.Printf("\033[32m %s \033[31m\n", v)
// 		cmd := exec.Command("/bin/zsh", "-c", v) //why no work
// 		output, err := cmd.CombinedOutput()
// 		if err != nil {
// 			log.Fatalf("Command execution failed: %v\nOutput: %s", err, string(output))
// 		}
// 		fmt.Println(string(output))
// 	}
// 	return false
// }

func main() {
	yaml_path := os.Args[1]
	clean_path := path.Clean(yaml_path)
	job_map := read_yaml(&clean_path)
	curr_os := runtime.GOOS
	platform := PlatformNone

	if job_map["jobs"] != nil {
		job_arr := job_builder(job_map)
		switch curr_os {
		case "windows":
			platform = Windows
			for _, v := range job_arr {
				handle_job(&v, platform)
			}
		case "darwin":
			platform = Darwin
			for _, v := range job_arr {
				handle_job(&v, platform)
			}
		}
	} else {
		fmt.Printf("missing tasks")
	}
}
