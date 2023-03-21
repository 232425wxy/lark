package bccsp

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