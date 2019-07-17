package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var (
	startCmd = &cobra.Command{
		Use:     `start`,
		Example: `/start`,
		Short:   `command invokes help tutorial`,
		Long:    `
list
create [opts]
delete [opts]
`,
	}
)
