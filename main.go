package main

import (
	"context"

	"github.com/sethvargo/go-githubactions"

	helloworld "github.com/strongishllama/golang-action/internal/hello-world"
)

func main() {
	action := githubactions.New()

	cfg, err := helloworld.NewConfig(action)
	if err != nil {
		action.Fatalf("%v", err)
	}

	if err := helloworld.Run(context.Background(), cfg); err != nil {
		action.Fatalf("%v", err)
	}
}
