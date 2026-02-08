package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run scripts/update_version.go <new_version>")
		fmt.Println("Example: go run scripts/update_version.go 1.0.2")
		os.Exit(1)
	}

	newVersion := os.Args[1]
	// Basic validation: x.y.z
	matched, _ := regexp.MatchString(`^\d+\.\d+\.\d+$`, newVersion)
	if !matched {
		fmt.Printf("Error: Version '%s' does not match format x.y.z\n", newVersion)
		os.Exit(1)
	}

	fmt.Printf("Updating files to version %s...\n", newVersion)

	updateGlobalGo(newVersion)
	updateManifest(newVersion)
	updateInstaller(newVersion)
	updateChangelog(newVersion)

	fmt.Println("Done! Please verify changes and commit.")
}

func updateGlobalGo(ver string) {
	path := "cmd/main.go"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", path, err)
		return
	}

	// const Version = "1.0.1"
	re := regexp.MustCompile(`const Version = ".*"`)
	newContent := re.ReplaceAllString(string(content), fmt.Sprintf(`const Version = "%s"`, ver))

	err = ioutil.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		fmt.Printf("Error writing %s: %v\n", path, err)
	} else {
		fmt.Printf("Updated %s\n", path)
	}
}

func updateManifest(ver string) {
	path := "wincu.manifest"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", path, err)
		return
	}

	// version="1.0.1.0"
	re := regexp.MustCompile(`version="[\d\.]+"`)
	newContent := re.ReplaceAllString(string(content), fmt.Sprintf(`version="%s.0"`, ver))

	err = ioutil.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		fmt.Printf("Error writing %s: %v\n", path, err)
	} else {
		fmt.Printf("Updated %s\n", path)
	}
}

func updateInstaller(ver string) {
	path := "installer.iss"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", path, err)
		return
	}

	// AppVersion=1.0.1
	re := regexp.MustCompile(`AppVersion=[\d\.]+`)
	newContent := re.ReplaceAllString(string(content), fmt.Sprintf(`AppVersion=%s`, ver))

	// OutputBaseFilename=wincu_installer -> wincu-installer-v1.0.1
	reFilename := regexp.MustCompile(`OutputBaseFilename=.*`)
	newContent = reFilename.ReplaceAllString(newContent, fmt.Sprintf(`OutputBaseFilename=wincu-installer-v%s`, ver))

	err = ioutil.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		fmt.Printf("Error writing %s: %v\n", path, err)
	} else {
		fmt.Printf("Updated %s\n", path)
	}
}

func updateChangelog(ver string) {
	path := "CHANGELOG.md"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", path, err)
		return
	}

	// Check if this version already exists
	if strings.Contains(string(content), fmt.Sprintf("[v%s]", ver)) {
		fmt.Printf("Changelog already contains v%s, skipping add.\n", ver)
		return
	}

	// Prepend new version section
	// Expected format: ## [vNewVer] - Unreleased
	re := regexp.MustCompile(`## \[v`)
	loc := re.FindStringIndex(string(content))

	newHeader := fmt.Sprintf("## [v%s] - Unreleased\n\n- No changes yet.\n\n", ver)

	var newContent string
	if loc != nil {
		// Insert before the first version header
		start := loc[0]
		newContent = string(content[:start]) + newHeader + string(content[start:])
	} else {
		// Appends to top if no version header found (unlikely but safe fallback)
		newContent = newHeader + string(content)
	}

	err = ioutil.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		fmt.Printf("Error writing %s: %v\n", path, err)
	} else {
		fmt.Printf("Updated %s\n", path)
	}
}
