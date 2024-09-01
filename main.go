/*
Package main is the entry point of the application. It sets up version information and executes the commands defined in the cmd package.

Copyright © 2024 Guilherme Thomazi Bonicontro <thomazi@linux.com>
*/
package main

import "github.com/guitmz/capable/cmd"

// Version information for the application.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// main function sets version information and executes the commands.
func main() {
	cmd.SetVersionInfo(version, commit, date)
	cmd.Execute()
}
