package myfuzz

import (
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/message"
)

var (
	dummyOnFinishedHandling = func() {}
	dummyNodeID = ids.EmptyNodeID
)

func Fuzz(data []byte) int {
	codec, err := message.NewCodecWithMemoryPool("", prometheus.NewRegistry(), 2*units.MiB, 10*time.Second)
	if err != nil { return 1 }
	codec.Parse(data, dummyNodeID, dummyOnFinishedHandling)
	return 0
}
