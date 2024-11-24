package suite

import (
	"context"
	"net/http"
	"testing"
)

type Suite struct {
	*testing.T
	HTTPClient http.Client
}

func NewSuite(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	ctx, cancelCtx := context.WithCancel(context.Background())

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	return ctx, &Suite{
		T:          t,
		HTTPClient: http.Client{},
	}
}
