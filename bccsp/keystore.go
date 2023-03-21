package bccsp

// KeyStore 表示一个加密密钥的存储系统，它允许存储和检索bccsp.Key对象。
// KeyStore可以是只读的，在这种情况下调用StoreKey方法将返回一个错误。
type KeyStore interface {
	// ReadOnly 如果KeyStore是只读的，则该方法返回true，否则返回false。在KeyStore是
	// 只读的情况下，调用StoreKey方法会失败。
	ReadOnly() bool

	// GetKey 给定密钥的标识符SKI，返回对应的密钥。
	GetKey(ski []byte) (key Key, err error)

	// StoreKey 在KeyStore里存储给定的密钥，如果KeyStore是只读的，调用此方法则会失败。
	StoreKey(k Key) (err error)
}