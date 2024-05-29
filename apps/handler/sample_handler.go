package handler

import (
	"context"
	"log/slog"

	"cloud.google.com/go/pubsub"
	"github.com/sdual/pubsub-sample/apps/usecase"
)

type PubsubHandler struct {
	usecase usecase.SampleUsecase
}

func (p PubsubHandler) CallBack(ctx context.Context, msg *pubsub.Message) {
	completed, err := p.usecase.Subscriber(ctx)
	if err != nil {
		slog.Error("failed", slog.Any("error", err))
	}
	ackResult := p.ackResult(completed, msg)
	p.logging(ctx, ackResult)
}

func (p PubsubHandler) ackResult(completed bool, msg *pubsub.Message) *pubsub.AckResult {
	if completed {
		return msg.AckWithResult()
	}
	return msg.NackWithResult()
}

func (p PubsubHandler) logging(ctx context.Context, ackResult *pubsub.AckResult) error {
	status, err := ackResult.Get(ctx)
	if err != nil {
		return err
	}
	slog.Info("pubsub ack result", slog.Any("status", status))
	return nil
}
