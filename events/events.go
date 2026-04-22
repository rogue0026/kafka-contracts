package events

import (
	"encoding/json"

	"github.com/rogue0026/kafka-contracts/contracts"
)

type Event interface {
	Type() contracts.EventType
	Raw() ([]byte, error)
}

type UserCreated struct {
	UserID uint64 `json:"user_id"`
}

func (e UserCreated) Raw() ([]byte, error) {
	raw, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

func (e UserCreated) Type() contracts.EventType {
	return contracts.UserCreated
}

type OrderCreated struct {
	UserID  uint64 `json:"user_id"`
	OrderID uint64 `json:"order_id"`
}

func (e OrderCreated) Raw() ([]byte, error) {
	raw, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

func (e OrderCreated) Type() contracts.EventType {
	return contracts.OrderCreated
}

type OrderPayedFor struct {
	PaymentID uint64 `json:"payment_id"`
}

func (e OrderPayedFor) Raw() ([]byte, error) {
	raw, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

func (e OrderPayedFor) Type() contracts.EventType {
	return contracts.OrderPayedFor
}

type FundsAdded struct {
	UserID uint64 `json:"user_id"`
}

func (e FundsAdded) Raw() ([]byte, error) {
	raw, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

func (e FundsAdded) Type() contracts.EventType {
	return contracts.FundsAdded
}

type FundsDebitted struct {
	UserID uint64 `json:"user_id"`
}

func (e FundsDebitted) Raw() ([]byte, error) {
	raw, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

func (e FundsDebitted) Type() contracts.EventType {
	return contracts.FundsDebitted
}
