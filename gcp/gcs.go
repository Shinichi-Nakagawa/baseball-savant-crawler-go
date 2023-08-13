package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"log"
)

func NewStorageClient(ctx context.Context) *storage.Client {
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}
	return client
}

func GetBucket(client *storage.Client, name string) *storage.BucketHandle {
	bucket := client.Bucket(name)
	return bucket

}

func WriteObject(bkt *storage.BucketHandle, name string, value string, ctx context.Context) {
	obj := bkt.Object(name)
	w := obj.NewWriter(ctx)
	if _, err := fmt.Fprintf(w, value); err != nil {
		log.Fatalf("fmt.Fprintf: %v", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("w.Close: %v", err)
	}
}
