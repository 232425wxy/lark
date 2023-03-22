package pkcs11

import "time"

const (
	defaultCreateSessionRetries    = 10
	defaultCreateSessionRetryDelay = 100 * time.Millisecond
	defaultSessionCacheSize        = 10
)

// PKCS11Opts 包含公钥加密标准的选项
type PKCS11Opts struct {
	// TODO fabric 官方在考虑舍弃掉？
	Security int    `json:"security"`
	Hash     string `json:"hash"`

	Library        string `json:"library"`
	Label          string `json:"label"`
	Pin            string `json:"pin"`
	SoftwareVerify bool   `json:"softwareverify,omitempty"`
	Immutable      bool   `json:"immutable,omitempty"`

	sessionCacheSize        int
	createSessionRetries    int
	createSessionRetryDelay time.Duration
}
