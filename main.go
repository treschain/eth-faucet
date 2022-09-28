package main

import (
	"https://github.com/bladepool/eth-faucet/cmd"
)

//go:generate npm run build
func main() {
	cmd.Execute()
}
