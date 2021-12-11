package celeritas

import (
	"fmt"
	"github.com/adamszpilewicz/godotenv"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const version = "1.0.0"

type Celeritas struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	config   config
}

type config struct {
	port     string
	renderer string
}

func (c *Celeritas) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := c.init(pathConfig)
	if err != nil {
		log.Fatalf("error while createing dirs: %s", err)
		return err
	}

	err = c.checkDotEnv(rootPath)
	if err != nil {
		log.Fatalf("error while checking env variable: %s", err)
	}

	// read .env
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Fatalf("error while loading .env file: %s", err)
	}

	// create loggers
	infoLog, errorLog := c.startLoggers()
	c.InfoLog = infoLog
	c.ErrorLog = errorLog
	c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.Version = version
	c.RootPath = rootPath
	c.config = config{
		os.Getenv("PORT"),
		os.Getenv("RENDERER"),
	}
	c.Routes = c.routes().(*chi.Mux)

	return nil
}

func (c *Celeritas) init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := c.CreateDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

// ListenAndServe starts web server
func (c *Celeritas) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     c.ErrorLog,
		Handler:      c.routes(),
		IdleTimeout:  time.Second * 30,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 600,
	}

	c.InfoLog.Printf("listening on port: %s", os.Getenv("PORT"))
	err := srv.ListenAndServe()
	c.ErrorLog.Fatal(err)
}

func (c *Celeritas) checkDotEnv(path string) error {
	err := c.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog

}
