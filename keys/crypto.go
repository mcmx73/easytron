package keys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/mcmx73/easytron/common/hexutil"
	"github.com/mcmx73/easytron/common/math"
	"github.com/mcmx73/easytron/keys/keccak"
	"github.com/mcmx73/easytron/keys/secp256k1"
	"math/big"
)

func generateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(secp256k1.P256k1(), rand.Reader)
}

func pubKeyToKeccak256HashBytes(p ecdsa.PublicKey) []byte {
	pubBytes := BytesFromECDSAPublicKey(&p)
	return keccak.Keccak256(pubBytes[1:])[12:]
}

func ECDSAKeysFromPrivateKeyHex(pk string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	pkBytes := hexutil.FromHex(pk)
	return ECDSAKeysFromPrivateKeyBytes(pkBytes)
}
func ECDSAKeysFromPrivateKeyBytes(pk []byte) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	curve := secp256k1.P256k1()
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

func BytesFromECDSAPrivateKey(privateKey *ecdsa.PrivateKey) []byte {
	if privateKey == nil {
		return nil
	}
	return math.PaddedBigBytes(privateKey.D, privateKey.Params().BitSize/8)
}

func BytesFromECDSAPublicKey(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(secp256k1.P256k1(), pub.X, pub.Y)
}
