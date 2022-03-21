package pubsub

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/pubsub"
)

const projectID = "project-competition-344803"
func Publish(w io.Writer, topicID, msg string) error {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("from pubsub.NewClient: %v", err)
	}

	defer client.Close()

	tp := client.Topic(topicID)
	result := tp.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("from Get: %v", err)
	}

	fmt.Fprintf(w, "Published a message - msg ID: %v\n", id)
	return nil
}
