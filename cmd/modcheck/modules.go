package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"golang.org/x/mod/semver"
)

//go:embed modules.json
var vulnsBytes []byte

type Module struct {
	Path  string
	Vulns []*Vuln
}

func (m *Module) String(version string) string {
	result := []string{}
	for _, v := range m.Vulns {
		if v.Fixed != "" {
			versionTrimmed := strings.Split(version, "-")[0]
			fixedTrimmed := strings.Split(v.Fixed, "-")[0]
			if semver.Compare(versionTrimmed, fixedTrimmed) < 0 {
				result = append(result, v.ID+" fixed in "+v.Fixed)
				continue
			}
			continue
		}
		result = append(result, v.ID)
	}
	if len(result) == 0 {
		return fmt.Sprintf("0 of %d", len(m.Vulns))
	}

	return strings.Join(result, ", ")
}

type Vuln struct {
	ID       string     `json:"id"`
	Fixed    string     `json:"fixed"`
	Modified *time.Time `json:"modified"`
}

type ModuleList []*Module

func (m ModuleList) Find(name string) *Module {
	for _, v := range m {
		if v.Path == name {
			return v
		}
	}
	return nil
}

func getVulns() (ModuleList, error) {
	result := ModuleList{}
	err := json.Unmarshal(vulnsBytes, &result)
	return result, err
}
