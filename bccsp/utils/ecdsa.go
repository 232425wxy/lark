package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/asn1"
	"errors"
	"fmt"
	"math/big"
)

type ECDSASignature struct {
	R, S *big.Int
}

var (
	// 这个被用来确保签名的S值低于或等于椭圆曲线阶的一半，fabric只接受低S值的签名
	curveHalfOrders = map[elliptic.Curve]*big.Int{
		elliptic.P224(): new(big.Int).Rsh(elliptic.P224().Params().N, 1),
		elliptic.P256(): new(big.Int).Rsh(elliptic.P256().Params().N, 1),
		elliptic.P384(): new(big.Int).Rsh(elliptic.P384().Params().N, 1),
		elliptic.P521(): new(big.Int).Rsh(elliptic.P521().Params().N, 1),
	}
)

func GetCurveHalfOrderAt(c elliptic.Curve) *big.Int {
	return big.NewInt(0).Set(curveHalfOrders[c])
}

func MarshalECDSASignature(r, s *big.Int) ([]byte, error) {
	// ASN.1: Abstract Syntax Notation One 计算机系统之间交换的数据消息如果用ASN.1进行描述，可以减少了双方的沟通成本。
	return asn1.Marshal(ECDSASignature{R: r, S: s})
}

// UnmarshalECDSASignature 反序列化得到ECDSA签名，返回签名的r和s。
func UnmarshalECDSASignature(raw []byte) (*big.Int, *big.Int, error) {
	// 反序列化
	sig := new(ECDSASignature)
	_, err := asn1.Unmarshal(raw, sig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed unmarshalling signature [%s]", err)
	}

	// 验证签名
	if sig.R == nil {
		return nil, nil, errors.New("invalid signature, R must be different from nil")
	}
	if sig.S == nil {
		return nil, nil, errors.New("invalid signature, S must be different from nil")
	}

	if sig.R.Sign() != 1 {
		return nil, nil, errors.New("invalid signature, R must be larger than 0")
	}
	if sig.S.Sign() != 1 {
		return nil, nil, errors.New("invalid signature, S must be larger than 0")
	}

	return sig.R, sig.S, nil
}

func SignatureToLowS(k *ecdsa.PublicKey, signature []byte) ([]byte, error) {
	r, s, err := UnmarshalECDSASignature(signature)
	if err != nil {
		return nil, err
	}
	s, err = ToLowS(k, s)
	if err != nil {
		return nil, err
	}
	return MarshalECDSASignature(r, s)
}

// IsLowS 判断ECDSA签名里的s是否小于椭圆曲线的阶的一半N/2。
func IsLowS(k *ecdsa.PublicKey, s *big.Int) (bool, error) {
	halfOrder, ok := curveHalfOrders[k.Curve]
	if !ok {
		return false, fmt.Errorf("curve not recognized [%s]", k.Curve)
	}
	return s.Cmp(halfOrder) != 1, nil
}

// TODO 为什么要让签名s小于椭圆曲线阶的一半呢？
func ToLowS(k *ecdsa.PublicKey, s *big.Int) (*big.Int, error) {
	lowS, err := IsLowS(k, s)
	if err != nil {
		return nil, err
	}
	if !lowS {
		s.Sub(k.Params().N, s)
		return s, nil
	}
	return s, nil
}
