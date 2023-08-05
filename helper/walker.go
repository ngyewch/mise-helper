package helper

import (
	"github.com/denormal/go-gitignore"
	"github.com/ngyewch/rtx-helper/rtx"
	"os"
	"path/filepath"
)

func walkDirectories(handler func(rtxHelper *rtx.Helper, path string) error) error {
	ignore, err := gitignore.NewRepository(".")
	if err != nil {
		return err
	}

	helper, err := rtx.NewHelper()
	if err != nil {
		return err
	}

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		match := ignore.Relative(path, info.IsDir())
		if match != nil {
			if match.Ignore() {
				if info.IsDir() {
					return filepath.SkipDir
				} else {
					return nil
				}
			}
		}
		if info.IsDir() {
			f, err := hasConfigFiles(path)
			if err != nil {
				return err
			}
			if f {
				err = handler(helper, path)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func hasConfigFiles(path string) (bool, error) {
	_, err := os.Stat(filepath.Join(path, ".rtx.toml"))
	if err == nil {
		return true, nil
	} else if !os.IsNotExist(err) {
		return false, err
	}
	_, err = os.Stat(filepath.Join(path, ".tool-versions"))
	if err == nil {
		return true, nil
	} else if !os.IsNotExist(err) {
		return false, err
	}
	return false, nil
}
