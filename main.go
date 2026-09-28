package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

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

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	yaml_path := os.Args[0]

	// if scanner.Scan() {
	// 	yaml_path = scanner.Text()
	// }

	clean_path := path.Clean(yaml_path)
	job_map := read_yaml(&clean_path)

	if job_map["jobs"] != nil {
		fmt.Println("the jobs are:", job_map["jobs"])
	}

}
