// This software encrypts and decrypts data given a secret password which is
// encoded in a hash function. The result is an hex encoded string of the
// encripted data.
//
// @author: Rafael Campos Nunes <rafaelnunes@engineer.com>
//

package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

const (
	s = `The secret password used to encrypt or decrypt data`
	d = `The data used to be encrypted or decrypted`
	e = `Encripts the data and return an hex encoded byte string of the encripted
data`
	de = `Decripts the data passed as input`
	// Version is the software number version (semver)
	Version = "0.0.1"
)

var (
	encrypt = flag.Bool("encrypt", false, e)
	decrypt = flag.Bool("decrypt", false, de)
	data    = flag.String("data", "", d)
	secret  = flag.String("secret", "", s)
	version = flag.Bool("version", false, "Shows software version")
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s\n\n", os.Args[0])

		flag.PrintDefaults()
		
		const note = `
Note that sucessive calls for encription does not returns the same output. This 
is the normal behavior of a good encription tool
		`

		fmt.Printf(note)
		fmt.Printf("\n\nExamples:\n\n")
		fmt.Printf(`
$ ./vault -encrypt -data "하파엘" -secret password-so-hard
$ 8ad07a8aa49840e9a8784729119bafd44d077ef649c41a17a1388f426cdab996f7660b64c7

$ ENCRYPTED=8ad07a8aa49840e9a8784729119bafd44d077ef649c41a17a1388f426cdab996f7660b64c7
$ ./vault -decrypt -data $ENCRYPTED -secret password-so-hard
$ 하파엘
		`)
		fmt.Printf("\n하파엘 감ㅂ수 누네수.\n")
	}
}

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("version %s", Version)
		os.Exit(0)
	}

	if *encrypt {
		if len(*data) == 0 && len(*secret) == 0 {
			flag.Usage()
			os.Exit(0)
		}

		encripted := Encrypt([]byte(*data), *secret)
		encoded := make([]byte, hex.EncodedLen(len(encripted)))
		hex.Encode(encoded, encripted)

		fmt.Printf(string(encoded))
	} else if *decrypt {
		if len(*data) == 0 && len(*secret) == 0 {
			flag.Usage()
			os.Exit(0)
		}

		dst := make([]byte, hex.DecodedLen(len(*data)))
		hex.Decode(dst, []byte(*data))

		decripted := Decrypt(dst, *secret)
		
		fmt.Printf(string(decripted))
	} else {
		flag.Usage()
	}
}
