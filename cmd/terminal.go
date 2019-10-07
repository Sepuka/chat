package cmd

import (
	"chat/src/context"
	"chat/src/def"
	"chat/src/def/source"
	commandSource "chat/src/source"
	"github.com/spf13/cobra"
)

func init() {
	terminalCmd.PersistentFlags().StringVar(&login, `login`, ``, `user's login'`)
	rootCmd.AddCommand(terminalCmd)
}

var (
	login       string
	terminalCmd = &cobra.Command{
		Use:     `terminal`,
		Example: `terminal -c /config/path [list|create]`,
		Short:   `hosting terminal controller`,
		Long:    `Manages hosting staff via terminal`,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			commandSourceListener, err := def.Container.SafeGet(source.TerminalDef)
			if err != nil {
				return err
			}

			req := context.NewRequest(login, args[0], args[1:]...)

			return commandSourceListener.(*commandSource.Terminal).Execute(req)
		},
	}
)
