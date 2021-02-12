package context

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// CollectStaticFiles collects all of the files in the `static/` directory relative to cwd
// and generates a map of their relative paths
func CollectStaticFiles(cwd string) (map[string]os.File, error) {
	staticDir := filepath.Join(cwd, "static")

	stat, err := os.Stat(staticDir)
	if err != nil {
		if err == os.ErrNotExist {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to Stat static directory")
	} else if !stat.IsDir() {
		return nil, errors.New("'static' is not a directory")
	}

	files := map[string]os.File{}

	filepath.Walk(staticDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return errors.Wrap(err, "failed to Open file: "+path)
		}

		relativePath := strings.TrimPrefix(path, staticDir+"/") // don't want to leave the leading slash

		files[relativePath] = *file

		return nil
	})

	return files, nil
}
