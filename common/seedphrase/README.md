## seedprase

Create mnemonic phrase from 128/256bit data array. Word list taken from BIP-39 standard.
This package demonstrate how to you can create words mnemonic from any data array.

[!] But please note, this is not BIP-39 implementation, it's just a educational mnemonic generator.

Example:
```go

import (
    "github.com/github.com/mcmx73/easytron/common/seedphrase"
)

func main(){
	...
    // create mnemonic from 128bit private key data:
    mnemonic, err := seedphrase.Bytes2Mnemonic(myKeyBytes)
    // restore key from mnemonic:
    keyBytes, err := seedphrase.Mnemonic2Bytes(myMnemonic)
	...
}

```

### TODO:
- [ ] Add tests
- [ ] Implement BIP-39 standard

### Attention! Don't use this in production code!