package main

import (
	"./cmd"
)

//go:generate npm run build
func main() {
	cmd.Execute()
}
