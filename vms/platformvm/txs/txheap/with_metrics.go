// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txheap

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/kukrer/savannahnode/ids"
	"github.com/kukrer/savannahnode/vms/platformvm/txs"
)

var _ Heap = &withMetrics{}

type withMetrics struct {
	Heap

	numTxs prometheus.Gauge
}

func NewWithMetrics(
	txHeap Heap,
	namespace string,
	registerer prometheus.Registerer,
) (Heap, error) {
	h := &withMetrics{
		Heap: txHeap,
		numTxs: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "count",
			Help:      "Number of transactions in the heap",
		}),
	}
	return h, registerer.Register(h.numTxs)
}

func (h *withMetrics) Add(tx *txs.Tx) {
	h.Heap.Add(tx)
	h.numTxs.Set(float64(h.Heap.Len()))
}

func (h *withMetrics) Remove(txID ids.ID) *txs.Tx {
	tx := h.Heap.Remove(txID)
	h.numTxs.Set(float64(h.Heap.Len()))
	return tx
}

func (h *withMetrics) RemoveTop() *txs.Tx {
	tx := h.Heap.RemoveTop()
	h.numTxs.Set(float64(h.Heap.Len()))
	return tx
}
