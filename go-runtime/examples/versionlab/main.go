package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

//go:embed loop.go.txt
var source []byte

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	dir, err := os.MkdirTemp("", "versionlab-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	if err := os.WriteFile(filepath.Join(dir, "main.go"), source, 0600); err != nil {
		return err
	}

	// Reuse the toolchain that built this runner, even if another go is on PATH.
	goCommand := filepath.Join(runtime.GOROOT(), "bin", "go")
	if runtime.GOOS == "windows" {
		goCommand += ".exe"
	}
	for _, test := range []struct {
		version string
		want    string
	}{
		{"1.21", "closures: [3 3 3]\npointers: [3 3 3]\npointers equal: true\nassignment closures: [3 3 3]\nshared backing array: [99 99 99]\n"},
		{"1.22", "closures: [1 2 3]\npointers: [1 2 3]\npointers equal: false\nassignment closures: [3 3 3]\nshared backing array: [99 99 99]\n"},
	} {
		mod := "module example.com/loopfixture\n\ngo " + test.version + "\n"
		if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(mod), 0600); err != nil {
			return err
		}
		cmd := exec.Command(goCommand, "run", ".")
		cmd.Dir = dir
		cmd.Env = cleanEnv()
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("go %s fixture: %w\n%s", test.version, err, output)
		}
		if string(output) != test.want {
			return fmt.Errorf("go %s fixture mismatch\nwant:\n%sgot:\n%s", test.version, test.want, output)
		}
		fmt.Printf("go.mod: go %s\n%s\n", test.version, output)
	}
	return nil
}

func cleanEnv() []string {
	var env []string
	for _, entry := range os.Environ() {
		name, _, _ := strings.Cut(entry, "=")
		switch name {
		case "GOTOOLCHAIN", "GOWORK", "GOEXPERIMENT", "GOFLAGS":
			continue
		}
		env = append(env, entry)
	}
	return append(env, "GOTOOLCHAIN=local", "GOWORK=off", "GOEXPERIMENT=", "GOFLAGS=")
}
