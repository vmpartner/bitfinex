package src

import (
	"encoding/hex"
	"hash"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/hmac"
)

const (
	HashSHA1       = iota
	HashSHA256
	HashSHA512
	HashSHA512_384
)

func HexEncodeToString(input []byte) string {
	return hex.EncodeToString(input)
}

func GetHMAC(hashType int, input, key []byte) []byte {
	var hash_ func() hash.Hash

	switch hashType {
	case HashSHA1:
		hash_ = sha1.New
	case HashSHA256:
		hash_ = sha256.New
	case HashSHA512:
		hash_ = sha512.New
	case HashSHA512_384:
		hash_ = sha512.New384
	}

	hmac_ := hmac.New(hash_, []byte(key))
	hmac_.Write(input)
	return hmac_.Sum(nil)
}
