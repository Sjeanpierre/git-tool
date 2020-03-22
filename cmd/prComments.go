package cmd

import (
	"github.com/sjeanpierre/git-tool/lib"
	"github.com/spf13/cobra"
	"strings"
)

// commentCmd represents the adding of comment command
var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Add a comment to a Github PR",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		lib.AddPRComment(args[0], strings.Join(args[1:],""))
	},
}

func init() {
	rootCmd.AddCommand(commentCmd)
}


