package main

import "github.com/romainDavaze/nagiosxi-cli/cmd"

func main() {
	cmd.GenMarkdownDocs()
	cmd.Execute()
}
