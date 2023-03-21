package bccsp

import (
	"crypto"
	"fmt"
)

const (
	// ECDSA 代表默认安全级别的椭圆曲线数字签名算法(KeyGen, Import, Sign, Verify)。
	ECDSA = "ECDSA"

	// ECDSAP256 代表P-256曲线上的椭圆曲线数字签名算法。
	ECDSAP256 = "ECDSAP256"

	// ECDSAP384 代表P-384曲线上的椭圆曲线数字签名算法。
	ECDSAP384 = "ECDSAP384"

	// ECDSAReRand 密钥重新随机化。
	// TODO 什么是密钥重新随机化？猜测是密钥派生。
	ECDSAReRand = "ECDSA_RERAND"

	// AES 代表默认安全级别的AES加密算法。
	AES = "AES"

	// AES128 代表128位安全级别的AES加密算法。
	AES128 = "AES128"

	// AES192 代表192位安全级别的AES加密算法。
	AES192 = "AES192"

	// AES256 代表256位安全级别的AES加密算法。
	AES256 = "AES256"

	// HMAC 代表哈希验证码。
	HMAC = "HAMC"

	// HMACTruncated256 代表截断前256位的哈希验证码。
	HMACTruncated256 = "HMAC_TRUNCATED_256"

	// SHA 代表默认系列的哈希算法。
	SHA = "SHA"

	// SHA2 代表SHA2哈希族的一个标识符。
	SHA2 = "SHA2"

	// SHA3 代表SHA3哈希族的一个标识符。
	SHA3 = "SHA3"

	SHA256 = "SHA256"

	SHA384 = "SHA384"

	SHA3_256 = "SHA3_256"

	SHA3_384 = "SHA3_384"

	// X509Certificate 代表用于X509证书的相关操作。
	X509Certificate = "X509Certificate"

	// IDEMIX 身份混淆
	IDEMIX = "IDEMIX"
)

// ECDSAKeyGenOpts 包含用于ECDSA密钥生成的选项。
type ECDSAKeyGenOpts struct {
	Temporary bool
}

// Algorithm 返回密钥生成算法的标识符。
func (opts *ECDSAKeyGenOpts) Algorithm() string {
	return ECDSA
}

// Ephemeral 如果要生成的密钥必须是短暂的，那么Ephemeral返回true，否则返回false。
func (opts *ECDSAKeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

// ECDSAPKIXPublicKeyImportOpts 包含用于以PKIX格式导入ECDSA公钥的选项。
type ECDSAPKIXPublicKeyImportOpts struct {
	Temporary bool
}

// Algorithm 返回密钥导入算法的标识符。
func (opts *ECDSAPKIXPublicKeyImportOpts) Algorithm() string {
	return ECDSA
}

// Ephemeral 如果要生成的密钥必须是短暂的，则该方法返回true，否则返回false。
func (opts *ECDSAPKIXPublicKeyImportOpts) Ephemeral() bool {
	return opts.Temporary
}

// ECDSAGoPublicKeyImportOpts 包含从ecdsa.PublicKey导入ECDCA密钥的选项。
type ECDSAGoPublicKeyImportOpts struct {
	Temporary bool
}

// Algorithm 返回密钥导入算法的标识符。
func (opts *ECDSAGoPublicKeyImportOpts) Algorithm() string {
	return ECDSA
}

// Ephemeral 如果生成的密钥必须是短暂的，则返回true，否则返回false。
func (opts *ECDSAGoPublicKeyImportOpts) Ephemeral() bool {
	return opts.Temporary
}

// ECDSAReRandKeyOpts 包含ECDSA密钥重新随机化的选项。
type ECDSAReRandKeyOpts struct {
	Temporary bool
	Expansion []byte
}

// Algorithm 返回密钥派生算法的标识符.
func (opts *ECDSAReRandKeyOpts) Algorithm() string {
	return ECDSAReRand
}

// Ephemeral 如果生成的密钥必须是短暂的，则返回true，否则返回false。
func (opts *ECDSAReRandKeyOpts) Ephemeral() bool {
	return opts.Temporary
}

// ExpansionValue 返回重新随机化系数.
func (opts *ECDSAReRandKeyOpts) ExpansionValue() []byte {
	return opts.Expansion
}

// AESKeyGenOpts 包含默认安全级别下的AES密钥生成选项。
type AESKeyGenOpts struct {
	Temporary bool
}

// Algorithm 返回密钥生成算法的标识符.
func (opts *AESKeyGenOpts) Algorithm() string {
	return AES
}

// Ephemeral 如果生成的密钥必须是短暂的,则该方法返回true,否则返回false。
func (opts *AESKeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

// HMACTruncated256AESDeriveKeyOpts 包含HMAC截断在256位密钥衍生的选项。
type HMACTruncated256AESDeriveKeyOpts struct {
	Temporary bool
	Arg       []byte
}

// Algorithm 返回密钥派生算法的标识符。
func (opts *HMACTruncated256AESDeriveKeyOpts) Algorithm() string {
	return HMACTruncated256
}

// Ephemeral 如果生成的密钥必须是短暂的,则该方法返回true,否则返回false。
func (opts *HMACTruncated256AESDeriveKeyOpts) Ephemeral() bool {
	return opts.Temporary
}

// Argument 返回传给HMAC的参数。
func (opts *HMACTruncated256AESDeriveKeyOpts) Argument() []byte {
	return opts.Arg
}

// HMACDeriveKeyOpts 包含HMAC密钥派生的选项。
type HMACDeriveKeyOpts struct {
	Temporary bool
	Arg       []byte
}

// Algorithm 返回密钥派生算法的标识符。
func (opts *HMACDeriveKeyOpts) Algorithm() string {
	return HMAC
}

// Ephemeral 如果生成的密钥必须是暂时的，则该方法返回true，否则返回false。
func (opts *HMACDeriveKeyOpts) Ephemeral() bool {
	return opts.Temporary
}

// Argument 返回传给HMAC的参数。
func (opts *HMACDeriveKeyOpts) Argument() []byte {
	return opts.Arg
}

// AES256ImportKeyOpts 包含导入AES256密钥的选项。
type AES256ImportKeyOpts struct {
	Temporary bool
}

// Algorithm 返回密钥导入算法标识符。
func (opts *AES256ImportKeyOpts) Algorithm() string {
	return AES
}

// Ephemeral 如果生成的密钥必须是暂时的，则该方法返回true，否则返回false。
func (opts *AES256ImportKeyOpts) Ephemeral() bool {
	return opts.Temporary
}

// HMACImportKeyOpts 包含导入HMAC密钥的选项。
type HMACImportKeyOpts struct {
	Temporary bool
}

// Algorithm 返回密钥导入算法的标识符。
func (opts *HMACImportKeyOpts) Algorithm() string {
	return HMAC
}

// Ephemeral 如果生成的密钥必须是暂时的，则该方法返回true，否则返回false。
func (opts *HMACImportKeyOpts) Ephemeral() bool {
	return opts.Temporary
}

// SHAOpts 包含计算SHA的选项。
type SHAOpts struct{}

// Algorithm 返回哈希算法标识符。
func (opts *SHAOpts) Algorithm() string {
	return SHA
}

// X509PublicKeyImportOpts 包含从X509证书里导入公钥的选项。
type X509PublicKeyImportOpts struct {
	Temporary bool
}

// Algorithm 返回密钥导入算法的标识符。
func (opts *X509PublicKeyImportOpts) Algorithm() string {
	return X509Certificate
}

// Ephemeral 如果生成的密钥必须是暂时的，则该方法返回true，否则返回false。
func (opts *X509PublicKeyImportOpts) Ephemeral() bool {
	return opts.Temporary
}

// SHA256Opts 包含于SHA-256算法相关的选项。
type SHA256Opts struct{}

// Algorithm 返回哈希算法的标识符
func (opts *SHA256Opts) Algorithm() string {
	return SHA256
}

// SHA384Opts 包含与SHA-384算法相关的选项。
type SHA384Opts struct{}

// Algorithm 返回哈希算法的标识符。
func (opts *SHA384Opts) Algorithm() string {
	return SHA384
}

// SHA3_256Opts 包含与SHA3_256算法相关的选项。
type SHA3_256Opts struct{}

// Algorithm 返回哈希算法的标识符。
func (opts *SHA3_256Opts) Algorithm() string {
	return SHA3_256
}

// SHA3_384Opts 包含与SHA3_384算法相关的选项。
type SHA3_384Opts struct{}

// Algorithm 返回哈希算法的标识符。
func (opts *SHA3_384Opts) Algorithm() string {
	return SHA3_384
}

// GetHashOpt 根据给定的哈希函数名，返回对应的哈希算法选项。
func GetHashOpt(hashFunction string) (HashOpts, error) {
	switch hashFunction {
	case SHA256:
		return &SHA256Opts{}, nil
	case SHA384:
		return &SHA384Opts{}, nil
	case SHA3_256:
		return &SHA3_256Opts{}, nil
	case SHA3_384:
		return &SHA3_384Opts{}, nil
	default:
		return nil, fmt.Errorf("hash function not recognized [%s]", hashFunction)
	}
}

// RevocationAlgorithm 标识撤销算法。
type RevocationAlgorithm int32

const (
	// AlgNoRevocation 意味着不支持撤销。
	AlgNoRevocation RevocationAlgorithm = iota
)

// IdemixIssuerKeyGenOpts 包含Idemix Issuer密钥生成的选项，属性名列表是可选的。
type IdemixIssuerKeyGenOpts struct {
	// Temporary 表示密钥是否是暂时的。
	Temporary bool
	// AttributeNames 属性的列表。
	AttributeNames []string
}

// Algorithm 返回密钥生成算法的标识符。
func (opts *IdemixIssuerKeyGenOpts) Algorithm() string {
	return IDEMIX
}

// Ephemeral 如果生成的密钥必须是暂时的，则该方法返回true，否则返回false。
func (opts *IdemixIssuerKeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

// IdemixIssuerPublicKeyImportOpts 包含导入Idemix Issuer公钥的选项。
type IdemixIssuerPublicKeyImportOpts struct {
	Temporary bool
	// AttributeNames 是一个属性列表，与导入的公钥的相关。
	AttributeNames []string
}

// Algorithm 返回密钥生成算法的标识符。
func (opts *IdemixIssuerPublicKeyImportOpts) Algorithm() string {
	return IDEMIX
}

// Ephemeral 如果生成的密钥必须是暂时的，则该方法返回true，否则返回false。
func (opts *IdemixIssuerPublicKeyImportOpts) Ephemeral() bool {
	return opts.Temporary
}

// IdemixUserSecretKeyGenOpts 包含用于生成Idemix凭证密钥的选项。
type IdemixUserSecretKeyGenOpts struct {
	Temporary bool
}

// Algorithm 返回密钥生成算法的标识符。
func (opts *IdemixUserSecretKeyGenOpts) Algorithm() string {
	return IDEMIX
}

// Ephemeral 如果生成的密钥必须是暂时的，则该方法返回true，否则返回false。
func (opts *IdemixUserSecretKeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

// IdemixUserSecretKeyImportOpts 包含用于导入Idemix凭证密钥的选项。
type IdemixUserSecretKeyImportOpts struct {
	Temporary bool
}

// Algorithm 返回密钥生成算法的标识符。
func (opts *IdemixUserSecretKeyImportOpts) Algorithm() string {
	return IDEMIX
}

// Ephemeral 如果生成的密钥必须是暂时的，则该方法返回true，否则返回false。
func (opts *IdemixUserSecretKeyImportOpts) Ephemeral() bool {
	return opts.Temporary
}

// IdemixNymKeyDerivationOpts 包含从凭证密钥创建一个新的无关联的假名的选项，该选项与指定的公钥发行者有关。
type IdemixNymKeyDerivationOpts struct {
	Temporary bool
	// IssuerPK 是发行者的公钥
	IssuerPK Key
}

// Algorithm 返回密钥派生算法的标识符。
func (opts *IdemixNymKeyDerivationOpts) Algorithm() string {
	return IDEMIX
}

// Ephemeral 如果生成的密钥必须是暂时的，则该方法返回true，否则返回false。
func (opts *IdemixNymKeyDerivationOpts) Ephemeral() bool {
	return opts.Temporary
}

// IssuerPublicKey 返回发行者的公钥，发行者的公钥用于从凭证密钥创建一个无关联的假名。
func (opts *IdemixNymKeyDerivationOpts) IssuerPublicKey() Key {
	return opts.IssuerPK
}

// IdemixNymPublicKeyImportOpts 包含导入假名的公共部分的选项。
type IdemixNymPublicKeyImportOpts struct {
	Temporary bool
}

// Algorithm 返回密钥派生算法的标识符。
func (opts *IdemixNymPublicKeyImportOpts) Algorithm() string {
	return IDEMIX
}

// Ephemeral 如果派生出的密钥必须是暂时的，则该方法返回true，否则返回false。
func (opts *IdemixNymPublicKeyImportOpts) Ephemeral() bool {
	return opts.Temporary
}

// IdemixCredentialRequestSignerOpts 包含创建Idemix凭证请求的选项。
type IdemixCredentialRequestSignerOpts struct {
	// Attributes 包含一个将被包含在证书中的属性索引的列表，这些索引是关于IdemixIssuerKeyGenOpts#AttributeNames的。
	Attributes []int
	IssuerPK   Key
	// IssuerNonce 由发行人生成，并由客户端用来生成凭证请求，一旦发行人得到了凭证请求，它就会检查nonce值是否相同。
	IssuerNonce []byte
	// H 是将被用到的哈希函数
	H crypto.Hash
}

func (opts *IdemixCredentialRequestSignerOpts) HashFunc() crypto.Hash {
	return opts.H
}

// IssuerPublicKey 返回发行者的公钥，发行者的公钥用于从凭证密钥创建一个无关联的假名。
func (opts *IdemixCredentialRequestSignerOpts) IssuerPublicKey() Key {
	return opts.IssuerPK
}

// IdemixAttributeType 表示一个Idemix属性的类型。
type IdemixAttributeType int

const (
	// IdemixHiddenAttribute 表示一个隐藏属性。
	IdemixHiddenAttribute IdemixAttributeType = iota
	// IdemixBytesAttribute 表示一串字节序列。
	IdemixBytesAttribute
	// IdemixIntAttribute 表示一个整数。
	IdemixIntAttribute
)

type IdemixAttribute struct {
	// Type 表示属性的类型。
	Type IdemixAttributeType
	// Value 表示属性的值。
	Value interface{}
}

// IdemixCredentialSignerOpts 包含从凭证请求开始产生凭证的选项。
type IdemixCredentialSignerOpts struct {
	// Attributes 是证书内要包含的属性，其中，IdemixHiddenAttribute类型的属性没有包含在内。
	Attributes []IdemixAttribute
	IssuerPK   Key
	H          crypto.Hash
}

func (opts *IdemixCredentialSignerOpts) HashFunc() crypto.Hash {
	return opts.H
}

func (opts *IdemixCredentialSignerOpts) IssuerPublicKey() Key {
	return opts.IssuerPK
}

// IdemixSignerOpts 包含用于生成Idemix签名的选项。
type IdemixSignerOpts struct {
	// Nym 是要使用的假名。
	Nym      Key
	IssuerPK Key
	// Credential 是由发行人签名的凭证。
	Credential []byte
	// Attributes 指定哪些属性应该被披露，哪些不应该。如果Attributes[i].Type = IdemixHiddenAttribute，
	// 那么第i个凭证属性不应该被披露，否则第i个凭证属性将被披露。在验证时，如果第i个属性被披露（Attributes[i].Type !=
	// IdemixHiddenAttribute），那么Attributes[i].Value必须被相应设置。
	Attributes []IdemixAttribute
	// RhIndex 是包含撤销处理程序的属性的索引，这个属性不能被公开。
	RhIndex int
	// CRI 包含凭证撤销信息。
	CRI []byte
	// Epoch 是指该签名应针对的撤销时间。
	Epoch int
	// RevocationPublicKey 是撤销的公钥。
	RevocationPublicKey Key
	H                   crypto.Hash
}

func (opts *IdemixSignerOpts) HashFunc() crypto.Hash {
	return opts.H
}

// IdemixNymSignerOpts 包含生成idemix假名签名的选项。
type IdemixNymSignerOpts struct {
	// Nym 是要使用的假名
	Nym Key
	IssuerPK Key
	H crypto.Hash
}

// HashFunc 返回用于生成传递给Signer.Sign的消息的散列函数的标识符，否则返回nil，表示没有进行散列。
func (opts *IdemixNymSignerOpts) HashFunc() crypto.Hash {
	return opts.H
}

// IdemixRevocationKeyGenOpts 包含Idemix撤销密钥生成的选项。
type IdemixRevocationKeyGenOpts struct {
	Temporary bool
}

// Algorithm 返回密钥生成算法标识符。
func (opts *IdemixRevocationKeyGenOpts) Algorithm() string {
	return IDEMIX
}

// Ephemeral 如果派生出的密钥必须是暂时的，则该方法返回true，否则返回false。
func (opts *IdemixRevocationKeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

// IdemixRevocationPublicKeyImportOpts 包含导入Idemix撤销公钥的选项。
type IdemixRevocationPublicKeyImportOpts struct {
	Temporary bool
}

// Algorithm 返回密钥生成算法的标识符。
func (opts *IdemixRevocationPublicKeyImportOpts) Algorithm() string {
	return IDEMIX
}

// Ephemeral 如果派生出的密钥必须是暂时的，则该方法返回true，否则返回false。
func (opts *IdemixRevocationPublicKeyImportOpts) Ephemeral() bool {
	return opts.Temporary
}

// IdemixCRISignerOpts 包含生成Idemix CRI的选项，CRI应该由发行机构生成，并可通过使用撤销公钥进行公开验证。
type IdemixCRISignerOpts struct {
	Epoch int
	RevocationAlgorithm RevocationAlgorithm
	UnrevokedHandles [][]byte // revoke的含义是”撤销“
	H crypto.Hash
}

func (opts *IdemixCRISignerOpts) HashFunc() crypto.Hash {
	return opts.H
}