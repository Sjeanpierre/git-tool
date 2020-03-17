package lib

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"os"
	"strconv"
	"strings"
)

var (
	githubToken = os.Getenv("GITHUB_API_TOKEN")
)

//NewClient returns a GH service client used to make request to Github's API
func NewClient() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client
}

func ParsePRLink(prLink string) (owner,repo string, prNumber int, err error) {
	linkParts := strings.Split(strings.Replace(prLink, "https://github.com/", "", 1), "/")
	owner = linkParts[0]
	repo = linkParts[1]
	prNumber, err = strconv.Atoi(linkParts[len(linkParts)-1])
	if err != nil {
		return "","",0, errors.New("Could not parse PR number from URL")
	}
	return
}
