// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metervm

import (
	"github.com/ava-labs/avalanchego/utils/metric"
	"github.com/ava-labs/avalanchego/utils/wrappers"
	"github.com/prometheus/client_golang/prometheus"
)

type stateSummaryMetrics struct {
	stateSyncEnabled,
	getLastStateSummary,
	parseStateSummary,
	getStateSummary,
	setSyncableStateSummaries,
	getOngoingStateSyncSummary,
	getStateSyncResult,
	setLastStateSummaryBlock metric.Averager
}

func newStateSummaryMetrics(namespace string, reg prometheus.Registerer) (stateSummaryMetrics, error) {
	errs := wrappers.Errs{}
	return stateSummaryMetrics{
		stateSyncEnabled:           newAverager(namespace, "state_sync_enabled", reg, &errs),
		getLastStateSummary:        newAverager(namespace, "get_last_state_summary", reg, &errs),
		parseStateSummary:          newAverager(namespace, "parse_state_summary", reg, &errs),
		getStateSummary:            newAverager(namespace, "get_state_summary", reg, &errs),
		setSyncableStateSummaries:  newAverager(namespace, "set_syncable_state_summaries", reg, &errs),
		getOngoingStateSyncSummary: newAverager(namespace, "get_ongoing_state_sync_summary", reg, &errs),
		getStateSyncResult:         newAverager(namespace, "get_state_sync_results", reg, &errs),
		setLastStateSummaryBlock:   newAverager(namespace, "set_last_state_summary_block", reg, &errs),
	}, errs.Err
}