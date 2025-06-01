# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based AtCoder Competitive Programming Library (ACL) containing data structures and algorithms commonly used in competitive programming contests. The main components include:

- **Core Library** (`lib.go`): Input/output utilities, array manipulation functions, and common helper functions optimized for competitive programming
- **Data Structures**: Priority queues (`pq.go`), double-ended priority queues (`depq.go`), segment trees (`segtree.go`), union-find (`unionfind.go`), splay trees (`splay.go`)  
- **Algorithms**: Mathematical functions (`math.go`), string algorithms (`strings.go`), graph algorithms (`graph.go`), sieve algorithms (`sieve.go`)
- **Verification System**: Online judge verification tests in `verify/` directory using a shell script system

## Development Commands

- **Run tests**: `go test -v` (runs all unit tests including fuzz tests)
- **Run single test**: `go test -v -run TestName`
- **Build**: `go build` or `go run .` 
- **Run main program**: `go run .` (expects input from stdin for competitive programming problem)

## Architecture

The codebase follows a competitive programming template pattern:

- `main.go` contains the solution logic for a specific problem (currently a pathfinding problem)
- `lib.go` provides core I/O functions (`I()`, `II()`, `S()`, `Ans()`) and utility functions optimized for contest environments
- Individual files contain specialized data structures and algorithms as standalone modules
- `testlib/assert.go` provides testing utilities using google/go-cmp for assertions
- Each module has corresponding `*_test.go` files with comprehensive unit and fuzz tests

## Verification System

The `verify/` directory contains online judge verification using:
- `new.sh` script to create new verification tests by copying `main.go` 
- Each subdirectory represents a problem with `verify.test.go` containing the solution
- Verification comments format: `// verification-helper: PROBLEM <url>`

## Key Design Patterns

- Generic types using Go constraints for data structures (e.g., `Ordered` interface)
- Functional programming patterns with `Map`, `Reduce`, `Filter` functions
- Memory-optimized I/O using buffered readers/writers for competitive programming
- Comprehensive fuzz testing for mathematical and algorithmic correctness