package cmd

import (
	"fmt"
	"os"

	"github.com/sepuka/chat/internal/def"

	"github.com/spf13/cobra"

	_ "github.com/sepuka/chat/internal/def/command"
)

var (
	configFile string

	rootCmd = &cobra.Command{
		Use:   "chat",
		Short: "Test telegram chat",
		Args:  cobra.MinimumNArgs(1),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return def.Build(configFile)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "/path/to/config.yml")
	_ = rootCmd.MarkPersistentFlagRequired("config")
}
