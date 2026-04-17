package events

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/rogue0026/kafka-contracts/contracts"
)

type Message struct {
	EventID    string              `json:"event_id"`
	EventType  contracts.EventType `json:"event_type"`
	OccurredAt time.Time           `json:"occurred_at"`
	Producer   contracts.Producer  `json:"producer"`
	Payload    json.RawMessage     `json:"payload"`
}

func NewMessage(e Event, p contracts.Producer, occurredAt time.Time) (*Message, error) {
	raw, err := e.Raw()
	if err != nil {
		return nil, err
	}

	msg := &Message{
		EventID:    uuid.New().String(),
		EventType:  e.Type(),
		OccurredAt: occurredAt,
		Producer:   p,
		Payload:    raw,
	}
	return msg, nil
}

func (m Message) Raw() ([]byte, error) {
	raw, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return raw, nil
}
