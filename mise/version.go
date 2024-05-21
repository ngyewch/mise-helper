package mise

import (
	"github.com/Masterminds/semver/v3"
	"regexp"
)

var (
	versionPrefixLocatorRegex1 = regexp.MustCompile(`^[0-9]`)
	versionPrefixLocatorRegex2 = regexp.MustCompile(`-[0-9]`)
)

type ToolVersion struct {
	Version       string
	VersionPrefix string
	VersionNumber string
	SemVer        *semver.Version
}

func NewToolVersion(version string) *ToolVersion {
	versionPrefixLocation := versionPrefixLocatorRegex1.FindStringIndex(version)
	if versionPrefixLocation != nil {
		v, _ := semver.NewVersion(version)
		return &ToolVersion{
			Version:       version,
			VersionPrefix: "",
			VersionNumber: version,
			SemVer:        v,
		}
	}
	versionPrefixLocation = versionPrefixLocatorRegex2.FindStringIndex(version)
	if versionPrefixLocation != nil {
		versionPrefix := version[0:versionPrefixLocation[0]]
		versionNumber := version[versionPrefixLocation[0]+1:]
		v, _ := semver.NewVersion(versionNumber)
		return &ToolVersion{
			Version:       version,
			VersionPrefix: versionPrefix,
			VersionNumber: versionNumber,
			SemVer:        v,
		}
	}
	return &ToolVersion{
		Version:       version,
		VersionPrefix: "",
		VersionNumber: "",
		SemVer:        nil,
	}
}

func (v *ToolVersion) Valid() bool {
	return v.SemVer != nil
}

func (v *ToolVersion) Prerelease() string {
	return v.SemVer.Prerelease()
}

func (v *ToolVersion) CheckConstraints(c *semver.Constraints) bool {
	if v.SemVer == nil {
		return false
	}
	if c != nil {
		return c.Check(v.SemVer)
	} else {
		return true
	}
}
