package helloworld

import "github.com/sethvargo/go-githubactions"

type config struct {
	ExampleInput string
	action       *githubactions.Action
}

func NewConfig(action *githubactions.Action) (config, error) {
	return config{
		ExampleInput: action.GetInput("example-input"),
		action:       action,
	}, nil
}
