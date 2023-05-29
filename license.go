package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"text/template"
)

type License struct {
	longName string
	template string
}

var licenses = map[string]License{
	"agpl-3.0":     {"GNU Affero General Public License v3.0", Agpl30Template},
	"apache-2.0":   {"Apache License 2.0", Apache20Template},
	"bsd-2-clause": {"BSD 2-Clause \"Simplified\" License", Bsd2ClauseTemplate},
	"bsd-3-clause": {"BSD 3-Clause \"New\" or \"Revised\" License", Bsd3ClauseTemplate},
	"cc0-1.0":      {"Creative Commons Zero v1.0 Universal", Cc010Template},
	"epl-2.0":      {"Eclipse Public License 2.0", Epl20Template},
	"free-art-1.3": {"Free Art License 1.3", FreeArt13Template},
	"gpl-2.0":      {"GNU General Public License v2.0", Gpl20Template},
	"gpl-3.0":      {"GNU General Public License v3.0", Gpl30Template},
	"lgpl-2.1":     {"GNU Lesser General Public License v2.1", Lgpl21Template},
	"lgpl-3.0":     {"GNU Lesser General Public License v3.0", Lgpl30Template},
	"mit":          {"MIT License", MitTemplate},
	"mpl-2.0":      {"Mozilla Public License 2.0", Mpl20Template},
	"unlicense":    {"The Unlicense", UnlicenseTemplate},
	"wtfpl":        {"Do What The Fuck You Want To Public License", WtfplTemplate},
	"resource":     {"Resource License", RlTemplate},
}

func getKeys(table map[string]License) []string {
	keys := make([]string, 0, len(licenses))

	for key := range licenses {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func printList() {
	keys := getKeys(licenses)

	for _, key := range keys {
		stdout.Printf("%-14s(%s)", key, licenses[key].longName)
	}
}

func printLicense(license, output, name, year string) {
	file, ok := licenses[license]
	if !ok {
		if match := findLicense(license); match == "" {
			userInputError.Abort(fmt.Errorf("unknown license %w\nrun \"license -list\" for list of available licenses", license))
		} else {
			file, _ = licenses[match]
		}
	}

	t, err := template.New("license").Parse(file.template)
	if err != nil {
		internalError.Abort(fmt.Errorf("internal: failed to parse license template for %w", license))
	}

	var outFile io.Writer = os.Stdout
	if output != "" {
		f, err := os.Create(filepath.Clean(output))
		if err != nil {
			internalError.Abort(fmt.Errorf("failed to create file %s: %w", output, err))
		}
		outFile = f
	}

	if err := t.Execute(outFile, struct {
		Name string
		Year string
	}{name, year}); err != nil {
		internalError.Abort(err)
	}

	noError.Abort(nil)
}

// find the most recent license that is a regex match for the argument
func findLicense(pattern string) string {
	keys := getKeys(licenses)
	sort.Slice(keys, func(i, j int) bool { return keys[i] > keys[j] })
	for _, key := range keys {
		if match, err := regexp.MatchString(pattern, key); err == nil && match {
			return key
		}
	}
	return ""
}
