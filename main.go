package main

import (
	"context"
	"fmt"

	"github.com/shurcool/githubv4"
	"golang.org/x/oauth2"
)

var query struct {
	Viewer struct {
		Name      githubv4.String
		Login     githubv4.String
		CreatedAt githubv4.DateTime
	}
}

func main() {
	fmt.Println("hello world.")
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("name:", query.Viewer.Name)
	fmt.Println("login:", query.Viewer.Login)
	fmt.Println("created:", query.Viewer.CreatedAt)
}
