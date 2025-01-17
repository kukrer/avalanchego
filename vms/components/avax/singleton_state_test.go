// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avax

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kukrer/savannahnode/database/memdb"
)

func TestSingletonState(t *testing.T) {
	require := require.New(t)

	db := memdb.New()
	s := NewSingletonState(db)

	isInitialized, err := s.IsInitialized()
	require.NoError(err)
	require.False(isInitialized)

	err = s.SetInitialized()
	require.NoError(err)

	isInitialized, err = s.IsInitialized()
	require.NoError(err)
	require.True(isInitialized)
}
