package keys

import "testing"

func TestNewAddress(t *testing.T) {
	key, address, err := generateAddress()
	if err != nil {
		t.Error(err)
	}
	t.Log(address, key)
}
