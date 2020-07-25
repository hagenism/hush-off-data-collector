package main

import (
	"log"
	"os"

	appsync "github.com/sony/appsync-client-go"
	"github.com/sony/appsync-client-go/graphql"
)

const (
	gqlAPIEnvKey = "gql.api"
)

func main() {

	// Appsync Client (Subscribe)
	gqlapi := os.Getenv(gqlAPIEnvKey)
	client := appsync.NewClient(appsync.NewGraphQLClient(graphql.NewClient(gqlapi)))
	subscription := `subscription SubscribeToEcho() { subscribeToEcho }`
	response, err := client.Post(graphql.PostRequest{
		Query: subscription,
	})
	if err != nil {
		log.Fatal(err)
	}

	ext, err := appsync.NewExtensions(response)
	if err != nil {
		log.Fatalln(err)
	}

	ch := make(chan *graphql.Response)
	subscriber := appsync.NewSubscriber(*ext,
		func(r *graphql.Response) { ch <- r },
		func(err error) { log.Println(err) },
	)

	if err := subscriber.Start(); err != nil {
		log.Fatalln(err)
	}
	defer subscriber.Stop()

	// Logic
}
