package main

import (
	"log"

	"github.com/fgrosse/servo"
	"github.com/fgrosse/servo-logxi"
	"github.com/fgrosse/servo-gorilla/routing"
	"github.com/fgrosse/servo/configuration"
)

func main() {
	loader := configuration.NewYAMLFileLoader("config/config.yml")
	kernel := servo.NewKernel(loader)
	kernel.Register(new(logxi.Bundle))
	kernel.Register(new(routing.Bundle))

	RegisterTypes(kernel.TypeRegistry)
	err := kernel.Run()
	if err != nil {
		log.Fatal(err)
	}
}
