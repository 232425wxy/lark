package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnmarshalECDSASignature(t *testing.T) {
	_, _, err := UnmarshalECDSASignature(nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed unmarshalling signature [")
}
