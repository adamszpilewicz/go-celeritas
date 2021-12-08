package celeritas

import "log"

const version = "1.0.0"

type Celeritas struct {
	AppName string
	Debug bool
	Version string
}

func (c *Celeritas) New(rootPath string) error{
	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := c.init(pathConfig)
	if err != nil {
		log.Fatalf("error while createing dirs: %s", err)
		return err
	}
	return nil
}

func (c *Celeritas) init(p initPaths) error {
	root := p.rootPath
	for _, path:= range p.folderNames {
		err := c.CreateDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}
