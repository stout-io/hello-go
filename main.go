// Package main is a minimal, dependency-free Go program — a curated
// build-from-source sample (module at repo root) for Stout end-to-end validation.
package main

import "fmt"

// Greeting returns the canonical greeting.
func Greeting() string { return "hello from a Stout curated build" }

func main() { fmt.Println(Greeting()) }
