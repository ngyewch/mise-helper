package mise

import (
	"bufio"
	"bytes"
	"github.com/adrg/xdg"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var (
	legacyVersionFiles = map[string][]string{
		"crystal":   {".crystal-version"},
		"elixir":    {".exenv-version"},
		"go":        {".go-version", "go.mod"},
		"java":      {".java-version", ".sdkmanrc"},
		"node":      {".nvmrc", ".node-version"},
		"python":    {".python-version"},
		"ruby":      {".ruby-version", "Gemfile"},
		"terraform": {".terraform-version", ".packer-version", "main.tf"},
		"yarn":      {".yarnrc"},
	}
)

const (
	defaultConfigFilename        = ".mise.toml"
	defaultToolVersionsFilename  = ".tool-versions"
	defaultMiseLegacyVersionFile = "1"
)

type Helper struct {
	ConfigDir                     string
	GlobalConfigFile              string
	ConfigFilename                string
	ToolVersionsFilename          string
	LegacyVersionFile             bool
	LegacyVersionFileDisableTools []string
}

func getEnvWithDefault(name string, defaultValue string) string {
	value := os.Getenv(name)
	if value != "" {
		return value
	} else {
		return defaultValue
	}
}

func getEnvStringArrayWithDefault(name string, defaultValue []string) []string {
	value := os.Getenv(name)
	if value != "" {
		parts := strings.Split(value, ",")
		for i := 0; i < len(parts); i++ {
			parts[i] = strings.TrimSpace(parts[i])
		}
		return parts
	} else {
		return defaultValue
	}
}

func NewHelper() (*Helper, error) {
	configDir := getEnvWithDefault("MISE_CONFIG_DIR", path.Join(xdg.ConfigHome, "mise"))
	return &Helper{
		ConfigDir:                     configDir,
		GlobalConfigFile:              getEnvWithDefault("MISE_GLOBAL_CONFIG_FILE", path.Join(configDir, "config.toml")),
		ConfigFilename:                getEnvWithDefault("MISE_DEFAULT_CONFIG_FILENAME", defaultConfigFilename),
		ToolVersionsFilename:          getEnvWithDefault("MISE_DEFAULT_TOOL_VERSIONS_FILENAME", defaultToolVersionsFilename),
		LegacyVersionFile:             getEnvWithDefault("MISE_LEGACY_VERSION_FILE", defaultMiseLegacyVersionFile) == "1",
		LegacyVersionFileDisableTools: getEnvStringArrayWithDefault("MISE_LEGACY_VERSION_FILE_DISABLE_TOOLS", nil),
	}, nil
}

func (helper *Helper) HasVersionFiles(path string) (bool, error) {
	f, err := hasFile(path, helper.ConfigFilename)
	if err != nil {
		return false, err
	}
	if f {
		return true, nil
	}

	f, err = hasFile(path, helper.ToolVersionsFilename)
	if err != nil {
		return false, err
	}
	if f {
		return true, nil
	}

	if helper.LegacyVersionFile {
		for toolName, filenames := range legacyVersionFiles {
			disabled := false
			for _, disabledToolName := range helper.LegacyVersionFileDisableTools {
				if disabledToolName == toolName {
					disabled = true
					break
				}
			}
			if disabled {
				continue
			}
			f, err = hasAtLeastOneOfTheListedFiles(path, filenames)
			if err != nil {
				return false, err
			}
			if f {
				return true, nil
			}
		}
	}

	return false, nil
}

func (helper *Helper) ListAvailable(name string) ([]*ToolVersion, error) {
	cmd := exec.Command("mise", "ls-remote", name)
	buf := bytes.NewBuffer(nil)
	cmd.Stdin = os.Stdin
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	var versions []*ToolVersion
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		version := scanner.Text()
		// HACK special handling for go versions
		if strings.HasPrefix(name, "go:") {
			if strings.HasPrefix(version, "v") {
				version = version[1:]
			}
		}
		versions = append(versions, NewToolVersion(version))
	}
	return versions, nil
}

func hasFile(path string, filename string) (bool, error) {
	_, err := os.Stat(filepath.Join(path, filename))
	if err == nil {
		return true, nil
	} else if !os.IsNotExist(err) {
		return false, err
	} else {
		return false, nil
	}
}

func hasAtLeastOneOfTheListedFiles(path string, filenames []string) (bool, error) {
	for _, filename := range filenames {
		f, err := hasFile(path, filename)
		if err != nil {
			return false, err
		}
		if f {
			return true, nil
		}
	}
	return false, nil
}
