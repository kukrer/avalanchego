// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"github.com/kukrer/savannahnode/ids"
	"github.com/kukrer/savannahnode/snow/networking/router"
	"github.com/kukrer/savannahnode/version"
)

var _ router.ExternalHandler = &testHandler{}

type testHandler struct {
	router.InboundHandler
	ConnectedF    func(nodeID ids.NodeID, nodeVersion *version.Application, subnetID ids.ID)
	DisconnectedF func(nodeID ids.NodeID)
}

func (h *testHandler) Connected(id ids.NodeID, nodeVersion *version.Application, subnetID ids.ID) {
	if h.ConnectedF != nil {
		h.ConnectedF(id, nodeVersion, subnetID)
	}
}

func (h *testHandler) Disconnected(id ids.NodeID) {
	if h.DisconnectedF != nil {
		h.DisconnectedF(id)
	}
}
