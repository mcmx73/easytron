// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package secp256k1

import (
	"bufio"
	"compress/bzip2"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"math/big"
	"os"
	"strings"
	"testing"
)

func testKeyGeneration(t *testing.T, c elliptic.Curve, tag string) {
	priv, err := ecdsa.GenerateKey(c, rand.Reader)
	if err != nil {
		t.Errorf("%s: error: %s", tag, err)
		return
	}
	if !c.IsOnCurve(priv.PublicKey.X, priv.PublicKey.Y) {
		t.Errorf("%s: public key invalid: %s", tag, err)
	}
}

func TestKeyGeneration(t *testing.T) {
	testKeyGeneration(t, elliptic.P224(), "p224")
	if testing.Short() {
		return
	}
	testKeyGeneration(t, elliptic.P256(), "p256")
	testKeyGeneration(t, P384(), "p384")
	testKeyGeneration(t, P521(), "p521")
	testKeyGeneration(t, P256k1(), "p256k1")
}

func BenchmarkSignP256(b *testing.B) {
	b.ResetTimer()
	p256 := elliptic.P256()
	hashed := []byte("testing")
	priv, _ := ecdsa.GenerateKey(p256, rand.Reader)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _, _, _ = Sign(rand.Reader, priv, hashed)
		}
	})
}

func BenchmarkSignP256k1(b *testing.B) {
	b.ResetTimer()
	p256k1 := P256k1()
	hashed := []byte("testing")
	priv, _ := ecdsa.GenerateKey(p256k1, rand.Reader)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _, _, _ = Sign(rand.Reader, priv, hashed)
		}
	})
}

func BenchmarkSignP384(b *testing.B) {
	b.ResetTimer()
	p384 := P384()
	hashed := []byte("testing")
	priv, _ := ecdsa.GenerateKey(p384, rand.Reader)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _, _, _ = Sign(rand.Reader, priv, hashed)
		}
	})
}

func BenchmarkVerifyP256(b *testing.B) {
	b.ResetTimer()
	p256 := elliptic.P256()
	hashed := []byte("testing")
	priv, _ := ecdsa.GenerateKey(p256, rand.Reader)
	r, s, _, _ := Sign(rand.Reader, priv, hashed)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ecdsa.Verify(&priv.PublicKey, hashed, r, s)
		}
	})
}

func BenchmarkKeyGeneration(b *testing.B) {
	b.ResetTimer()
	p256 := elliptic.P256()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ecdsa.GenerateKey(p256, rand.Reader)
		}
	})
}

func testSignAndVerify(t *testing.T, c elliptic.Curve, tag string) {
	priv, _ := ecdsa.GenerateKey(c, rand.Reader)

	hashed := []byte("testing")
	r, s, _, err := Sign(rand.Reader, priv, hashed)
	if err != nil {
		t.Errorf("%s: error signing: %s", tag, err)
		return
	}

	if !ecdsa.Verify(&priv.PublicKey, hashed, r, s) {
		t.Errorf("%s: Verify failed", tag)
	}

	hashed[0] ^= 0xff
	if ecdsa.Verify(&priv.PublicKey, hashed, r, s) {
		t.Errorf("%s: Verify always works!", tag)
	}
}

func TestSignAndVerify(t *testing.T) {
	testSignAndVerify(t, elliptic.P224(), "p224")
	if testing.Short() {
		return
	}
	testSignAndVerify(t, elliptic.P256(), "p256")
	testSignAndVerify(t, P384(), "p384")
	testSignAndVerify(t, P521(), "p521")
	testSignAndVerify(t, P256k1(), "p256k1")
}

func testSignAndVerifyASN1(t *testing.T, c elliptic.Curve, tag string) {
	priv, _ := ecdsa.GenerateKey(c, rand.Reader)

	hashed := []byte("testing")
	sig, err := SignASN1(rand.Reader, priv, hashed)
	if err != nil {
		t.Errorf("%s: error signing: %s", tag, err)
		return
	}

	if !ecdsa.VerifyASN1(&priv.PublicKey, hashed, sig) {
		t.Errorf("%s: VerifyASN1 failed", tag)
	}

	hashed[0] ^= 0xff
	if ecdsa.VerifyASN1(&priv.PublicKey, hashed, sig) {
		t.Errorf("%s: VerifyASN1 always works!", tag)
	}
}

func TestSignAndVerifyASN1(t *testing.T) {
	testSignAndVerifyASN1(t, elliptic.P224(), "p224")
	if testing.Short() {
		return
	}
	testSignAndVerifyASN1(t, elliptic.P256(), "p256")
	testSignAndVerifyASN1(t, P384(), "p384")
	testSignAndVerifyASN1(t, P521(), "p521")
	testSignAndVerifyASN1(t, P256k1(), "p256k1")
}

func testNonceSafety(t *testing.T, c elliptic.Curve, tag string) {
	priv, _ := ecdsa.GenerateKey(c, rand.Reader)

	hashed := []byte("testing")
	r0, s0, _, err := Sign(zeroReader, priv, hashed)
	if err != nil {
		t.Errorf("%s: error signing: %s", tag, err)
		return
	}

	hashed = []byte("testing...")
	r1, s1, _, err := Sign(zeroReader, priv, hashed)
	if err != nil {
		t.Errorf("%s: error signing: %s", tag, err)
		return
	}

	if s0.Cmp(s1) == 0 {
		// This should never happen.
		t.Errorf("%s: the signatures on two different messages were the same", tag)
	}

	if r0.Cmp(r1) == 0 {
		t.Errorf("%s: the nonce used for two different messages was the same", tag)
	}
}

func TestNonceSafety(t *testing.T) {
	testNonceSafety(t, elliptic.P224(), "p224")
	if testing.Short() {
		return
	}
	testNonceSafety(t, elliptic.P256(), "p256")
	testNonceSafety(t, P384(), "p384")
	testNonceSafety(t, P521(), "p521")
	testNonceSafety(t, P256k1(), "p256k1")
}

func testINDCCA(t *testing.T, c elliptic.Curve, tag string) {
	priv, _ := ecdsa.GenerateKey(c, rand.Reader)

	hashed := []byte("testing")
	r0, s0, _, err := Sign(rand.Reader, priv, hashed)
	if err != nil {
		t.Errorf("%s: error signing: %s", tag, err)
		return
	}

	r1, s1, _, err := Sign(rand.Reader, priv, hashed)
	if err != nil {
		t.Errorf("%s: error signing: %s", tag, err)
		return
	}

	if s0.Cmp(s1) == 0 {
		t.Errorf("%s: two signatures of the same message produced the same result", tag)
	}

	if r0.Cmp(r1) == 0 {
		t.Errorf("%s: two signatures of the same message produced the same nonce", tag)
	}
}

func TestINDCCA(t *testing.T) {
	testINDCCA(t, elliptic.P224(), "p224")
	if testing.Short() {
		return
	}
	testINDCCA(t, elliptic.P256(), "p256")
	testINDCCA(t, P384(), "p384")
	testINDCCA(t, P521(), "p521")
	testINDCCA(t, P256k1(), "p256k1")
}

func fromHex(s string) *big.Int {
	r, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("bad hex")
	}
	return r
}

func TestVectors(t *testing.T) {
	// This test runs the full set of NIST test vectors from
	// https://csrc.nist.gov/groups/STM/cavp/documents/dss/186-3ecdsatestvectors.zip
	//
	// The SigVer.rsp file has been edited to remove test vectors for
	// unsupported algorithms and has been compressed.

	if testing.Short() {
		return
	}

	f, err := os.Open("testdata/SigVer.rsp.bz2")
	if err != nil {
		t.Fatal(err)
	}

	buf := bufio.NewReader(bzip2.NewReader(f))

	lineNo := 1
	var h hash.Hash
	var msg []byte
	var hashed []byte
	var r, s *big.Int
	pub := new(ecdsa.PublicKey)

	for {
		line, err := buf.ReadString('\n')
		if len(line) == 0 {
			if err == io.EOF {
				break
			}
			t.Fatalf("error reading from input: %s", err)
		}
		lineNo++
		// Need to remove \r\n from the end of the line.
		if !strings.HasSuffix(line, "\r\n") {
			t.Fatalf("bad line ending (expected \\r\\n) on line %d", lineNo)
		}
		line = line[:len(line)-2]

		if len(line) == 0 || line[0] == '#' {
			continue
		}

		if line[0] == '[' {
			line = line[1 : len(line)-1]
			parts := strings.SplitN(line, ",", 2)

			switch parts[0] {
			case "P-224":
				pub.Curve = elliptic.P224()
			case "P-256":
				pub.Curve = elliptic.P256()
			case "P-384":
				pub.Curve = P384()
			case "P-521":
				pub.Curve = P521()
			default:
				pub.Curve = nil
			}

			switch parts[1] {
			case "SHA-1":
				h = sha1.New()
			case "SHA-224":
				h = sha256.New224()
			case "SHA-256":
				h = sha256.New()
			case "SHA-384":
				h = sha512.New384()
			case "SHA-512":
				h = sha512.New()
			default:
				h = nil
			}

			continue
		}

		if h == nil || pub.Curve == nil {
			continue
		}

		switch {
		case strings.HasPrefix(line, "Msg = "):
			if msg, err = hex.DecodeString(line[6:]); err != nil {
				t.Fatalf("failed to decode message on line %d: %s", lineNo, err)
			}
		case strings.HasPrefix(line, "Qx = "):
			pub.X = fromHex(line[5:])
		case strings.HasPrefix(line, "Qy = "):
			pub.Y = fromHex(line[5:])
		case strings.HasPrefix(line, "R = "):
			r = fromHex(line[4:])
		case strings.HasPrefix(line, "S = "):
			s = fromHex(line[4:])
		case strings.HasPrefix(line, "Result = "):
			expected := line[9] == 'P'
			h.Reset()
			h.Write(msg)
			hashed := h.Sum(hashed[:0])
			if ecdsa.Verify(pub, hashed, r, s) != expected {
				t.Fatalf("incorrect result on line %d", lineNo)
			}
		default:
			t.Fatalf("unknown variable on line %d: %s", lineNo, line)
		}
	}
}

func testNegativeInputs(t *testing.T, curve elliptic.Curve, tag string) {
	key, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		t.Errorf("failed to generate key for %q", tag)
	}

	var hash [32]byte
	r := new(big.Int).SetInt64(1)
	r.Lsh(r, 550 /* larger than any supported curve */)
	r.Neg(r)

	if ecdsa.Verify(&key.PublicKey, hash[:], r, r) {
		t.Errorf("bogus signature accepted for %q", tag)
	}
}

func TestNegativeInputs(t *testing.T) {
	testNegativeInputs(t, elliptic.P224(), "p224")
	testNegativeInputs(t, elliptic.P256(), "p256")
	testNegativeInputs(t, P384(), "p384")
	testNegativeInputs(t, P521(), "p521")
	testNegativeInputs(t, P256k1(), "p256k1")
}

func TestZeroHashSignature(t *testing.T) {
	zeroHash := make([]byte, 64)

	for _, curve := range []elliptic.Curve{
		elliptic.P224(),
		elliptic.P256(),
		P384(),
		P521(),
		P256k1(),
	} {
		privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
		if err != nil {
			panic(err)
		}

		// Sign a hash consisting of all zeros.
		r, s, _, err := Sign(rand.Reader, privKey, zeroHash)
		if err != nil {
			panic(err)
		}

		// Confirm that it can be verified.
		if !ecdsa.Verify(&privKey.PublicKey, zeroHash, r, s) {
			t.Errorf("zero hash signature verify failed for %T", curve)
		}
	}
}

func TestSignBytes(t *testing.T) {
	for _, curve := range []elliptic.Curve{
		elliptic.P224(),
		elliptic.P256(),
		P384(),
		P521(),
		P256k1(),
	} {
		privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
		if err != nil {
			panic(err)
		}

		for _, hashed := range [][]byte{
			make([]byte, 64),
			[]byte("testing"),
		} {

			for _, flag := range []byte{
				Normal,
				LowerS,
				RecID,
				LowerS | RecID,
			} {
				b, err := SignBytes(privKey, hashed, flag)
				if err != nil {
					t.Errorf("SignBytes failed for %T", curve)
				}
				testSignBytes(t, &privKey.PublicKey, hashed, b, flag)
			}
		}
	}
}

func testSignBytes(t *testing.T, pubkey *ecdsa.PublicKey, hash, sig []byte, flag byte) {
	if !VerifyBytes(pubkey, hash, sig, flag) {
		t.Error("VerifyBytes failed")
	}

	// test wrong length
	size := len(sig)
	if VerifyBytes(pubkey, hash, sig[:size-1], flag) {
		t.Error("VerifyBytes pass with shorter length")
	}
	sig = append(sig, 0)
	if VerifyBytes(pubkey, hash, sig, flag) {
		t.Error("VerifyBytes pass with longer length")
	}
	sig = sig[:size]

	if (flag & RecID) != 0 {
		// invalid recovery id fails verification
		v := sig[size-1]
		sig[size-1] = 4
		if VerifyBytes(pubkey, hash, sig, RecID) {
			t.Error("VerifyBytes pass with invalid recovery id")
		}
		if VerifyBytes(pubkey, hash, sig, Normal) {
			t.Error("VerifyBytes pass flag = Normal with invalid recovery id")
		}
		sig[size-1] = v

		// remove recovery id fails verification
		if VerifyBytes(pubkey, hash, sig[:size-1], RecID) {
			t.Error("VerifyBytes pass w/o recovery id")
		}
		if !VerifyBytes(pubkey, hash, sig[:size-1], Normal) {
			t.Error("VerifyBytes fail flag = Normal w/o recovery id")
		}
	} else {
		// signature w/o recovery id fails flag = RecID
		if VerifyBytes(pubkey, hash, sig, RecID) {
			t.Error("VerifyBytes pass flag = RecID")
		}
	}

	if (flag & LowerS) != 0 {
		_, s, v := decodeSigBytes(pubkey.Params(), sig)
		curveParam := pubkey.Curve.Params()
		rSize := (curveParam.BitSize + 7) >> 3
		s.Sub(curveParam.N, s)
		s.FillBytes(sig[rSize : 2*rSize])

		// (r, N-s) fails verification
		if VerifyBytes(pubkey, hash, sig, flag) {
			t.Error("VerifyBytes pass with (r, N-s)")
		}

		// (r, N-s) can pass flag = Normal
		if (flag & RecID) != 0 {
			if v > 3 {
				t.Errorf("Invalid recovery id = %d", v)
			}
			sig = sig[:size-1]
		}
		if !VerifyBytes(pubkey, hash, sig, Normal) {
			t.Error("VerifyBytes fail flag = Normal with (r, N-s)")
		}
	}
}
