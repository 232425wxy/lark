package bccsp

import (
	"crypto"
	"hash"
)

// Key 代表加密方案中的密钥。
type Key interface {
	// Bytes 将密钥转换为原始的字节切片。
	Bytes() ([]byte, error)

	// SKI 返回该密钥的标识符：subject key identifier。
	SKI() []byte

	// Symmetric 如果该密钥是对称密钥，则返回true，如果是非对称密钥，则返回false。
	Symmetric() bool

	// Private 如果该密钥是一个私钥，则返回true，否则返回false。
	Private() bool

	// PublicKey 如果该密钥是非对称密钥方案里的密钥，则返回公私钥对相对应的公钥，如果该密钥是对称密钥方案里的密钥，调用此方法则会返回错误。
	PublicKey() (Key, error)
}

// KeyGenOpts 包含用密码方案服务提供商(Cryptographic Service Provider, CSP)生成密钥的选项。
type KeyGenOpts interface {
	// Algorithm 返回密钥生成算法标识符。
	Algorithm() string

	// Ephemeral 如果要生成的密钥必须是短暂的，则Ephemeral返回true，否则返回false。
	Ephemeral() bool
}

// KeyDerivOpts 包含用密码方案服务提供商(Cryptographic Service Provider, CSP)进行密钥派生的选项。
type KeyDerivOpts interface {
	// Algorithm 返回密钥派生算法标识符。
	Algorithm() string

	// Ephemeral 如果要派生的密钥必须是短暂的，则返回true，否则返回false。
	Ephemeral() bool
}

// KeyImportOpts 包含用密码方案服务提供商(Cryptographic Service Provider, CSP)导入密钥原材料的选项。
type KeyImportOpts interface {
	// Algorithm 返回密钥导入算法的标识符。
	Algorithm() string

	// Ephemeral 如果生成的密钥必须是短暂的，则返回true，否则返回false。
	Ephemeral() bool
}

// HashOpts 包含用密码方案服务提供商(Cryptographic Service Provider, CSP)进行哈希计算的选项。
type HashOpts interface {
	// Algorithm 返回哈希算法的标识符。
	Algorithm() string
}

// SignerOpts 包含用密码方案服务提供商(Cryptographic Service Provider, CSP)进行签名的选项。
type SignerOpts interface {
	crypto.SignerOpts
}

// EncrypterOpts 包含用密码方案服务提供商(Cryptographic Service Provider, CSP)进行加密的选项。
type EncrypterOpts interface{}

// DecrypterOpts 包含用密码方案服务提供商(Cryptographic Service Provider, CSP)进行解密的选项。
type DecrypterOpts interface{}

// BCCSP 是区块链加密服务提供商，提供加密标准和算法的实施。
type BCCSP interface {
	// KeyGen 根据给定的密钥生成选项生成一个密钥。
	KeyGen(opts KeyGenOpts) (k Key, err error)

	// KeyDeriv 使用给定的密钥派生选项从给定的密钥派生出一个密钥。
	KeyDeriv(k Key, opts KeyDerivOpts) (dk Key, err error)

	// KeyImport 使用opts从其原始数据中导入一个密钥。
	KeyImport(raw interface{}, opts KeyImportOpts) (k Key, err error)

	// GetKey 返回该CSP与主题密钥标识符ski相关的密钥。
	GetKey(ski []byte) (k Key, err error)

	// Hash 根据给定的哈希算法选项求给定的消息的哈希值，如果选项是空的，则采用默认的哈希算法求哈希值。
	Hash(msg []byte, opts HashOpts)

	// GetHash 根据给定的选项返回hash.Hash实例，如果选项是空的，则返回默认的哈希函数。
	GetHash(opts HashOpts) (h hash.Hash, err error)

	// Sign 给定密钥k、消息的摘要值digest和签名选项opts，对消息的摘要值进行签名。注意，签名选项opts决定了采用什么签名算法。
	Sign(k Key, digest []byte, opts SignerOpts) (signature []byte, err error)

	// Verify 对签名进行验证。
	Verify(k Key, signature, digest []byte, opts SignerOpts) (valid bool, err error)

	// Encrypt 利用给定的密钥k和加密选项opts，对给定的明文plaintext进行加密。
	Encrypt(k Key, plaintext []byte, opts EncrypterOpts) (ciphertext []byte, err error)

	// Decrypt 利用给定的密钥k解密密文得到明文。
	Decrypt(k Key, ciphertext []byte, opts DecrypterOpts) (plaintext []byte, err error)
}
