// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package fxs

import (
	"github.com/kukrer/savannahnode/ids"
	"github.com/kukrer/savannahnode/snow"
	"github.com/kukrer/savannahnode/vms/components/avax"
	"github.com/kukrer/savannahnode/vms/components/verify"
	"github.com/kukrer/savannahnode/vms/nftfx"
	"github.com/kukrer/savannahnode/vms/propertyfx"
	"github.com/kukrer/savannahnode/vms/secp256k1fx"
)

var (
	_ Fx = &secp256k1fx.Fx{}
	_ Fx = &nftfx.Fx{}
	_ Fx = &propertyfx.Fx{}
)

type ParsedFx struct {
	ID ids.ID
	Fx Fx
}

// Fx is the interface a feature extension must implement to support the AVM.
type Fx interface {
	// Initialize this feature extension to be running under this VM. Should
	// return an error if the VM is incompatible.
	Initialize(vm interface{}) error

	// Notify this Fx that the VM is in bootstrapping
	Bootstrapping() error

	// Notify this Fx that the VM is bootstrapped
	Bootstrapped() error

	// VerifyTransfer verifies that the specified transaction can spend the
	// provided utxo with no restrictions on the destination. If the transaction
	// can't spend the output based on the input and credential, a non-nil error
	// should be returned.
	VerifyTransfer(tx, in, cred, utxo interface{}) error

	// VerifyOperation verifies that the specified transaction can spend the
	// provided utxos conditioned on the result being restricted to the provided
	// outputs. If the transaction can't spend the output based on the input and
	// credential, a non-nil error  should be returned.
	VerifyOperation(tx, op, cred interface{}, utxos []interface{}) error
}

type FxOperation interface {
	verify.Verifiable
	snow.ContextInitializable
	avax.Coster

	Outs() []verify.State
}

type FxCredential struct {
	FxID              ids.ID `serialize:"false" json:"fxID"`
	verify.Verifiable `serialize:"true" json:"credential"`
}
