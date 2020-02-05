package session

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/qlova/seed"
)

var cachedKey []byte

//Key returns the key used by session to encrypt data. Reads key from SESSION_KEY env.
//Will create a key and store it in seed.Dir if it env is not set.
func Key() (key [32]byte) {
	if env := os.Getenv("SESSION_KEY"); env != "" {
		copy(key[:], env)
		return
	}

	if cachedKey == nil {

		var err error
		cachedKey, err = ioutil.ReadFile(seed.Dir + "/session.key")
		if err != nil {
			cachedKey = make([]byte, 32)
			_, err = rand.Read(cachedKey)
			if err != nil {
				_, err = rand.Read(cachedKey)
				if err != nil {
					fmt.Println(err)
				}
			}

			err := ioutil.WriteFile(seed.Dir+"/session.key", cachedKey, 0755)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

	copy(key[:], cachedKey)
	return
}

//Encrypt encrypts data with the session encryption scheme.
func Encrypt(data []byte) string {
	var key = Key()

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key[:])
	// if there are any errors, handle them
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
		return ""
	}

	var encrypted = gcm.Seal(nonce, nonce, data, nil)

	var result bytes.Buffer
	var encoder = base64.NewEncoder(base64.URLEncoding, &result)
	_, err = encoder.Write(encrypted)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	encoder.Close()

	return result.String()
}

func Decrypt(data string) []byte {

	var decoder = base64.NewDecoder(base64.URLEncoding, strings.NewReader(data))

	var encrypted, err = ioutil.ReadAll(decoder)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	key := Key()

	c, err := aes.NewCipher(key[:])
	if err != nil {
		fmt.Println(err)
		return nil
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	nonceSize := gcm.NonceSize()
	if len(encrypted) < nonceSize {
		fmt.Println(err)
		return nil
	}

	nonce, ciphertext := []byte(encrypted[:nonceSize]), []byte(encrypted[nonceSize:])
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return plaintext
}
