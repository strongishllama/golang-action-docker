package helloworld_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/sethvargo/go-githubactions"
	"github.com/stretchr/testify/require"

	helloworld "github.com/strongishllama/golang-action/internal/hello-world"
)

func TestRun(t *testing.T) {
	actionLog := bytes.NewBuffer(nil)
	envMap := map[string]string{
		"INPUT_EXAMPLE-INPUT": "hello-world",
	}
	getenv := func(key string) string {
		return envMap[key]
	}

	cfg, err := helloworld.NewConfig(githubactions.New(
		githubactions.WithWriter(actionLog),
		githubactions.WithGetenv(getenv),
	))
	require.NoError(t, err)

	require.NoError(t, helloworld.Run(context.Background(), cfg))
	require.Equal(t, "Example input: hello-world\n", actionLog.String())
}
