//
// Implements functions of encription and decription
//
// @author: Rafael Campos Nunes <rafaelnunes@engineer.com>
//

package vault

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// Encrypt a given data with a secret
func Encrypt(data []byte, secret string) []byte {
	block, err := aes.NewCipher(MD5(secret))

	if err != nil {
		msg := fmt.Sprintf("Couldn't create a block. Reason %v ", err)
		panic(msg)
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		msg := fmt.Sprintf("Couldn't create a chiper block. Reason %v ",
			err.Error())
		panic(msg)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	encrypted := gcm.Seal(nonce, nonce, data, nil)

	return encrypted
}

// Decrypt data with a given secret
func Decrypt(data []byte, secret string) []byte {
	block, err := aes.NewCipher(MD5(secret))

	if err != nil {
		msg := fmt.Sprintf("Couldn't create a block. Reason %v", err)
		panic(msg)
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		msg := fmt.Sprintf("Couldn't create a wrapper for the block cipher. "+
			"Reason %v", err)
		panic(msg)
	}

	nonceSize := gcm.NonceSize()
	nonce, cipher := data[:nonceSize], data[nonceSize:]

	decrypted, err := gcm.Open(nil, nonce, cipher, nil)

	if err != nil {
		panic(fmt.Sprintf("Couldn't decrypt message. Reason %v", err))
	}

	return decrypted
}
