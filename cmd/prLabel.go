package cmd

import (
	"../lib"
	"context"
	"github.com/spf13/cobra"
	"log"
)

// labelCmd represents the adding of labels command
var labelCmd = &cobra.Command{
	Use:   "label",
	Short: "Add a label to a Github PR",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		prLabel(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(labelCmd)
}

func prLabel(prLink, label string) {
	owner, repo, number, err := lib.ParsePRLink(prLink)
	if err != nil {
		log.Fatalf("Could not parse PR Link, Error: %s", err.Error())
	}
	client := lib.NewClient()
	labels, _, err := client.Issues.AddLabelsToIssue(context.Background(), owner, repo, number, []string{label})
	if err != nil {
		log.Fatalf("Encountered error pushing label %s to PR %s, Error: %s",label,prLink,err.Error())
	}
	log.Printf("Successfully pushed label to PR, labels present are %s",labels)
}
