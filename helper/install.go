package helper

import (
	"fmt"
	"github.com/ngyewch/rtx-helper/rtx"
)

func Install(recursive bool) error {
	return walkDirectories(recursive, func(rtxHelper *rtx.Helper, rtxDirHelper *rtx.DirHelper, path string) error {
		fmt.Println(path)
		err := rtxDirHelper.InstallAll()
		if err != nil {
			return err
		}
		fmt.Println()
		return nil
	})
}
