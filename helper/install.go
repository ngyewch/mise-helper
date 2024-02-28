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

		listMissingResponse, err := miseDirHelper.ListMissing()
		if err != nil {
			return err
		}

		if listMissingResponse != nil {
			missingCount := 0
			for name, entries := range *listMissingResponse {
				for _, entry := range entries {
					fmt.Printf("[missing] %s %s\n", name, entry.RequestedVersion)
					missingCount++
				}
			}
			if missingCount > 0 {
				return fmt.Errorf("some tools were not installed")
			}
		}

		fmt.Println()
		return nil
	})
}
