package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Person struct {
	Name    string   `yaml:"name"`
	Age     int      `yaml:"age"`
	Hobbies []string `yaml:"hobbies"`
}

func main() {
	p := Person{
		Name:    "Alex",
		Age:     8,
		Hobbies: []string{"Art", "Sport", "TV"},
	}

	toYaml(p)
	fmt.Println("---------------")
	fromYaml("Nick.yaml")

}

func toYaml(p Person) {
	fmt.Printf("%v >>> to yaml\n", p)
	y, err := yaml.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(y))
}

func fromYaml(file string) {
	fmt.Printf("%s from yaml to object\n", file)
	b, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	p := &Person{}

	err = yaml.Unmarshal([]byte(b), &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*p)

}
