package rtx

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type ListAllResponse map[string][]*Listing

type Listing struct {
	Version          string  `json:"version"`
	RequestedVersion string  `json:"requested_version,omitempty"`
	InstallPath      string  `json:"install_path,omitempty"`
	Source           *Source `json:"source,omitempty"`
}

type Source struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type Helper struct {
}

func NewHelper() (*Helper, error) {
	return &Helper{}, nil
}

func (helper *Helper) InstallAll(path string) error {
	cmd := exec.Command("rtx", "install")
	cmd.Dir = path
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if (cmd.ProcessState != nil) && (cmd.ProcessState.ExitCode() != 0) {
		fmt.Printf("exit code = %d\n", cmd.ProcessState.ExitCode())
	} else {
		if err != nil {
			return err
		}
	}
	return nil
}

func (helper *Helper) ListInstalled(path string) (*ListAllResponse, error) {
	cmd := exec.Command("rtx", "list", "--json")
	cmd.Dir = path
	buf := bytes.NewBuffer(nil)
	cmd.Stdin = os.Stdin
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if (cmd.ProcessState != nil) && (cmd.ProcessState.ExitCode() != 0) {
		fmt.Printf("exit code = %d\n", cmd.ProcessState.ExitCode())
	} else {
		if err != nil {
			return nil, err
		}
	}
	var result ListAllResponse
	err = json.Unmarshal(buf.Bytes(), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (helper *Helper) ListAvailable(name string) ([]string, error) {
	cmd := exec.Command("rtx", "ls-remote", name)
	buf := bytes.NewBuffer(nil)
	cmd.Stdin = os.Stdin
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if (cmd.ProcessState != nil) && (cmd.ProcessState.ExitCode() != 0) {
		fmt.Printf("exit code = %d\n", cmd.ProcessState.ExitCode())
	} else {
		if err != nil {
			return nil, err
		}
	}
	var versions []string
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		version := scanner.Text()
		versions = append(versions, version)
	}
	return versions, nil
}
