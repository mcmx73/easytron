package seedphrase

import (
	"github.com/mcmx73/easytron/common/hexutil"
	"testing"
)

const (
	_KEY_BYTES_256 = "f0ff001b62b82d2c0198eb4cc43c68a31e2543b55180dffa34654e3f9048a400"
	_KEY_BYTES_128 = "f0ff00b55180dffa34654e3f9048a400"
)

func TestBitwise(t *testing.T) {
	keyBytes256 := hexutil.FromHex(_KEY_BYTES_256)
	keyBytes128 := hexutil.FromHex(_KEY_BYTES_128)
	keyBytesBadLen := hexutil.FromHex("ffff00b55180dffa34654e3f9048")

	keyBitsArray256 := bytesToBitArray(keyBytes256)
	t.Log(len(keyBitsArray256))
	keyBitsArray128 := bytesToBitArray(keyBytes128)
	t.Log(len(keyBitsArray128))

	signedKey256BitsArray := addChecksumBits(keyBytes256)
	signedKey128BitsArray := addChecksumBits(keyBytes128)
	err := validateChecksumBits(signedKey256BitsArray)
	if err != nil {
		t.Error(err)
	}
	err = validateChecksumBits(signedKey128BitsArray)
	if err != nil {
		t.Error(err)
	}
	t.Log("Signed Key 256:", "(", len(signedKey256BitsArray), ")", signedKey256BitsArray)
	t.Log("Signed Key 128:", "(", len(signedKey128BitsArray), ")", signedKey128BitsArray)
	wordsCount256 := len(signedKey256BitsArray) / 11
	t.Log("256 From 0 to", wordsCount256)
	wordsCount128 := len(signedKey128BitsArray) / 11
	t.Log("128 From 0 to", wordsCount128)
	mnemonic256, err := Bytes2Mnemonic(keyBytes256)
	if err != nil {
		t.Error(err)
	}
	t.Log("Mnemonic 256:", mnemonic256)
	mnemonic128, err := Bytes2Mnemonic(keyBytes128)
	if err != nil {
		t.Error(err)
	}
	t.Log("Mnemonic 128:", mnemonic128)
	mnemonicBad, err := Bytes2Mnemonic(keyBytesBadLen)
	if err == nil {
		t.Error("Bad mnemonic:", mnemonicBad)
	}
	t.Log("Mnemonic Error:", err)
	t.Log("Validate mnemonic word count 256:", validateMnemonicWordCount(mnemonic256))
	t.Log("Validate mnemonic word count 128:", validateMnemonicWordCount(mnemonic128))

	checkBitArray := uint16ToBitArray(0xAF0)
	t.Log(checkBitArray)
	recovered256, _ := Mnemonic2Bytes(mnemonic256)
	recovered128, _ := Mnemonic2Bytes(mnemonic128)
	t.Log("Recovered 256:", hexutil.Encode(recovered256))
	t.Log("         from:", "0x"+_KEY_BYTES_256)
	t.Log("Recovered 128:", hexutil.Encode(recovered128))
	t.Log("         from:", "0x"+_KEY_BYTES_128)
}
