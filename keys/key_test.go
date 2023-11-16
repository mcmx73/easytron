package keys

import (
	"github.com/mcmx73/easytron/common/hexutil"
	"github.com/mcmx73/easytron/common/seedphrase"
	"testing"
)

func TestECDSAKeys(t *testing.T) {
	key, err := generateKey()
	if err != nil {
		t.Error(err)
	}
	publicKeyBytes := BytesFromECDSAPublicKey(&key.PublicKey)
	privateKeyBytes := BytesFromECDSAPrivateKey(key)
	t.Log("Private key len:", len(privateKeyBytes), "/", len(privateKeyBytes)*8, "bit")
	t.Log("Public key len:", len(publicKeyBytes), "/", len(publicKeyBytes)*8, "bit")
	pubHex := hexutil.Encode(publicKeyBytes)
	privateHex := hexutil.Encode(privateKeyBytes)
	recoveredPrivateKey, _ := ECDSAKeysFromPrivateKeyBytes(privateKeyBytes)
	recoveredPKBytes := BytesFromECDSAPrivateKey(recoveredPrivateKey)
	recoveredPKHex := hexutil.Encode(recoveredPKBytes)
	t.Log(pubHex, "\n", privateHex, "\n", recoveredPKHex)
	if privateHex != recoveredPKHex {
		t.Error("ECDSAKeysFromPrivateKeyBytes: Private key mismatch")
	}
	recoveredFromHex, _ := ECDSAKeysFromPrivateKeyHex(privateHex)
	recoveredFromHexBytes := BytesFromECDSAPrivateKey(recoveredFromHex)
	recoveredFromHexHex := hexutil.Encode(recoveredFromHexBytes)
	if privateHex != recoveredFromHexHex {
		t.Error("ECDSAKeysFromPrivateKeyHex: Private key mismatch")
	}
	backupPhrase, err := seedphrase.Bytes2Mnemonic(privateKeyBytes)
	if err != nil {
		t.Error(err)
	}
	t.Log("Backup phrase:", backupPhrase)
	recoveredFromBackup, err := seedphrase.Mnemonic2Bytes(backupPhrase)
	if err != nil {
		t.Error(err)
	}
	if hexutil.Encode(recoveredFromBackup) != privateHex {
		t.Error("Backup phrase mismatch")
	}
	t.Log("Restored key:", hexutil.Encode(recoveredFromBackup))
}
