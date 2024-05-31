package usecase

import (
	"context"

	"github.com/sdual/pubsub-sample/apps/domain/service"
)

type SampleUsecase struct {
	subscribeService service.MsgSubscribe
}

func NewSampleUsecase() SampleUsecase {
	return SampleUsecase{}
}

func (s SampleUsecase) Subscriber(ctx context.Context, msg string) (bool, error) {
	if err := s.subscribeService.Subscribe(ctx, msg); err != nil {
		return false, err
	}
	return true, nil
}
