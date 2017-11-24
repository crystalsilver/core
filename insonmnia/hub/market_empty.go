// +build infra

package hub

import (
	"golang.org/x/net/context"
)

type emptyMarker struct{}

func (m *emptyMarker) OrderExists(ID string) (bool, error) { return true, nil }
func (m *emptyMarker) CancelOrder(ID string) error         { return nil }

// NewMarket constructs a new SONM marketplace client.
func NewMarket(ctx context.Context, addr string) (Market, error) {
	return &emptyMarker{}, nil
}
