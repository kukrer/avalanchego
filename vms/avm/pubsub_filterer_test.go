// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kukrer/savannahnode/ids"
	"github.com/kukrer/savannahnode/pubsub"
	"github.com/kukrer/savannahnode/vms/avm/txs"
	"github.com/kukrer/savannahnode/vms/components/avax"
	"github.com/kukrer/savannahnode/vms/secp256k1fx"
)

type mockFilter struct {
	addr []byte
}

func (f *mockFilter) Check(addr []byte) bool {
	return bytes.Equal(addr, f.addr)
}

func TestFilter(t *testing.T) {
	require := require.New(t)

	addrID := ids.ShortID{1}
	tx := txs.Tx{Unsigned: &txs.BaseTx{BaseTx: avax.BaseTx{
		Outs: []*avax.TransferableOutput{
			{
				Out: &secp256k1fx.TransferOutput{
					OutputOwners: secp256k1fx.OutputOwners{
						Addrs: []ids.ShortID{addrID},
					},
				},
			},
		},
	}}}
	addrBytes := addrID[:]

	fp := pubsub.NewFilterParam()
	err := fp.Add(addrBytes)
	require.NoError(err)

	parser := NewPubSubFilterer(&tx)
	fr, _ := parser.Filter([]pubsub.Filter{&mockFilter{addr: addrBytes}})
	require.Equal([]bool{true}, fr)
}
