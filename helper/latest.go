package helper

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/ngyewch/rtx-helper/rtx"
	"github.com/pelletier/go-toml"
	"os"
	"path/filepath"
)

type Config struct {
	Constraints map[string]string
}

func Latest(hideLatest bool, includePrereleases bool, recursive bool) error {
	return walkDirectories(recursive, func(rtxHelper *rtx.Helper, rtxDirHelper *rtx.DirHelper, path string) error {
		fmt.Println(path)
		absPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		config, err := ReadConfig(filepath.Join(path, ".rtx-helper.toml"))
		if err != nil {
			return err
		}
		response, err := rtxDirHelper.ListInstalled()
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
				toolVersion := rtx.NewToolVersion(matchedListing.RequestedVersion)
				availableVersions, err := rtxHelper.ListAvailable(toolName)
				if err != nil {
					return err
				}
				if toolVersion.Valid() {
					matchedAvailableVersion := false
					var latestToolVersion *rtx.ToolVersion
					for _, availableVersion := range availableVersions {
						if matchedListing.RequestedVersion == availableVersion.Version {
							matchedAvailableVersion = true
						}
						if availableVersion.Valid() && (toolVersion.VersionPrefix == availableVersion.VersionPrefix) && (includePrereleases || availableVersion.SemVer.Prerelease() == "") {
							var constraints *semver.Constraints
							if config != nil && config.Constraints != nil {
								constraintsStr, ok := config.Constraints[toolName]
								if ok {
									constraints, err = semver.NewConstraint(constraintsStr)
									if err != nil {
										return err
									}
								}
							}
							if (constraints == nil) || constraints.Check(availableVersion.SemVer) {
								if (latestToolVersion == nil) || availableVersion.SemVer.GreaterThan(latestToolVersion.SemVer) {
									latestToolVersion = availableVersion
								}
							}
						}
					}
					if !matchedAvailableVersion {
						if latestToolVersion != nil {
							fmt.Printf("- %s@%s [REQUESTED UNKNOWN] (latest %s)\n", toolName, matchedListing.RequestedVersion, latestToolVersion.Version)
						} else {
							fmt.Printf("- %s@%s [REQUESTED UNKNOWN]\n", toolName, matchedListing.RequestedVersion)
						}
					} else if latestToolVersion != nil {
						if !hideLatest || latestToolVersion.SemVer.GreaterThan(toolVersion.SemVer) {
							fmt.Printf("- %s@%s (latest %s)\n", toolName, matchedListing.RequestedVersion, latestToolVersion.Version)
						}
					} else {
						fmt.Printf("- %s@%s (LATEST UNKNOWN)\n", toolName, matchedListing.RequestedVersion)
					}
				} else {
					fmt.Printf("- %s@%s (skipped, invalid semver)\n", toolName, matchedListing.RequestedVersion)
				}
			}
		}
		fmt.Println()
		return nil
	})
}

func ReadConfig(path string) (*Config, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	configBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = toml.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
