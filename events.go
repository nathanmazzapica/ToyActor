package main

type Event interface{ isEvent() }

type Changed struct {
	Value int
}

func (ch *Changed) isEvent() {}
