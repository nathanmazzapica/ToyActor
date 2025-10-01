package main

type Command interface{ isCommand() }

type Inc struct{}

func (Inc) isCommand() {}

type Dec struct{}

func (Dec) isCommand() {}

type Reset struct{}

func (Reset) isCommand() {}
