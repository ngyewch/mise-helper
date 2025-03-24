package helper

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/ngyewch/mise-helper/mise"
	"github.com/pelletier/go-toml"
	"os"
	"path/filepath"
)

type Config struct {
	Constraints map[string]string
}

type LatestHandler struct {
	global             bool
	hideLatest         bool
	includePrereleases bool
}

func (handler *LatestHandler) Handle(miseHelper *mise.Helper, miseDirHelper *mise.DirHelper, path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	if handler.global {
		fmt.Println(miseHelper.GlobalConfigFile)
	} else {
		fmt.Println(path)
	}
	config, err := ReadConfig(filepath.Join(path, ".mise-helper.toml"))
	if err != nil {
		return err
	}
	response, err := miseDirHelper.ListInstalled()
	if err != nil {
		return err
	}
	for toolName, listings := range *response {
		var matchedListing *mise.Listing
		for _, listing := range listings {
			if listing.Source != nil {
				if handler.global {
					if (path == ".") && (listing.Source.Path == miseHelper.GlobalConfigFile) {
						matchedListing = listing
						break
					}
				} else {
					dir := filepath.Dir(listing.Source.Path)
					if dir == absPath {
						matchedListing = listing
						break
					}
				}
			}
		}
		if matchedListing != nil {
			toolVersion := mise.NewToolVersion(matchedListing.RequestedVersion)
			availableVersions, err := miseHelper.ListAvailable(toolName)
			if err != nil {
				return err
			}
			if toolVersion.Valid() {
				matchedAvailableVersion := false
				var latestToolVersion *mise.ToolVersion
				for _, availableVersion := range availableVersions {
					if matchedListing.RequestedVersion == availableVersion.Version {
						matchedAvailableVersion = true
					}
					if availableVersion.Valid() && (toolVersion.VersionPrefix == availableVersion.VersionPrefix) && (handler.includePrereleases || availableVersion.SemVer.Prerelease() == "") {
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
					if !handler.hideLatest || latestToolVersion.SemVer.GreaterThan(toolVersion.SemVer) {
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
}

func Latest(hideLatest bool, includePrereleases bool, recursive bool, includeGlobal bool) error {
	if includeGlobal {
		globalHandler := &LatestHandler{
			global:             true,
			hideLatest:         hideLatest,
			includePrereleases: includePrereleases,
		}
		err := walkDirectories(recursive, globalHandler.Handle)
		if err != nil {
			return err
		}
	}
	localHandler := &LatestHandler{
		global:             false,
		hideLatest:         hideLatest,
		includePrereleases: includePrereleases,
	}
	return walkDirectories(recursive, localHandler.Handle)
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
