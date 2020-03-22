package cmd

import (
	"github.com/sjeanpierre/git-tool/lib"
	"github.com/spf13/cobra"
	"log"
)

// prGet represents getting a PR object
var prGet = &cobra.Command{
	Use:   "get",
	Short: "Get a whole Github PR",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pr, _ := lib.GetPR(args[0])
		log.Printf("%s", pr)
	},
}

func init() {
	Cmd.AddCommand(prGet)
}
