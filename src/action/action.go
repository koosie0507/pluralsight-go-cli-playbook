package main

import (
	"flag"
	"log"
	"plugin"
)

func main() {
	path := flag.String("plugin", "", "Execute this plugin")
	flag.Parse()

	if *path == "" {
		log.Fatal("Please provide path to plugin")
	}

	p, err := plugin.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

	f, err := p.Lookup("ThingToDo")
	if err != nil {
		log.Fatal(err)
	}
	thingToDo, ok := f.(func())
	if !ok {
		log.Fatal("not ok")
	}
	thingToDo()
	log.Println("amazing")
}
