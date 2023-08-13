package main

import "os"

// Project Environment
var GcsBucketName string = os.Getenv("GCS_BUCKET_NAME")
var GcsPathName string = os.Getenv("GCS_PATH_NAME")
var GoogleCloudProjectID string = os.Getenv("GOOGLE_CLOUD_PROJECT")
var PubSubSubscriptionID string = os.Getenv("PUBSUB_SUBSCRIPTION_ID")
