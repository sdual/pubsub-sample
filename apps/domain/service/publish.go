package service

import (
	"context"
	"fmt"
	"sync"

	"cloud.google.com/go/pubsub"
)

type MsgPublish struct {
	client *pubsub.Client
	topic  *pubsub.Topic
	wg     sync.WaitGroup
}

func NewMsgPublish(ctx context.Context, projectID, topicID string) (*MsgPublish, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %w", err)
	}
	topic := client.Topic(topicID)
	return &MsgPublish{
		client: client,
		topic:  topic,
	}, nil
}

func (mp *MsgPublish) Publish(ctx context.Context, msg string) {
	result := mp.topic.Publish(ctx, &pubsub.Message{Data: []byte(msg)})
	go func(ctx context.Context, r *pubsub.PublishResult) {
		id, err := r.Get(ctx)
		if err != nil {
			fmt.Printf("failed to publish. id: %s, %v", id, err)
		}
	}(ctx, result)
}

func (mp *MsgPublish) Stop() {
	mp.topic.Stop()
}
