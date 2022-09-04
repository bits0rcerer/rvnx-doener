package data

import (
	"io/fs"
	"net/http"
	"strings"
)

type ServeFileSystemFS struct {
	fs.FS
	Root string
}

func (s *ServeFileSystemFS) Open(name string) (http.File, error) {
	subFS, _ := fs.Sub(fs.FS(s.FS), s.Root)
	file, err := http.FS(subFS).Open(name)

	if !strings.HasSuffix(name, ".html") && err != nil {
		// retry with html extension
		return http.FS(subFS).Open(name + ".html")
	}

	return file, err
}

func (s *ServeFileSystemFS) Exists(prefix string, filepath string) bool {
	if strings.HasSuffix(filepath, "/") {
		return s.Exists(prefix, filepath+"index.html")
	}

	if p := strings.TrimPrefix(filepath, prefix); len(p) <= len(filepath) {
		_, err := s.Open(p)
		if !strings.HasSuffix(filepath, ".html") && err != nil {
			// retry with html extension
			return s.Exists(prefix, filepath + ".html")
		}

		return err == nil
	}
	return false
}
