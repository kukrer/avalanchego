// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"github.com/kukrer/savannahnode/snow"
	"github.com/kukrer/savannahnode/snow/uptime"
	"github.com/kukrer/savannahnode/utils"
	"github.com/kukrer/savannahnode/utils/timer/mockable"
	"github.com/kukrer/savannahnode/vms/platformvm/config"
	"github.com/kukrer/savannahnode/vms/platformvm/fx"
	"github.com/kukrer/savannahnode/vms/platformvm/reward"
	"github.com/kukrer/savannahnode/vms/platformvm/utxo"
)

type Backend struct {
	Config       *config.Config
	Ctx          *snow.Context
	Clk          *mockable.Clock
	Fx           fx.Fx
	FlowChecker  utxo.Verifier
	Uptimes      uptime.Manager
	Rewards      reward.Calculator
	Bootstrapped *utils.AtomicBool
}
