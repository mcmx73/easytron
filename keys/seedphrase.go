package keys

import "github.com/mcmx73/easytron/common/seedphrase"

func (k *Key) GetKeyBackupPhrase() []string {
	if k == nil {
		panic("Key is nil")
	}
	if k.privateKey == nil && k.PrivateKey == "" {
		panic("Key is empty")
	} else if k.privateKey == nil {
		k.privateKey, _ = ECDSAKeysFromPrivateKeyHex(k.PrivateKey)
		if k.privateKey == nil {
			panic("Key is invalid")
		}
	}
	privateKeyBytes := BytesFromECDSAPrivateKey(k.privateKey)
	backupPhrase, err := seedphrase.Bytes2Mnemonic(privateKeyBytes)
	if err != nil {
		panic(err)
	}
	return backupPhrase
}
