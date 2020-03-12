//
//
// @author: Rafael Campos Nunes <rafaelnunes@engineer.com>
//

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// SHA1 returns an hex encoded 160 bit hash of whatever data is passed
func SHA1(data string) []byte {
	hasher := sha1.New()
	return hashit(hasher, data)
}

// SHA512 returns an hex encoded 512 bit hash of whatever data is passed
func SHA512(data string) []byte {
	hasher := sha512.New()
	return hashit(hasher, data)
}

// MD5 returns an hex encoded 512 bit hash of whatever data is passed
func MD5(data string) []byte {
	hasher := md5.New()
	return hashit(hasher, data)
}

func hashit(h hash.Hash, data string) []byte {
	h.Write([]byte(data))
	result := h.Sum(nil)

	b := make([]byte, hex.EncodedLen(len(result)))
	hex.Encode(b, result)

	return b
}
