package cmd

import (
	"../lib"
	"context"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"log"
)

// commentCmd represents the adding of comment command
var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Add a comment to a Github PR",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		prComment(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(commentCmd)
}

func prComment(prLink, body string) {
	owner, repo, number, err := lib.ParsePRLink(prLink)
	if err != nil {
		log.Fatalf("Could not parse PR Link, Error: %s", err.Error())
	}
	client := lib.NewClient()
	comment := &github.IssueComment{
		Body:           &body,
	}
	comm, _, err := client.Issues.CreateComment(context.Background(), owner, repo, number, comment)
	if err != nil {
		log.Fatalf("Encountered error posting comment to PR %s - %s - %d, error: %s", owner, repo, number, err.Error())
	}
	log.Printf("Comment '%s' added to PR %s",*comm.Body, *comm.HTMLURL)
}
