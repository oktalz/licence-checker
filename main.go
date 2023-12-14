package main

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"slices"
	"sort"

	"gopkg.in/yaml.v3"
)

type Package struct {
	Type    string `yaml:"license"`
	Package string `yaml:"package"`
}
type Group struct {
	Types    []string  `yaml:"type"`
	Packages []Package `yaml:"packages"`
}
type GroupBlocked struct {
	Types    []string `yaml:"type"`
	Packages []string `yaml:"packages"`
}

type Config struct {
	Allowed Group        `yaml:"allowed"`
	Blocked GroupBlocked `yaml:"blocked"`
}

func main() {
	cfgFilePath := os.Getenv("LICENSE_CHECKER")
	if cfgFilePath == "" {
		cfgFilePath = ".license-checker.yml"
	}

	cfg := Config{}

	cfgFileData, err := os.ReadFile(cfgFilePath)
	if err != nil {
		cfg.Allowed = Group{}
		cfg.Allowed.Packages = []Package{}
		cfg.Allowed.Types = []string{}
	} else {
		cfg, err = parseConfig(cfgFileData)
		if err != nil {
			slog.Error("", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}

	cmd := exec.Command("go-licenses", "report", "./...")

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err = cmd.Run()
	if err != nil {
		slog.Error("", slog.String("error", err.Error()))
		os.Exit(1)
	}

	stdoutData := stdoutBuf.Bytes()
	// stderrData := stderrBuf.Bytes()

	licences, err := check(stdoutData, cfg)
	if err != nil {
		slog.Error("", slog.String("error", err.Error()))
		os.Exit(1)
	}
	yamlData, err := yaml.Marshal(licences)
	if err != nil {
		slog.Error("Failed to marshal licenses to YAML", slog.String("error", err.Error()))
		os.Exit(1)
	}
	fmt.Println(string(yamlData)) //nolint:forbidigo
}

func parseConfig(cfgFileData []byte) (Config, error) {
	cfg := Config{}

	err := yaml.Unmarshal(cfgFileData, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("error parsing config file: %w", err)
	}

	return cfg, nil
}

type Error struct {
	Msg    string `json:"msg,omitempty" yaml:"msg,omitempty"`
	Reason string `json:"reason"        yaml:"reason"`
}

type Licenses struct {
	Packages []string `json:"packages" yaml:"packages"`
	Total    int      `json:"total"    yaml:"total"`
}

type CheckResult struct {
	Licenses map[string]Licenses `json:"licenses"         yaml:"licenses"`
	Errors   []Error             `json:"errors,omitempty" yaml:"errors,omitempty"`
}

func check(stdoutData []byte, cfg Config) (CheckResult, error) { //nolint:funlen
	result := CheckResult{
		Licenses: map[string]Licenses{},
		Errors:   []Error{},
	}

	lines := bytes.Split(stdoutData, []byte("\n"))
	for _, line := range lines {
		if len(line) > 0 {
			parts := bytes.Split(line, []byte(","))
			pkg := result.Licenses[string(parts[2])]
			pkg.Packages = append(pkg.Packages, string(parts[0]))
			result.Licenses[string(parts[2])] = pkg
		}
	}
	licTypes := []string{}
	for k := range result.Licenses {
		licTypes = append(licTypes, k)
	}
	sort.Strings(licTypes)
	HasError := false
	for _, licType := range licTypes {
		if !slices.Contains(cfg.Allowed.Types, licType) {
			// slog.Error("", slog.String("msg", "License type not allowed"), slog.String("type", licType))
			for _, pkg := range result.Licenses[licType].Packages {
				found := false
				for _, allowedPackage := range cfg.Allowed.Packages {
					if pkg == allowedPackage.Package && licType == allowedPackage.Type {
						found = true
						// slog.Info("   ", slog.String("msg", "Package in allowed list"), slog.String("package", pkg))
					}
				}
				if !found {
					// slog.Warn("    ", slog.String("type", licType), slog.String("package", pkg))
					result.Errors = append(result.Errors, Error{
						Reason: fmt.Sprintf("License type %s is not allowed", licType),
						Msg:    fmt.Sprintf("Package %s is not allowed", pkg),
					})
					HasError = true
				}
			}

			continue
		}
		// slog.Info("", slog.String("type", licType), slog.Int("packages", len(licenses[licType])))
		lt := result.Licenses[licType]
		lt.Total = len(lt.Packages)
		result.Licenses[licType] = lt
		for _, pkg := range result.Licenses[licType].Packages {
			if slices.Contains(cfg.Blocked.Packages, pkg) {
				// slog.Error("", slog.String("msg", "Package not allowed"), slog.String("package", pkg))
				result.Errors = append(result.Errors, Error{
					Reason: fmt.Sprintf("Package %s is not allowed", pkg),
				})
				HasError = true
			}
			// slog.Warn("    ", slog.String("type", licType), slog.String("package", pkg))
		}
	}
	if HasError {
		return result, fmt.Errorf("license check failed")
	}

	return result, nil
}
