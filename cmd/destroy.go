package cmd

import (
	"github.com/spf13/cobra"
)

// DestroyCmd ...
var DestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "TODO",
	Long:  "TODO",
}

func init() {
	RootCmd.AddCommand(DestroyCmd)
}
