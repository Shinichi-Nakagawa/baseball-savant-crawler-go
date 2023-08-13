package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/Shinichi-Nakagawa/baseball-savant-crawler-go/gcp"
	"github.com/Shinichi-Nakagawa/baseball-savant-crawler-go/savant"
)

func main() {
	ctx := context.Background()
	gcs := gcp.NewStorageClient(ctx)
	bucket := gcp.GetBucket(gcs, GcsBucketName)
	client := gcp.NewPubSubClient(ctx, GoogleCloudProjectID)
	sub := gcp.Subscription(client, PubSubSubscriptionID)
	sub.ReceiveSettings.MaxOutstandingMessages = 10
	err := sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		form, err2 := savant.CreateForm(string(msg.Data))
		if err2 != nil {
			log.Print(err2)
			return
		}
		query := savant.Query(form)
		filename := savant.Filename(form)
		body, _ := savant.FetchAndAsString(query)
		gcp.WriteObject(bucket, fmt.Sprintf("%s/%s", GcsPathName, filename), body, ctx)
		log.Print(fmt.Sprintf("saved: %s", filename))
		msg.Ack()
	})
	if err != nil {
		log.Fatalf("sub.Receive: %s", err)
	}
}
