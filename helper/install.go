package helper

import (
	"fmt"
	"github.com/ngyewch/rtx-helper/rtx"
)

func Install() error {
	return walkDirectories(func(rtxHelper *rtx.Helper, path string) error {
		fmt.Println(path)
		err := rtxHelper.InstallAll(path)
		if err != nil {
			return err
		}
		fmt.Println()
		return nil
	})
}
