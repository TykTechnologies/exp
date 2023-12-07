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

func start() error {
	var (
		gomodPath = "go.mod"
	)

	vulns, err := getVulns()
	if err != nil {
		return err
	}

	content, err := os.ReadFile(gomodPath)
	if err != nil {
		return err
	}

	f, err := modfile.ParseLax(gomodPath, content, nil)
	if err != nil {
		return err
	}

	// Exclude org-wide version checks
	//
	// The go module path gets stripped of the repository, so we can
	// avoid hitting proxy
	pkg := f.Module.Mod.String()
	pkg = path.Dir(pkg)

	output := &strings.Builder{}

	w := tablewriter.NewWriter(output)
	w.SetHeader([]string{"import", "version", "latest", "warnings", "cves"})
	w.SetAutoWrapText(false)
	w.SetAutoFormatHeaders(true)
	w.SetTablePadding(" ")
	w.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	w.SetAlignment(tablewriter.ALIGN_LEFT)
	w.SetRowSeparator("")
	w.SetHeaderLine(true)
	w.SetCenterSeparator("|")
	//	w.EnableBorder(true)
	w.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})

	for _, r := range f.Require {
		if r.Indirect {
			continue
		}

		var err error
		name, version, latest := r.Mod.Path, r.Mod.Version, "Skipped"

		if !strings.HasPrefix(name, pkg) {
			latest, err = getLatestVersion(name)
			if err != nil {
				latest = err.Error()
			}
		}

		warnings := lintImport(name, version, latest)
		var cves string
		if m := vulns.Find(name); m != nil {
			cves = m.String(version)
		}

		if latest == version {
			latest = "✓ Up to date"
		}
		if strings.HasPrefix(latest, "bad request:") {
			latest = "✖ No info"
		}

		// nicer word wrap for github markdown
		version = strings.ReplaceAll(version, "-", " ")
		version = strings.ReplaceAll(version, "+", " +")

		// strip github.com for less data
		name = strings.ReplaceAll(name, "github.com/", "")

		w.Append(toStringSlice(name, version, latest, warnings, cves))
	}

	w.Render()

	tableString := output.String()
	for strings.Contains(tableString, "||") {
		tableString = strings.Replace(tableString, "||", "|:---|", 1)
	}

	fmt.Print(tableString)

	return nil
}

func getLatestVersion(name string) (string, error) {
	var result string

	res, err := http.Get("https://proxy.golang.org/" + url.QueryEscape(name) + "/@v/list")
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	parts := strings.Split(strings.TrimSpace(string(body)), "\n")

	semver.Sort(parts)

	if len(parts) > 0 {
		result = parts[len(parts)-1]
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
