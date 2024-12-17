package main

import (
	"github.com/nzajk/password-manager/src/cmd"
)

func main() {
	cmd.RootCmd.AddCommand(cmd.LoginCmd)
	cmd.RootCmd.AddCommand(cmd.SaveCmd)
	cmd.RootCmd.AddCommand(cmd.GetCmd)
	cmd.RootCmd.Execute()
}
