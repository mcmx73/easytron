package keys

import (
	"fmt"
	"github.com/mcmx73/easytron/common/base58"
)

const (
	HEX_PREFIX = "41"
	TRON_NETID = 65
)

func tronNewAddress() (privateKeyHex, address string, err error) {
	key, err := generateKey()
	if err != nil {
		return "", "", err
	}
	privateKeyBytes := BytesFromECDSAPrivateKey(key)
	privateKeyHex = fmt.Sprintf("%x", privateKeyBytes)
	addressBytes := pubKeyToKeccak256HashBytes(key.PublicKey)
	address = base58.CheckEncode(addressBytes, TRON_NETID)
	return privateKeyHex, address, nil
}
