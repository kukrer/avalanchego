// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package secp256k1fx

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kukrer/savannahnode/ids"
	"github.com/kukrer/savannahnode/vms/components/verify"
)

func TestMintOperationVerifyNil(t *testing.T) {
	require := require.New(t)
	op := (*MintOperation)(nil)
	require.ErrorIs(op.Verify(), errNilMintOperation)
}

func TestMintOperationOuts(t *testing.T) {
	require := require.New(t)
	op := &MintOperation{
		MintInput: Input{
			SigIndices: []uint32{0},
		},
		MintOutput: MintOutput{
			OutputOwners: OutputOwners{
				Threshold: 1,
				Addrs: []ids.ShortID{
					addr,
				},
			},
		},
		TransferOutput: TransferOutput{
			Amt: 1,
			OutputOwners: OutputOwners{
				Locktime:  0,
				Threshold: 1,
			},
		},
	}

	require.Len(op.Outs(), 2)
}

func TestMintOperationState(t *testing.T) {
	require := require.New(t)
	intf := interface{}(&MintOperation{})
	_, ok := intf.(verify.State)
	require.False(ok)
}
