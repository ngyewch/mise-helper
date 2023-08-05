package helper

import (
	"fmt"
	"github.com/ngyewch/rtx-helper/rtx"
	"path/filepath"
)

func Latest(hideLatest bool, includePrereleases bool) error {
	return walkDirectories(func(rtxHelper *rtx.Helper, path string) error {
		fmt.Println(path)
		absPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		response, err := rtxHelper.ListInstalled(path)
		if err != nil {
			return err
		}
		for toolName, listings := range *response {
			var matchedListing *rtx.Listing
			for _, listing := range listings {
				if listing.Source != nil {
					dir := filepath.Dir(listing.Source.Path)
					if dir == absPath {
						matchedListing = listing
						break
					}
				}
			}
			if matchedListing != nil {
				fmt.Printf("- %s@%s\n", toolName, matchedListing.RequestedVersion)
				versions, err := rtxHelper.ListAvailable(toolName)
				if err != nil {
					return err
				}
				fmt.Printf("  - %v\n", versions)
			}
		}
		fmt.Println()
		return nil
	})
}
