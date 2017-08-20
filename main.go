package main

import (
	"github.com/google/go-github/github"
	"fmt"
	"context"
)

func main() {
	ctx := context.Background()
	client := github.NewClient(nil)

	// list all organizations for user "willnorris"
	orgs, _, err := client.Organizations.List(ctx, "mislav", nil)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	for _, element := range orgs {
		fmt.Printf("%v\n", *element.Login)
	}
}

