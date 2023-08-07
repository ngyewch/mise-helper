package helper

import (
	"fmt"
	"github.com/ngyewch/rtx-helper/rtx"
)

func Install() error {
	return walkDirectories(func(rtxHelper *rtx.Helper, rtxDirHelper *rtx.DirHelper, path string) error {
		fmt.Println(path)
		err := rtxDirHelper.InstallAll()
		if err != nil {
			return err
		}
		fmt.Println()
		return nil
	})
}
