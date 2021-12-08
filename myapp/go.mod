module myapp

go 1.17

replace github.com/adamszpilewicz/go-laravel/celeritas => ../celeritas

replace github.com/adamszpilewicz/godotenv => ../../godotenv

require github.com/adamszpilewicz/go-laravel/celeritas v0.0.0-00010101000000-000000000000

require github.com/adamszpilewicz/godotenv v0.0.0-00010101000000-000000000000 // indirect
