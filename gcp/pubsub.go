package gcp

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
)

func NewPubSubClient(ctx context.Context, projectId string) *pubsub.Client {
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("pubsub.NewClient: %v", err)
	}
	return client
}

func Subscription(client *pubsub.Client, subID string) *pubsub.Subscription {
	sub := client.Subscription(subID)
	return sub
}
