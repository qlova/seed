package session

import (
	"bytes"
	"testing"
)

func TestEncrypt(t *testing.T) {
	if !bytes.Equal(Decrypt(Encrypt([]byte("Hello World"))), []byte("Hello World")) {
		t.Fail()
	}
	if Encrypt([]byte("Hello World")) == "Hello World" {
		t.Fail()
	}
}
