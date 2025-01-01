package service

import (
	"context"
)

type MsgSubscribe struct {
}

func NewMsgSubscribe() MsgSubscribe {
	return MsgSubscribe{}
}

func (ms MsgSubscribe) Subscribe(ctx context.Context, msg string) error {
	_ = make(chan string)
	return nil
}
