package main

import (
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

const projectID string = "my-project-id"
const subID string = "message-sub"

func main() {
	ctx := context.Background()
	log.Print("SUBSUBSUB")
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	subs, err := list(client)
	for _, sub := range subs {
		fmt.Println(sub)
	}

	sub := client.Subscription(subID)
	log.Println(sub)
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		log.Printf("GOT MESSSGE %v\n", m)
		m.Ack()
	})
}

func list(client *pubsub.Client) ([]*pubsub.Subscription, error) {
	ctx := context.Background()
	// [START pubsub_list_subscriptions]
	var subs []*pubsub.Subscription
	it := client.Subscriptions(ctx)
	for {
		s, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}
	// [END pubsub_list_subscriptions]
	return subs, nil
}
