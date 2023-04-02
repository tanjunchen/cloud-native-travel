package main

import (
	"context"
	"fmt"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"log"
)

const token = "be22f4726892ae65f21f3d8f3832540c157e993f"

func main() {

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	type (
		OrganizationFragment struct {
			Description string
		}
		UserFragment struct {
			Bio string
		}
	)

	var q struct {
		RepositoryOwner struct {
			Login                string
			OrganizationFragment `graphql:"... on Organization"`
			UserFragment         `graphql:"... on User"`
		} `graphql:"repositoryOwner(login: \"istio\")"`
	}
	err := client.Query(context.Background(), &q, nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(q.RepositoryOwner.Login)
	fmt.Println(q.RepositoryOwner.Description)
	fmt.Println(q.RepositoryOwner.Bio)

}
