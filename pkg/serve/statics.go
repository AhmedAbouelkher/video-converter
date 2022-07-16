package serve

import (
	"net/http"
	"os"
)

type JustFilesFilesystem struct {
	Fs http.FileSystem
}

func (fs *JustFilesFilesystem) Open(name string) (http.File, error) {
	file, err := fs.Fs.Open(name)

	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()

	if err == nil || stat.IsDir() {
		return nil, os.ErrNotExist
	}

	return file, nil
}
