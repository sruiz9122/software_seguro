package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDummy(t *testing.T) {
	assertions := require.New(t)

	result := 22

	assertions.Equal(22, result)
}
