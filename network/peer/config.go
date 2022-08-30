// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"time"

	"github.com/kukrer/savannahnode/ids"
	"github.com/kukrer/savannahnode/message"
	"github.com/kukrer/savannahnode/network/throttling"
	"github.com/kukrer/savannahnode/snow/networking/router"
	"github.com/kukrer/savannahnode/snow/networking/tracker"
	"github.com/kukrer/savannahnode/snow/validators"
	"github.com/kukrer/savannahnode/utils/logging"
	"github.com/kukrer/savannahnode/utils/timer/mockable"
	"github.com/kukrer/savannahnode/version"
)

type Config struct {
	// Size, in bytes, of the buffer this peer reads messages into
	ReadBufferSize int
	// Size, in bytes, of the buffer this peer writes messages into
	WriteBufferSize      int
	Clock                mockable.Clock
	Metrics              *Metrics
	MessageCreator       message.Creator
	Log                  logging.Logger
	InboundMsgThrottler  throttling.InboundMsgThrottler
	Network              Network
	Router               router.InboundHandler
	VersionCompatibility version.Compatibility
	MySubnets            ids.Set
	Beacons              validators.Set
	NetworkID            uint32
	PingFrequency        time.Duration
	PongTimeout          time.Duration
	MaxClockDifference   time.Duration

	// Unix time of the last message sent and received respectively
	// Must only be accessed atomically
	LastSent, LastReceived int64

	// Tracks CPU/disk usage caused by each peer.
	ResourceTracker tracker.ResourceTracker

	PingMessage message.OutboundMessage
}
