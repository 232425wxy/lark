package utils

import (
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestUnmarshalECDSASignature(t *testing.T) {
	_, _, err := UnmarshalECDSASignature(nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed unmarshalling signature [")

	_, _, err = UnmarshalECDSASignature([]byte{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed unmarshalling signature [")

	_, _, err = UnmarshalECDSASignature([]byte{0})
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed unmarshalling signature [")

	sig, err := MarshalECDSASignature(big.NewInt(-1), big.NewInt(1))
	require.NoError(t, err)
	_, _, err = UnmarshalECDSASignature(sig)
	require.Error(t, err)
	require.Contains(t, err, "invalid signature, R must be larger than 0")

	sig, err = MarshalECDSASignature(big.NewInt(0), big.NewInt(1))
	require.NoError(t, err)
	_, _, err = UnmarshalECDSASignature(sig)
	require.Error(t, err)
	require.Contains(t, err, "invalid signature, R must be larger than 0")

	sig, err = MarshalECDSASignature(big.NewInt(1), big.NewInt(0))
	require.NoError(t, err)
	_, _, err = UnmarshalECDSASignature(sig)
	require.Error(t, err)
	require.Contains(t, err, "invalid signature, S must be larger than 0")

	sig, err = MarshalECDSASignature(big.NewInt(1), big.NewInt(-1))
	require.NoError(t, err)
	_, _, err = UnmarshalECDSASignature(sig)
	require.Error(t, err)
	require.Contains(t, err, "invalid signature, S must be larger than 0")

	sig, err = MarshalECDSASignature(big.NewInt(1), big.NewInt(1))
	require.NoError(t, err)
	R, S, err := UnmarshalECDSASignature(sig)
	require.NoError(t, err)
	require.Equal(t, R, big.NewInt(1))
	require.Equal(t, S, big.NewInt(1))
}
