package broadcaster

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

const (
	headersChanCap = 100
)

type HeadStreamer interface {
	Start(ctx context.Context)
	Stop()
	Next() *types.Header
}
