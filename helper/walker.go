package helper

import (
	"github.com/denormal/go-gitignore"
	"github.com/ngyewch/rtx-helper/rtx"
	"os"
	"path/filepath"
)

func walkDirectories(recursive bool, handler func(rtxHelper *rtx.Helper, rtxDirHelper *rtx.DirHelper, path string) error) error {
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
		if !recursive && (path != ".") && info.IsDir() {
			return filepath.SkipDir
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
			f, err := helper.HasVersionFiles(path)
			if err != nil {
				return err
			}
			if f {
				dirHelper := rtx.NewDirHelper(path)
				err = handler(helper, dirHelper, path)
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
