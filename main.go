package main

import "go-api-todolist/config"

var cfg config.Config

func init() {
	if err := config.ParseEnv(&cfg); err != nil {
		panic(err)
	}
}

func main() {

}
