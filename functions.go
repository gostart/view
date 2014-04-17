package view

import (
	"time"

	"github.com/ungerik/go-dry"
)

func FindStaticFile(filename string) (filePath string, found bool, modified time.Time) {
	// todo optimize
	return dry.FileFindModified(append(Config.BaseDirs, Config.StaticDirs...), filename)
}

func FindTemplateFile(filename string) (filePath string, found bool, modified time.Time) {
	// todo optimize
	return dry.FileFindModified(append(Config.BaseDirs, Config.StaticDirs...), filename)
}
