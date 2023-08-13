# Baseball Savant Crawler

[Baseball Savant(Statcast Search)]() Crawler for Go

## Requirement

Go latest version

## Installation

```bash
$ go get -u .
```

## Setting

### Environment

```shell
export GOOGLE_CLOUD_PROJECT=your-project-id
export GCS_BUCKET_NAME=your-storage-bucket
export GCS_PATH_NAME=your-storage-path
export PUBSUB_SUBSCRIPTION_ID=your-cloud-pubsub-subxcription-id
```

## Usage

### Running

```bash
$ go build .
$ go run .
```

### Test

```bash
$ go test -v ./...
```

## Use

### sample request

```json
{
  "season": 2023,
  "player_type": "batter",
  "game_date": "2023-08-10T00:00:00Z"
}
```

### Publisher sample

```go
package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"log"
	"time"
)

type PlayerType string

const (
	PITCHER PlayerType = "pitcher"
	BATTER  PlayerType = "batter"
)

type Form struct {
	Season     int        `validate:"required, min=2015,max=2999" json:"season"`
	PlayerType PlayerType `validate:"required" json:"player_type"`
	GameDate   time.Time  `validate:"required" json:"game_date"`
}

func Publish(ctx context.Context, topic *pubsub.Topic, form Form) {
	value, err := json.Marshal(form)
	if err != nil {
		log.Printf("Request Error: %s", err)
	}
	result := topic.Publish(ctx, &pubsub.Message{
		Data: value,
	})
	id, err := result.Get(ctx)
	if err != nil {
		log.Printf("Publish Error: %s", err)
	}
	log.Printf("Published a message; msg ID: %v\n", id)
}

func NewPubSubClient(ctx context.Context, projectId string) *pubsub.Client {
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("pubsub.NewClient: %v", err)
	}
	return client
}

func Topic(client *pubsub.Client, topicID string) *pubsub.Topic {
	topic := client.Topic(topicID)
	return topic
}

func main() {
	t := time.Now().UTC()
	gameDate := t.Add(DATE_DIFF * time.Hour * 24)
	ctx := context.Background()
	client := NewPubSubClient(ctx, GoogleCloudProjectID)
	topicExporter := Topic(client, PubTopicIDExporter)
	formBatter := Form{Season: Season, GameDate: gameDate, PlayerType: BATTER}
	log.Printf("export batter game_date: %s", formBatter.GameDate)
	Publish(ctx, topicExporter, formBatter)
	log.Print("export batter end")
	formPitcher := Form{Season: Season, GameDate: gameDate, PlayerType: PITCHER}
	log.Printf("export pitcher game_date: %s", formPitcher.GameDate)
	Publish(ctx, topicExporter, formPitcher)
	log.Print("export pitcher end")
}

```
