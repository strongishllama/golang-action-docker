package helloworld

import (
	"context"
)

func Run(ctx context.Context, cfg config) error {
	cfg.action.Infof("Example input: %s", cfg.ExampleInput)
	return nil
}
