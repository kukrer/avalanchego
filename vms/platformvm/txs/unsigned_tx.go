// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"github.com/kukrer/savannahnode/ids"
	"github.com/kukrer/savannahnode/snow"
	"github.com/kukrer/savannahnode/vms/components/avax"
)

// UnsignedTx is an unsigned transaction
type UnsignedTx interface {
	// TODO: Remove this initialization pattern from both the platformvm and the
	// avm.
	snow.ContextInitializable
	Initialize(unsignedBytes []byte)
	Bytes() []byte

	// InputIDs returns the set of inputs this transaction consumes
	InputIDs() ids.Set

	Outputs() []*avax.TransferableOutput

	// Attempts to verify this transaction without any provided state.
	SyntacticVerify(ctx *snow.Context) error

	// Visit calls [visitor] with this transaction's concrete type
	Visit(visitor Visitor) error
}
