package helper

import (
	"fmt"
	"github.com/ngyewch/mise-helper/mise"
)

func Install(recursive bool) error {
	return walkDirectories(recursive, func(miseHelper *mise.Helper, miseDirHelper *mise.DirHelper, path string) error {
		fmt.Println(path)
		err := miseDirHelper.InstallAll()
		if err != nil {
			return err
		}
		fmt.Println()
		return nil
	})
}
