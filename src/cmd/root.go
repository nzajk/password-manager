package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "password-manager",
	Short: "A simple password manager.",
}
