package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	c := Counter{
		value:       0,
		subscribers: make(map[string]chan Event),
		commands:    make(chan Command, 16),
	}

	go c.Run(ctx)

	subs := make([]Subscriber, 10)

	for i := 0; i < 3; i++ {
		subs[i] = Subscriber{id: strconv.Itoa(i), ch: c.Subscribe(strconv.Itoa(i))}
		go subs[i].Listen(ctx, c.CommandsSink())
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("We're done!")
			return
		default:
			choice := r.Intn(8)
			switch choice {
			case 1:
				c.Dec()
			case 2:
				c.Reset()
			default:
				c.Inc()
			}

		}
		time.Sleep(250 * time.Millisecond)
	}

}
