package cmd

import (
	"errors"

	"github.com/sepuka/chat/src/telegram"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(telegramCmd)
}

const (
	repeaterMode = `repeater`
	hostingMode  = `hosting`
)

var (
	mode = `repeater`

	telegramCmd = &cobra.Command{
		Use:     `telegram`,
		Example: `telegram -c /config/path`,
		Short:   `hosting telegram controller`,
		Long:    `Manages hosting staff via telegram`,

		RunE: func(cmd *cobra.Command, args []string) error {
			switch mode {
			case repeaterMode:
				return telegram.Repeater()
			case hostingMode:
				return telegram.Hosting()
			default:
				return errors.New(`unknown mode`)
			}

		},
	}
)

func init() {
	telegramCmd.Flags().StringVar(&mode, "mode", "", "--mode=repeater")
}
