# Blockchain cryptographic service provider (bccsp)

fabric的bccsp全称为blockchain cryptographic service provider，区块链密码方案服务提供方，是fabric用于提供密钥生成、加解密以及签名验证的服务接口，接口的定义如下。

## Identity Mixer (idemix)

Idemix（Identity Mixer）的核心是零知识证明（Zero Knowledge Proof），用户无需暴露私有数据以及任何有用的信息，也能证明自己拥有这些私有数据，对方能够进行有效验证，这就是零知识证明。
Idemix是一个密码协议套件（X.509+加密算法），保留隐私实现匿名性，交易时不用透露交易者的身份，而且交易间是无关联的，不可往前追溯。