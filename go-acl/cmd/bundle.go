package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func bundle() {

}

func main() {
	// Parse command line arguments
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <dstBundledPath> <solveCodePath>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "example: %s ./bundle/abc123/abc123.go ./abc123/abc123.go\n", os.Args[0])
		os.Exit(1)
	}
	dstBundledPath, solveCodePath := args[0], args[1]
	fmt.Printf("dstBundledPath: %s\n", dstBundledPath)
	fmt.Printf("solveCodePath: %s\n", solveCodePath)

	// Create the destination directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(dstBundledPath), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create directory: %v\n", err)
		os.Exit(1)
	}
	bundledAcl, err := os.CreateTemp("", "bundle-acl*.go")
	if err != nil {
		panic(err)
	}
	fmt.Printf("bundledAcl: %s\n", bundledAcl.Name())
	defer os.Remove(bundledAcl.Name())

	// Bundle acl package to bundledAcl
	cmd := exec.Command(
		"go", "run",
		"golang.org/x/tools/cmd/bundle@latest",
		"-o", bundledAcl.Name(),
		".",
	)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "bundle failed: %v\nstderr: %s\n", err, stderr.String())
	}

}
