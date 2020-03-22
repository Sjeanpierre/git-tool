package lib

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
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

func ParsePRLink(prLink string) (owner, repo string, prNumber int, err error) {
	linkParts := strings.Split(strings.Replace(prLink, "https://github.com/", "", 1), "/")
	owner = linkParts[0]
	repo = linkParts[1]
	prNumber, err = strconv.Atoi(linkParts[len(linkParts)-1])
	if err != nil {
		return "", "", 0, fmt.Errorf("could not parse PR number from URL %s", err.Error())
	}
	return
}

func AddPRComment(prLink, body string) (ic github.IssueComment, e error) {
	owner, repo, number, err := ParsePRLink(prLink)
	if err != nil {
		return ic, fmt.Errorf("could not parse PR Link, Error: %s", err.Error())
	}
	client := NewClient()
	comment := &github.IssueComment{
		Body: &body,
	}
	comm, _, err := client.Issues.CreateComment(context.Background(), owner, repo, number, comment)
	if err != nil {
		return ic, fmt.Errorf("encountered error posting comment to PR %s - %s - %d, error: %s", owner, repo, number, err.Error())
	}
	log.Printf("comment '%s' added to PR %s", *comm.Body, *comm.HTMLURL)
	return *comm, nil
}

func AddPRLabel(prLink, label string) (l []*github.Label, e error) {
	owner, repo, number, err := ParsePRLink(prLink)
	if err != nil {
		return l, fmt.Errorf("could not parse PR Link, Error: %s", err.Error())
	}
	client := NewClient()
	labels, _, err := client.Issues.AddLabelsToIssue(context.Background(), owner, repo, number, []string{label})
	if err != nil {
		return l, fmt.Errorf("encountered error pushing label %s to PR %s, Error: %s", label, prLink, err.Error())
	}
	log.Printf("Successfully pushed label to PR, labels present are %s", labels)
	return labels, nil
}

func RemovePRLabel(prLink, label string) ( e error) {
	owner, repo, number, err := ParsePRLink(prLink)
	if err != nil {
		return fmt.Errorf("could not parse PR Link, Error: %s", err.Error())
	}
	client := NewClient()
	_, err = client.Issues.RemoveLabelForIssue(context.Background(), owner, repo, number, label)
	if err != nil {
		return fmt.Errorf("encountered error pushing label %s to PR %s, Error: %s", label, prLink, err.Error())
	}
	return nil
}

func GetPRComments(prURL string) (result []*github.IssueComment, e error) {
	log.Printf("Retrieving comments for pull request %s", prURL)
	c := NewClient()
	owner, repo, prNumber, err := ParsePRLink(prURL)
	if err != nil {
		return result, fmt.Errorf("could not retrieve comments from URL %s encountered error %s", prURL, err.Error())
	}
	comments, resp, err := c.Issues.ListComments(context.Background(), owner, repo, prNumber, &github.IssueListCommentsOptions{})
	if err != nil {
		return result, fmt.Errorf("error in Github request to retrieve comments from %s, error: %s-%s", prURL, err.Error(), resp)
	}
	return comments, nil
}

func GetPRsWithLabel(label string) (result github.IssuesSearchResult, e error) {
	log.Printf("Searching for pull requests with label %s", label)
	c := NewClient()
	q := fmt.Sprintf("is:open is:pr label:\"%s\"", label)
	searchResults, _, err := c.Search.Issues(context.Background(), q, &github.SearchOptions{})
	if err != nil {
		return result, fmt.Errorf("could not perform search, Error %s", err.Error())
	}
	return *searchResults, nil
}

func GetPR(prURL string) (i github.PullRequest, e error) {
	owner, repo, number, err := ParsePRLink(prURL)
	if err != nil {
		return i, fmt.Errorf("could not parse PR link, error %s", err.Error())
	}
	c := NewClient()
	pr, _, err := c.PullRequests.Get(context.Background(), owner, repo, number)
	if err != nil {
		return i, fmt.Errorf("encountered error retrieving PR %s, error %s", prURL, err.Error())
	}
	return *pr, nil
}

func GetPRComment(commentURL string) (ic github.IssueComment, e error) {
	commentURLParts := strings.Split(commentURL, "#")
	if len(commentURLParts) < 2 {
		return ic, fmt.Errorf("comment URL is invalid %s", commentURL)
	}
	commentID := strings.Replace(commentURLParts[1], "issuecomment-", "", 1)
	owner, repo, number, err := ParsePRLink(commentURLParts[0])
	if err != nil {
		return ic, fmt.Errorf("could not parse PR Link, Error: %s", err.Error())
	}
	cid, err := strconv.Atoi(commentID)
	if err != nil {
		fmt.Errorf("comment id invalid, cannot be coverted to int")
	}
	client := NewClient()
	comm, _, err := client.Issues.GetComment(context.Background(), owner, repo, cid)
	if err != nil {
		return ic, fmt.Errorf("encountered error posting comment to PR %s - %s - %d, error: %s", owner, repo, number, err.Error())
	}
	return *comm, nil
}

func GetPRReviews(prURL string) (i []*github.PullRequestReview, e error) {
	owner, repo, number, err := ParsePRLink(prURL)
	if err != nil {
		return i, fmt.Errorf("could not parse PR link, error %s", err.Error())
	}
	c := NewClient()
	reviews, _, err := c.PullRequests.ListReviews(context.Background(), owner, repo, number, &github.ListOptions{}) //Get(context.Background(),owner,repo,number)
	if err != nil {
		return i, fmt.Errorf("could not retrieve PR reviews, error %s", err.Error())
	}
	return reviews, nil
}
