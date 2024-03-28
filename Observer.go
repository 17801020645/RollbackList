package main

import (
	"context"
	"fmt"
	"sync"
)

type Event struct {
	Topic string
	Val   interface{}
}

type Observer interface {
	OnChange(ctx context.Context, e *Event) error
}

type EventBus interface {
	Subscribe(topic string, o Observer)
	Unsubscribe(topic string, o Observer)
	Publish(ctx context.Context, e *Event)
}

type BaseObserver struct {
	name string
}

func NewBaseObserver(name string) *BaseObserver {
	return &BaseObserver{
		name: name,
	}
}

func (b *BaseObserver) OnChange(ctx context.Context, e *Event) error {
	fmt.Printf("observer: %s, event key: %s, event val: %v", b.name, e.Topic, e.Val)
	// ...
	return nil
}

type BaseEventBus struct {
	mux       sync.RWMutex
	observers map[string]map[Observer]struct{}
}

func NewBaseEventBus() BaseEventBus {
	return BaseEventBus{
		observers: make(map[string]map[Observer]struct{}),
	}
}

func (b *BaseEventBus) Subscribe(topic string, o Observer) {
	b.mux.Lock()
	defer b.mux.Unlock()
	_, ok := b.observers[topic]
	if !ok {
		b.observers[topic] = make(map[Observer]struct{})
	}
	b.observers[topic][o] = struct{}{}
}

func (b *BaseEventBus) Unsubscribe(topic string, o Observer) {
	b.mux.Lock()
	defer b.mux.Unlock()
	delete(b.observers[topic], o)
}

type SyncEventBus struct {
	BaseEventBus
}

func NewSyncEventBus() *SyncEventBus {
	return &SyncEventBus{
		BaseEventBus: NewBaseEventBus(),
	}
}

func (s *SyncEventBus) Publish(ctx context.Context, e *Event) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	subscribers := s.observers[e.Topic]

	errs := make(map[Observer]error)
	for subscriber := range subscribers {
		if err := subscriber.OnChange(ctx, e); err != nil {
			errs[subscriber] = err
		}
	}

	s.handleErr(ctx, errs)
}

func (s *SyncEventBus) handleErr(ctx context.Context, errs map[Observer]error) {
	for o, err := range errs {
		// 处理 publish 失败的 observer
		fmt.Printf("observer: %v, err: %v", o, err)
	}
}
