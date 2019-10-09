package password

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

//Hash is a base64 encoded 256-bit password with entropy distributed across 256 bits.
//The passwordbox HashAndGo method can derive a password on the client side that is suitable as an input to this function.
type Hash string

//Key returns the 256-bit key encoded in this hash.
func (hash Hash) Key() ([32]byte, error) {
	var decoder = base64.NewDecoder(base64.StdEncoding, strings.NewReader(string(hash)))

	var decoded, err = ioutil.ReadAll(decoder)
	if err != nil {
		return [32]byte{}, err
	}

	var key [32]byte
	copy(key[:], decoded)
	return key, nil
}

//PasswordFor returns the username encrypted with the password encoded as base64.
func (hash Hash) PasswordFor(username string) (string, error) {
	var key, err = hash.Key()
	if err != nil {
		return "", err
	}

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key[:])
	// if there are any errors, handle them
	if err != nil {
		return "", err
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		return "", err
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	var encrypted = gcm.Seal(nonce, nonce, []byte(username), nil)

	var result bytes.Buffer
	var encoder = base64.NewEncoder(base64.StdEncoding, &result)
	_, err = encoder.Write(encrypted)
	if err != nil {
		return "", err
	}
	encoder.Close()

	return result.String(), nil
}

//VerifyFor verifies a previous result of a hash.PasswordFor result.
func (hash Hash) VerifyFor(username, password string) bool {
	var decoder = base64.NewDecoder(base64.StdEncoding, strings.NewReader(password))

	var encrypted, err = ioutil.ReadAll(decoder)
	if err != nil {
		fmt.Println(err)
	}

	key, err := hash.Key()
	if err != nil {
		return false
	}

	c, err := aes.NewCipher(key[:])
	if err != nil {
		return false
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return false
	}

	nonceSize := gcm.NonceSize()
	if len(encrypted) < nonceSize {
		return false
	}

	nonce, ciphertext := []byte(encrypted[:nonceSize]), []byte(encrypted[nonceSize:])
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return false
	}

	if string(plaintext) == username {
		return true
	}

	return false
}
