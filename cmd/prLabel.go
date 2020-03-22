package cmd

import (
	"github.com/sjeanpierre/git-tool/lib"
	"github.com/spf13/cobra"
)

// labelCmd represents the adding of labels command
var labelCmd = &cobra.Command{
	Use:   "label",
	Short: "Add a label to a Github PR",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		lib.AddPRLabel(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(labelCmd)
}


