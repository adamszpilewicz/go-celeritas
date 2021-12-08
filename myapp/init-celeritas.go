package main

import (
	"github.com/adamszpilewicz/go-laravel/celeritas"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "myapp"
	cel.Debug = true

	app := &application{
		App: cel,
	}

	return app

}