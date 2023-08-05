package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToolVersion(t *testing.T) {
	{
		toolVersion := NewToolVersion("adoptopenjdk-19.0.0-beta+36.0.202208190932")
		expectToolVersion(t, toolVersion, true, "adoptopenjdk", "19.0.0-beta+36.0.202208190932", "beta")
	}
	{
		toolVersion := NewToolVersion("adoptopenjdk-19.0.1+10")
		expectToolVersion(t, toolVersion, true, "adoptopenjdk", "19.0.1+10", "")
	}
	{
		toolVersion := NewToolVersion("adoptopenjdk-openj9-8.0.192+12.OpenJDK8U-jdk_x64_linux_openj9_8u192b12.tar.gz")
		expectToolVersion(t, toolVersion, false, "adoptopenjdk-openj9", "8.0.192+12.OpenJDK8U-jdk_x64_linux_openj9_8u192b12.tar.gz", "")
	}
	{
		toolVersion := NewToolVersion("adoptopenjdk-openj9-8.0.192+12.openj9-0.11.0")
		expectToolVersion(t, toolVersion, true, "adoptopenjdk-openj9", "8.0.192+12.openj9-0.11.0", "")
	}
	{
		toolVersion := NewToolVersion("7.27.0-0")
		expectToolVersion(t, toolVersion, true, "", "7.27.0-0", "0")
	}
}

func expectToolVersion(t *testing.T, toolVersion *ToolVersion, expectedValid bool, expectedVersionPrefix string, expectedVersionNumber string, expectedPrerelease string) {
	assert.Equalf(t, expectedValid, toolVersion.Valid(), `toolVersion.VersionPrefix = %v, expected = %v`,
		toolVersion.Valid(), expectedValid)
	assert.Equalf(t, expectedVersionPrefix, toolVersion.VersionPrefix,
		`toolVersion.VersionPrefix = "%s", expected = "%s"`, toolVersion.VersionPrefix, expectedVersionPrefix)
	assert.Equalf(t, expectedVersionNumber, toolVersion.VersionNumber,
		`toolVersion.VersionNumber = "%s", expected = "%s"`, toolVersion.VersionNumber, expectedVersionNumber)
	if toolVersion.Valid() {
		assert.Equalf(t, expectedPrerelease, toolVersion.Prerelease(),
			`toolVersion.Prerelease() = "%s", expected = "%s"`, toolVersion.Prerelease(), expectedPrerelease)
	}
}
