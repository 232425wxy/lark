package bccsp

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAESOpts(t *testing.T) {
	test := func(ephemeral bool) {
		for _, opts := range []KeyGenOpts{
			&AES128KeyGenOpts{Temporary: ephemeral},
			&AES192KeyGenOpts{Temporary: ephemeral},
			&AES256KeyGenOpts{Temporary: ephemeral},
		} {
			expectedAlgorithm := reflect.TypeOf(opts).String()
			require.Contains(t, expectedAlgorithm, "AES")
			require.Equal(t, ephemeral, opts.Ephemeral())
		}
	}
	test(true)
	test(false)
}

func TestECDSAOpts(t *testing.T) {
	test := func(ephemeral bool) {
		for _, opts := range []KeyGenOpts{
			&ECDSAP256KeyGenOpts{Temporary: ephemeral},
			&ECDSAP384KeyGenOpts{Temporary: ephemeral},
		} {
			expectedAlgorithm := reflect.TypeOf(opts).String()
			require.Contains(t, expectedAlgorithm, "ECDSA")
			require.Equal(t, ephemeral, opts.Ephemeral())
		}
	}
	test(true)
	test(false)
	opts := &ECDSAReRandKeyOpts{Temporary: true}
	require.Equal(t, true, opts.Ephemeral())
	opts.Temporary = false
	require.Equal(t, false, opts.Temporary)
	require.Equal(t, "ECDSA_RERAND", opts.Algorithm())
	require.Empty(t, opts.ExpansionValue())
}

func TestHashOpts(t *testing.T) {
	for _, opts := range []HashOpts{&SHA256Opts{}, &SHA384Opts{}, &SHA3_256Opts{}, &SHA3_384Opts{}} {
		s := strings.Replace(reflect.TypeOf(opts).String(), "*bccsp.", "", -1)
		algorithm := strings.Replace(s, "Opts", "", -1)
		require.Equal(t, algorithm, opts.Algorithm())
		opts2, err := GetHashOpt(algorithm)
		require.NoError(t, err)
		require.Equal(t, opts.Algorithm(), opts2.Algorithm())
	}
	_, err := GetHashOpt("fool")
	require.Error(t, err)
	require.Contains(t, err.Error(), "hash function not recognized")
	require.Equal(t, "SHA", (&SHAOpts{}).Algorithm())
}

func TestHMAC(t *testing.T) {
	opts := &HMACTruncated256AESDeriveKeyOpts{
		Arg:       []byte("arg"),
	}
	require.False(t, opts.Ephemeral())
	opts.Temporary = true
	require.True(t, opts.Ephemeral())
	require.Equal(t, "HMAC_TRUNCATED_256", opts.Algorithm())
	require.Equal(t, []byte("arg"), opts.Argument())

	opts2 := &HMACDeriveKeyOpts{
		Arg:       []byte("arg"),
	}
	require.False(t, opts2.Ephemeral())
	opts2.Temporary = true
	require.True(t, opts2.Ephemeral())
	require.Equal(t, "HMAC", opts2.Algorithm())
	require.Equal(t, []byte("arg"), opts2.Argument())
}

func TestKeyGenOpts(t *testing.T) {
	expectedAlgorithms := map[reflect.Type]string{
		reflect.TypeOf(&HMACImportKeyOpts{}): "HMAC",
		reflect.TypeOf(&X509PublicKeyImportOpts{}): "X509Certificate",
		reflect.TypeOf(&AES256ImportKeyOpts{}): "AES",
	}
	test := func(ephemeral bool) {
		for _, opts := range []KeyGenOpts{
			&HMACImportKeyOpts{ephemeral},
			&X509PublicKeyImportOpts{ephemeral},
			&AES256ImportKeyOpts{ephemeral},
		} {
			expectedAlgorithm := expectedAlgorithms[reflect.TypeOf(opts)]
			require.Equal(t, expectedAlgorithm, opts.Algorithm())
			require.Equal(t, ephemeral, opts.Ephemeral())
		}
	}

	test(true)
	test(false)
}