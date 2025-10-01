package main

import (
	"context"
	"fmt"
)

// Counter is an example Actor for fan in/fan out
type Counter struct {
	value       int
	subscribers map[string]chan Event
	commands    chan Command
}

func (c *Counter) Inc() {
	c.commands <- Inc{}
}

func (c *Counter) Dec() {
	c.commands <- Dec{}
}

func (c *Counter) Reset() {
	c.commands <- Reset{}
}

// Run is our Actor's loop, ran as a goroutine. It processes commands until it receives ctx.Done()
func (c *Counter) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			c.cleanup()
			return ctx.Err()
		case cmd, ok := <-c.commands:
			if !ok {
				return nil
			}

			c.handle(cmd)
		}
	}
}

func (c *Counter) handle(cmd Command) {
	switch cmd.(type) {
	case Inc:
		c.value++
		c.broadcast()
	case Dec:
		c.value--
		c.broadcast()
	case Reset:
		c.value = 0
		c.broadcast()
	}
}

// broadcast fans out events to all subscribers.
func (c *Counter) broadcast() {
	for id, ch := range c.subscribers {
		select {
		// if our message pushes fine, do nothing
		case ch <- &Changed{c.value}:
		default:
			// If we cannot push to the channel, drop the subscriber and close its channel
			fmt.Printf("Dropping unresponsive sub: %s\n", id)
			close(ch)
			delete(c.subscribers, id)
		}
	}
}

// CommandsSink returns the Actor's command channel to allow subscribers to send messages back
func (c *Counter) CommandsSink() chan<- Command {
	return c.commands
}

// Subscribe returns a read only channel for the subscriber to read events through
func (c *Counter) Subscribe(id string) <-chan Event {
	c.subscribers[id] = make(chan Event, 4)
	return c.subscribers[id]
}

func (c *Counter) cleanup() {
	for id, ch := range c.subscribers {
		fmt.Println("Cleaning up subscriber:", id)
		close(ch)
		delete(c.subscribers, id)
	}
}
