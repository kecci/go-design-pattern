package main

import "github.com/kecci/go-design-pattern/behavioral_patterns/command/usecase"

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
