package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/sjeanpierre/git-tool/lib"
	"github.com/spf13/cobra"
)

// prGet represents getting a PR object
var prGet = &cobra.Command{
	Use:   "get",
	Short: "Get a whole Github PR",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pr, _ := lib.GetPR(args[0])
		prJson, _ := json.Marshal(pr)
		fmt.Println(string(prJson))
	},
}

func init() {
	rootCmd.AddCommand(prGet)
}
