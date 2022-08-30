// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"github.com/kukrer/savannahnode/codec"
	"github.com/kukrer/savannahnode/codec/linearcodec"
	"github.com/kukrer/savannahnode/utils/units"
	"github.com/kukrer/savannahnode/utils/wrappers"
)

const (
	codecVersion   uint16 = 0
	maxMessageSize        = 512 * units.KiB
	maxSliceLen           = maxMessageSize
)

// Codec does serialization and deserialization
var c codec.Manager

func init() {
	c = codec.NewManager(maxMessageSize)
	lc := linearcodec.NewCustomMaxLength(maxSliceLen)

	errs := wrappers.Errs{}
	errs.Add(
		lc.RegisterType(&Tx{}),
		c.RegisterCodec(codecVersion, lc),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
}
