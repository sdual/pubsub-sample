package usecase

import "context"

type SampleUsecase struct {
}

func (s SampleUsecase) Subscriber(ctx context.Context) (bool, error) {
	return true, nil
}
