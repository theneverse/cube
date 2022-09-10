package hash

import (
	"crypto/sha256"
	"math/big"
)

func Hash(content []byte) (*big.Int, error) {
	h := sha256.New()
	_, err := h.Write(content)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(h.Sum(nil)), nil
}
