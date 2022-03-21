package pubsub

import (
	"io"
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

func PullMsgs(w io.Writer, subID string) error {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("from pubsub.NewClient: %v", err)
	}

	defer client.Close()

	sub := client.Subscription(subID)

	// to simplify testing
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("from Receive: %v", err)
	}

	return nil
}
