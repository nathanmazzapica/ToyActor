package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Subscriber struct {
	id   string
	ch   <-chan Event
	Dead bool
}

func (s *Subscriber) Listen(ctx context.Context, commands chan<- Command) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Simulate unresponsive client for the Actor (Counter) to drop
	if s.Dead {
		return
	}
	for {
		select {
		case <-ctx.Done():
			fmt.Println("we're done here.")
			return
		case evt, ok := <-s.ch:
			c, ok := evt.(*Changed)
			if ok {
				fmt.Printf("Subscriber%s received: %v\n", s.id, c.Value)
			}

		}

		choice := r.Float32()
		if choice > 0.95 {
			fmt.Printf("Subscriber{%s} is sending an inc command.\n", s.id)
			commands <- Inc{}
		}

	}
}
