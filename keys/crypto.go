package keys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/mcmx73/easytron/common/base58"
	"github.com/mcmx73/easytron/common/hexutil"
	"github.com/mcmx73/easytron/common/math"
	"github.com/mcmx73/easytron/keys/keccak"
	"github.com/mcmx73/easytron/keys/secp256k1"
	"math/big"
)

func generateAddress() (privateKeyHex, address string, err error) {
	key, err := generateKey()
	if err != nil {
		return "", "", err
	}
	privateKeyBytes := fromECDSA(key)
	privateKeyHex = fmt.Sprintf("%x", privateKeyBytes)
	addressBytes := pubKeyToAddressBytes(key.PublicKey)
	address = base58.CheckEncode(addressBytes, TRON_NETID)
	return privateKeyHex, address, nil
}

func generateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(secp256k1.P256k1(), rand.Reader)
}

func addressFromPrivateKeyHexString(privateKeyStr string) (address string, err error) {
	pkBytes := hexutil.FromHex(privateKeyStr)
	_, publicKey := privateKeyFromBytes(secp256k1.P256k1(), pkBytes)
	addressBytes := pubKeyToAddressBytes(*publicKey)
	return base58.CheckEncode(addressBytes, TRON_NETID), nil
}

func addressFromPrivate(privateKeyStr string) (address string, err error) {
	pkBytes := hexutil.FromHex(privateKeyStr)
	_, publicKey := privateKeyFromBytes(secp256k1.P256k1(), pkBytes)
	addressBytes := pubKeyToAddressBytes(*publicKey)
	return base58.CheckEncode(addressBytes, TRON_NETID), nil
}

func pubKeyToAddressBytes(p ecdsa.PublicKey) []byte {
	pubBytes := fromECDSAPub(&p)
	return keccak.Keccak256(pubBytes[1:])[12:]
}

func privateKeyFromBytes(curve elliptic.Curve, pk []byte) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	x, y := curve.ScalarBaseMult(pk)
	priv := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     x,
			Y:     y,
		},
		D: new(big.Int).SetBytes(pk),
	}
	return priv, &priv.PublicKey
}

func fromECDSA(priv *ecdsa.PrivateKey) []byte {
	if priv == nil {
		return nil
	}
	return math.PaddedBigBytes(priv.D, priv.Params().BitSize/8)
}

func fromECDSAPub(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(secp256k1.P256k1(), pub.X, pub.Y)
}
