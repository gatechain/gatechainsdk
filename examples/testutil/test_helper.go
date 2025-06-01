package testutil

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/gatechain/crypto/keys"
)

type PrivateKeyInfo struct {
	Name       string
	Content    string
	PassPhrase string
}

var AllPrivateKeys = []PrivateKeyInfo{
	{
		Name:       "validator1",
		PassPhrase: "12345678",
		Content: `-----BEGIN GATECHAIN PRIVATE KEY-----
kdf: bcrypt
salt: 079483018A5BCDBC0D09544DDDC6986F

qV8Zdwx2YXeZ5kd+mEPwhHaIPGYaEcC6Hl5tJyDGGqs1+NeF2PKQQFIRvBKJyd1C
0ZAjh7HCy4HBRnxbtqxNh3XnY7X1FoFBk1Dq+4vCYpgewLSzcZ9LNkwPbN4JYP+q
9FOGx0r4JGk9OyBe3bKs3ZIJBK7fzmk37yWQbAgwUA44qwoElPQljSabrTm6
=vOf7
-----END GATECHAIN PRIVATE KEY-----`,
	},
	{
		Name:       "delegator",
		PassPhrase: "12345678",
		Content: `-----BEGIN GATECHAIN PRIVATE KEY-----
kdf: bcrypt
salt: 15277662EC61A86B1C549572DE8FC549

kJp4h7n2+lmLRaybDLAa0hka4QwJNb6Di41EWY7850NsDc5QE2BQ3C3OCLF0pziC
3pj68gdNWQ/WEdRvTDH0WiJHO2bptwXTwBg6++nL6eGrc0LXGRbvspYPqIqXcP9S
kuw5rW0F5rnxYhd4tJSChT5DitMabDLng8j9Pb3R+r8IKdFN517J/NX2b6Mu
=WFSc
-----END GATECHAIN PRIVATE KEY-----`,
	},
	{
		Name:       "toaccount1",
		PassPhrase: "12345678",
		Content: `-----BEGIN GATECHAIN PRIVATE KEY-----
kdf: bcrypt
salt: 81C055EDB525BD97754F6760619CB25A

1b/89+OgGAm4TzqjUG2qBrgwlGLooEVttAyHsFJ4MD9M09u8pPbdl99buTOVMI9+
oJNq30Mkj2c5UppXwOCEmcjKS3YiwYkWD5pP4IwOtdpdOozSNshqk8L51aPl6Ozp
kfURK2cjjUdk/mSqmr668cdVN5ewhnbmxlhDdH4qa+pvVUdWCHmowtK0RsXj
=RHn6
-----END GATECHAIN PRIVATE KEY-----`,
	},
	{
		Name:       "security_account1",
		PassPhrase: "12345678",
		Content: `-----BEGIN GATECHAIN PRIVATE KEY-----
kdf: bcrypt
salt: AB0B32D9565A673832A40C7F824C344D

guf9qnq6T+coR2CWjrFk5ASN53dp3UYRWPuhWxez/37gon31KC74TGECfEZB08A1
jVL4d8ZtY7Yadqbnza5mk6sizn4JggMarzx/kCXnHVIBQq8y5BEp4naFeM6LtWWu
8JMUpQg1N1cpK5Zfl6J0dbnnb22+auApRJH4rGbM3CqN5T9TfINClnxh6nsa
=NIUL
-----END GATECHAIN PRIVATE KEY-----`,
	},
}

// ImportAllPrivateKeys
func ImportAllPrivateKeys() keys.Keybase {
	kb := keys.NewInMemory()
	for _, pk := range AllPrivateKeys {
		if err := kb.ImportPrivKey(pk.Name, pk.Content, pk.PassPhrase); err != nil {
			fmt.Printf("[Warn] Import %s failed: %v\n", pk.Name, err)
		} else {
			fmt.Printf("[Info] %s key imported successfully\n", pk.Name)
		}
	}
	return kb
}
