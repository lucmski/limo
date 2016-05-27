package cmd

import (
	"fmt"

	"github.com/hoop33/limo/config"
	"github.com/spf13/cobra"
)

// VersionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	Long:  fmt.Sprintf("Display version information for %s", config.ProgramName),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.Version)
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
