package handler

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Value string
}

func ParseMassege(rawMsg []byte) (*Message, error) {
	var msg Message
	if err := json.Unmarshal(rawMsg, &msg); err != nil {
		return nil, fmt.Errorf("failed to parse pubsub message: %w", err)
	}
	return &msg, nil
}
