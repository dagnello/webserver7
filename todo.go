package main

import "time"

type Todo struct {
    Id          int         `json:"id"`
    Name        string      `json:"name"`
    Completed   bool        `json:"completed"`
    Due         time.Time   `json:"due"`
}

type Todos []Todo

type Service struct {
    Image string `yaml:"image"`
 	Ports string `yaml:"ports"`
 	Links string `yaml:"links"`
}

type Network struct {
    Driver string `yaml:"driver"`
}

type Compose struct {
    Version string `yaml:"version"`
    Services map[string]Service `yaml:"services"`
    Networks map[string]Network `yaml:"networks"`
}