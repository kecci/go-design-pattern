package main

import "github.com/kecci/design-pattern-go/behavioral_patterns/command/usecase"

func main() {
	tv := &usecase.Tv{}

	onCommand := &usecase.OnCommand{
		Device: tv,
	}

	offCommand := &usecase.OffCommand{
		Device: tv,
	}

	onButton := &usecase.Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton := &usecase.Button{
		Command: offCommand,
	}
	offButton.Press()
}
