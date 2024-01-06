package mise

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type DirHelper struct {
	path string
}

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

func NewDirHelper(path string) *DirHelper {
	return &DirHelper{
		path: path,
	}
}

func (helper *DirHelper) InstallAll() error {
	cmd := exec.Command("mise", "install")
	cmd.Dir = helper.path
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

func (helper *DirHelper) ListInstalled() (*ListAllResponse, error) {
	cmd := exec.Command("mise", "list", "--json")
	cmd.Dir = helper.path
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
