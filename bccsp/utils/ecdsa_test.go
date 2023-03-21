package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
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
	require.Contains(t, err.Error(), "invalid signature, R must be larger than 0")

	sig, err = MarshalECDSASignature(big.NewInt(0), big.NewInt(1))
	require.NoError(t, err)
	_, _, err = UnmarshalECDSASignature(sig)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid signature, R must be larger than 0")

	sig, err = MarshalECDSASignature(big.NewInt(1), big.NewInt(0))
	require.NoError(t, err)
	_, _, err = UnmarshalECDSASignature(sig)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid signature, S must be larger than 0")

	sig, err = MarshalECDSASignature(big.NewInt(1), big.NewInt(-1))
	require.NoError(t, err)
	_, _, err = UnmarshalECDSASignature(sig)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid signature, S must be larger than 0")

	sig, err = MarshalECDSASignature(big.NewInt(1), big.NewInt(1))
	require.NoError(t, err)
	R, S, err := UnmarshalECDSASignature(sig)
	require.NoError(t, err)
	require.Equal(t, R, big.NewInt(1))
	require.Equal(t, S, big.NewInt(1))
}

func TestIsLowS(t *testing.T) {
	// 生成ECDSA私钥
	lowLevelKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	lowS, err := IsLowS(&lowLevelKey.PublicKey, big.NewInt(0)) // 0肯定小于 elliptic.P256().N / 2
	require.NoError(t, err)
	require.True(t, lowS)

	s := new(big.Int)
	s.Set(GetCurveHalfOrderAt(elliptic.P256()))

	lowS, err = IsLowS(&lowLevelKey.PublicKey, s)
	require.NoError(t, err)
	require.True(t, lowS)

	s.Add(s, big.NewInt(1))
	lowS, err = IsLowS(&lowLevelKey.PublicKey, s)
	require.NoError(t, err)
	require.False(t, lowS)

	s, err = ToLowS(&lowLevelKey.PublicKey, s)
	require.NoError(t, err)
	lowS, err = IsLowS(&lowLevelKey.PublicKey, s)
	require.NoError(t, err)
	require.True(t, lowS)
}

func TestSignatureToLowS(t *testing.T) {
	lowLevelKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	s := new(big.Int)
	s.Set(GetCurveHalfOrderAt(elliptic.P256()))
	s.Add(s, big.NewInt(1))

	lowS, err := IsLowS(&lowLevelKey.PublicKey, s)
	require.NoError(t, err)
	require.False(t, lowS)

	sig, err := MarshalECDSASignature(big.NewInt(1), s)
	require.NoError(t, err)

	// PublicKey公钥提供了椭圆曲线的参数阶N，将其作为SignatureToLows的输入，可以帮助将signature的s转换成小于N/2的值
	sig2, err := SignatureToLowS(&lowLevelKey.PublicKey, sig)
	require.NoError(t, err)
	_, s, err = UnmarshalECDSASignature(sig2)
	require.NoError(t, err)
	lowS, err = IsLowS(&lowLevelKey.PublicKey, s)
	require.NoError(t, err)
	require.True(t, lowS)
}
