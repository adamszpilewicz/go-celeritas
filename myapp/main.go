package main

import "github.com/adamszpilewicz/go-laravel/celeritas"

type application struct {
	App *celeritas.Celeritas
}

func main() {
	initApplication()
}