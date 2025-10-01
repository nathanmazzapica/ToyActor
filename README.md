# Toy Actor

This sample project demonstrates a small **Actor Model + Pub/Sub (fan-out)** system in Go.

I built it to understand how a multiplayer lobby can:

- receive commands from clients (e.g. Inc, Dec, Reset)
- process them in a single goroutine actor loop
- broadcast events back out to all subscribers (Incremented, Decremented, Reset)
- handle unresponsive/slow clients
- propagate shutdown signals to close things gracefully with context


## What is all this?
**Actor**: the `Counter` goroutine, which owns all state and processes `Command`s
**Pub/Sub**: a notifier that fans out `Event`s to all subscribed goroutines

These patterns give you a clean model for multiplayer game state or any event-driven backend.

## Why did you publish this?
As a reference for myself and for others who want to explore concurrency in Go. It's small, but it demonstrates the basic structure of real-time systems without any locking or global mutable state

## How do I run this?
It's pretty simple, just run: `go run .`
