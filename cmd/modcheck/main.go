package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/olekukonko/tablewriter"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/semver"
)

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	Name     string
	Version  string
	Latest   string
	Upgrade  bool
	Warnings string
}

func (d *Dependency) StringSlice() []string {
	var (
		version = d.Version
		name    = d.Name
	)

	// nicer word wrap for github markdown
	version = strings.ReplaceAll(version, "-", " ")
	version = strings.ReplaceAll(version, "+", " +")

	// strip github.com for less data
	name = strings.ReplaceAll(name, "github.com/", "")

	return toStringSlice(name, version, d.Latest, d.Warnings)
}

func load(gomodPath string) ([]*Dependency, error) {
	var result []*Dependency

	content, err := os.ReadFile(gomodPath)
	if err != nil {
		return nil, err
	}

	f, err := modfile.ParseLax(gomodPath, content, nil)
	if err != nil {
		return nil, err
	}

	// Exclude org-wide version checks
	//
	// The go module path gets stripped of the repository, so we can
	// avoid hitting proxy
	pkg := f.Module.Mod.String()
	pkg = path.Dir(pkg)

	for _, r := range f.Require {
		if r.Indirect {
			continue
		}

		dep := &Dependency{
			Name:    r.Mod.Path,
			Version: r.Mod.Version,
			Upgrade: false,
			Latest:  "Skipped",
		}

		if !strings.HasPrefix(dep.Name, pkg) {
			latest, err := getLatestVersion(dep.Name)
			if err != nil {
				dep.Latest = err.Error()
			} else {
				dep.Latest = latest
			}
		}

		dep.Warnings = lintImport(dep.Name, dep.Version, dep.Latest)

		switch {
		case dep.Latest == dep.Version:
			dep.Latest = "✓ Up to date"
		case strings.HasPrefix(dep.Latest, "bad request:"):
			dep.Latest = "✖ No info"
		default:
			if dep.Warnings == "" && semver.IsValid(dep.Version) && semver.IsValid(dep.Latest) {
				if semver.Compare(dep.Version, dep.Latest) < 0 {
					dep.Upgrade = true
				}
			}
		}

		result = append(result, dep)
	}

	return result, nil
}

func isSkipped(conf *options, name string) bool {
	for _, skipped := range conf.skip {
		if strings.HasPrefix(name, skipped) {
			return true
		}
	}
	return false
}

func start() error {
	conf := NewOptions()
	deps, err := load(conf.goModPath)
	if err != nil {
		return nil
	}

	switch {
	case conf.suggest:
		for _, dep := range deps {
			if dep.Upgrade {
				if isSkipped(conf, dep.Name) {
					log.Println(dep.Name, "held back from upgrade")
					continue
				}
				fmt.Printf("go get %s@%s\t\t# upgrade from %s\n", dep.Name, dep.Latest, dep.Version)
			}
		}
	default:
		output := &strings.Builder{}

		w := tablewriter.NewWriter(output)
		w.SetHeader([]string{"import", "version", "latest", "warnings"})
		w.SetAutoWrapText(false)
		w.SetAutoFormatHeaders(true)
		w.SetTablePadding(" ")
		w.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		w.SetAlignment(tablewriter.ALIGN_LEFT)
		w.SetRowSeparator("")
		w.SetHeaderLine(true)
		w.SetCenterSeparator("|")
		w.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})

		for _, dep := range deps {
			if isSkipped(conf, dep.Name) {
				dep.Warnings = "Held back from upgrade"
			}

			if conf.forUpgrade && !dep.Upgrade {
				continue
			}

			w.Append(dep.StringSlice())
		}

		w.Render()

		tableString := output.String()
		for strings.Contains(tableString, "||") {
			tableString = strings.Replace(tableString, "||", "|:---|", 1)
		}

		fmt.Print(tableString)
	}

	return nil
}

func getLatestVersion(name string) (string, error) {
	var result string

	res, err := http.Get("https://proxy.golang.org/" + url.QueryEscape(strings.ToLower(name)) + "/@v/list")
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	parts := strings.Split(strings.TrimSpace(string(body)), "\n")
	cleanParts := []string{}
	for _, part := range parts {
		// Skip `-rc`, `-dev` and similar suffixes
		if strings.Contains(part, "-") {
			continue
		}
		cleanParts = append(cleanParts, part)
	}

	semver.Sort(cleanParts)

	if len(cleanParts) > 0 {
		result = cleanParts[len(cleanParts)-1]
	}
	if len(result) == 0 {
		err = errors.New("No versions available")
	}

	return result, err
}

func toStringSlice(vars ...string) []string {
	return vars
}

func lintImport(name, version, latest string) string {
	if strings.Contains(name, "gopkg.in/") {
		return "Deprecated import (gopkg.in)"
	}

	if strings.HasPrefix(version, "v0.0.0-") {
		return "Dependency without go.mod"
	}

	if strings.HasPrefix(latest, "bad request:") {
		return "Bad request, possibly renamed"
	}

	if latest == "Skipped" {
		return ""
	}

	versionTrimmed := strings.Split(version, "-")[0]
	if semver.Compare(versionTrimmed, latest) > 0 {
		return "Version ahead of latest release"
	}

	return ""
}
