package pkcs11

import (
	"crypto/ecdsa"
	"crypto/x509"
	"errors"
	"fmt"

	"github.com/232425wxy/lark/bccsp"
)

type ecdsaPublicKey struct {
	ski []byte // subject key identifier
	pub *ecdsa.PublicKey
}

// Bytes 将公钥转换为一串字节序列。
func (epk *ecdsaPublicKey) Bytes() (raw []byte, err error) {
	raw, err = x509.MarshalPKIXPublicKey(epk.pub)
	if err != nil {
		return nil, fmt.Errorf("Failed marshalling key [%s]", err)
	}
	return raw, nil
}

// SKI 返回公钥的标识符。
func (epk *ecdsaPublicKey) SKI() []byte {
	return epk.ski
}

// Symmetric ECDSA是一个非对称密码方案，所以此方法返回false。
func (epk *ecdsaPublicKey) Symmetric() bool {
	return false
}

// Private 返回一个bool值，指示该密钥是否是私钥。
func (epk *ecdsaPublicKey) Private() bool {
	return false
}

// PublicKey 返回非对称公钥/私钥对中相应的公钥部分，在对称密钥方案中，该方法返回一个错误。
func (epk *ecdsaPublicKey) PublicKey() (bccsp.Key, error) {
	return epk, nil
}

type ecdsaPrivateKey struct {
	ski []byte
	pub ecdsaPublicKey
}

// Bytes ECDSA的私钥的字节序列表现形式不予支持。
func (epk *ecdsaPrivateKey) Bytes() ([]byte, error) {
	return nil, errors.New("Not supported.")
}

// SKI 返回ECDSA私钥的标识符。
func (epk *ecdsaPrivateKey) SKI() []byte {
	return epk.ski
}

// Symmetric ECDSA是一个非对称密码方案，所以此方法返回false。
func (epk *ecdsaPrivateKey) Symmetric() bool {
	return false
}

// Private ECDSA是非对称密码方案，且该密钥是私钥，所以返回true。
func (epk *ecdsaPrivateKey) Private() bool {
	return true
}

// PublicKey 返回非对称公钥/私钥对中相应的公钥部分，在对称密钥方案中，该方法返回一个错误。
func (epk *ecdsaPrivateKey) PublicKey() (bccsp.Key, error) {
	return &epk.pub, nil
}